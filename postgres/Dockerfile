# Dockerfile metadata
FROM postgres:16
LABEL version="${VERSION}"
LABEL description="ctfconsole | a Capture the Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# make a folder for logs
RUN mkdir -p /var/log/postgresql

# copy the db schema init script
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql

# start postgres
CMD [ "postgres", \
  "-c", "logging_collector=on", \
  "-c", "log_destination=jsonlog", \
  "-c", "log_directory=/var/log/postgresql", \
  "-c", "log_filename='postgresql-%a.log'", \
  "-c", "log_statement=all" \
]