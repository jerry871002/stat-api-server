FROM postgres:16.4

ENV POSTGRES_DB=mydatabase
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword

COPY ./init-table.sql /docker-entrypoint-initdb.d/

EXPOSE 5432