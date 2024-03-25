FROM mysql

COPY ./schema.sql /docker-entrypoint-initdb.d/

ENV MYSQL_ROOT_PASSWORD livraria2024
ENV MYSQL_DATABASE livraria

EXPOSE 3306
