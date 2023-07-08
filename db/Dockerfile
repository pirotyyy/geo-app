FROM mysql:8.0

ENV TZ=Asia/Tokyo

COPY db/my.cnf /etc/mysql/conf.d/my.cnf
COPY db/initdb.d /docker-entrypoint-initdb.d

CMD [ "mysqld", "--character-set-server=utf8", "--collation-server=utf8_unicode_ci" ]