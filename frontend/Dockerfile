# Dockerfile metadata
FROM node:13.12.0-alpine
LABEL version="${VERSION}"
LABEL description="ctfconsole | a Capture the Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# add "/app/node_modules/.bin" to $PATH
WORKDIR /opt/ctfconsole/
ENV PATH /opt/ctfconsole/node_modules/.bin:$PATH

# install dependencies
COPY package.json .
COPY package-lock.json .
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

# copy source code
COPY . .

# start ctfconsole frontend
CMD [ "npm", "start" ]