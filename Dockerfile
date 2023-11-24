FROM postgres
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
COPY db/db.sql /docker-entrypoint-initdb.d/
