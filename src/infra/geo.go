package infra

import (
	"database/sql"
	"encoding/json"
	"errors"
	"go-geo-backend/src/domain/model"
	"go-geo-backend/src/domain/repository"
	"io"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type GeoRepository struct {
	SqlHandler
}

func NewGeoRepository(sqlHandler SqlHandler) repository.GeoRepository {
	geoRepository := GeoRepository{sqlHandler}
	return &geoRepository
}

func (geoRepo *GeoRepository) GetGeo(postalCode string) (geoInfo *model.GeoInfo, err error) {
	url := "https://geoapi.heartrails.com/api/json?method=searchByPostal&postal=" + postalCode

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geoApiResp model.GeoApiResponse
	if err := json.Unmarshal(body, &geoApiResp); err != nil {
		return nil, err
	}

	if geoApiResp.Response.Location == nil {
		return nil, errors.New(postalCode + " does not exist")
	}

	locations := make([]*model.Location, 0)
	for i := range geoApiResp.Response.Location {
		locations = append(locations, &geoApiResp.Response.Location[i])
	}

	commonPart := findCommonPart(locations)

	geoInfo = &model.GeoInfo{}
	geoInfo.PostalCode = postalCode
	geoInfo.HitCount = len(locations)
	geoInfo.Address = commonPart
	geoInfo.TokyoStaDistance = maxTokyoDistance(locations)

	addAccessLog(geoRepo.SqlHandler.Conn, postalCode)

	return
}

func (geoRepo *GeoRepository) GetAccessLogs() (accessLogs []*model.AccessLog, err error) {
	res, err := geoRepo.SqlHandler.Conn.Query("SELECT postal_code, COUNT(id) FROM access_logs GROUP BY postal_code ORDER BY COUNT(id) DESC")
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var postalCode string
		var requestCount int
		var accessLog model.AccessLog
		res.Scan(&postalCode, &requestCount)
		accessLog.PostalCode = postalCode
		accessLog.RequestCount = requestCount
		accessLogs = append(accessLogs, &accessLog)
	}

	return accessLogs, nil
}

func findCommonPart(locations []*model.Location) string {
	re := regexp.MustCompile(`(.+町)?(.+)`)
	base := re.FindStringSubmatch(locations[0].Town)
	commonParts := []string{locations[0].Prefecture, locations[0].City, base[1], base[2]}
	flags := []bool{true, true}

	for i := 1; i < len(locations)-1; i++ {
		if flags[0] && flags[1] {
			matches := re.FindStringSubmatch(locations[i].Town)

			if strings.Compare(base[1], matches[1]) != 0 {
				commonParts[2] = ""
				flags[0] = false
			}

			if strings.Compare(base[2], matches[2]) != 0 || base[2] == "（その他）" {
				commonParts[3] = ""
				flags[1] = false
			}
		} else {
			break
		}
	}

	return strings.Join(commonParts, "")
}

func maxTokyoDistance(locations []*model.Location) float64 {
	dist := 0.0
	for _, location := range locations {
		x, _ := strconv.ParseFloat(location.X, 64)
		y, _ := strconv.ParseFloat(location.Y, 64)

		if dist < distance(x, y) {
			dist = distance(x, y)
		}
	}

	dist = math.Round(dist*10) / 10

	return dist
}

func distance(x float64, y float64) float64 {
	const R = 6371.0
	const tokyoX = 139.7673068
	const tokyoY = 35.6809591
	radian := math.Pi * R / 180

	return radian * math.Pow(math.Pow((x-tokyoX)*math.Cos(math.Pi*(y+tokyoY)/360), 2)+math.Pow(y-tokyoY, 2), 0.5)
}

func addAccessLog(client *sql.DB, postalCode string) error {
	ins, err := client.Prepare("INSERT INTO " + "access_logs" + "(postal_code,created_at) VALUES(?,?)")
	if err != nil {
		return err
	}

	ins.Exec(postalCode, time.Now())
	return nil
}
