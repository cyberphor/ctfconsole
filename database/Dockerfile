# Dockerfile metadata
FROM postgres:16
LABEL version="${VERSION}"
LABEL description="ctfconsole | a Capture the Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# copy the db schema init script
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql
