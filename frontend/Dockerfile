# Dockerfile metadata
FROM node:13.12.0-alpine
LABEL version="v0.3.0"
LABEL description="ctfconsole is a Capture The Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# add `/app/node_modules/.bin` to $PATH
WORKDIR /opt/ctfconsole/frontend
ENV PATH /opt/ctfconsole/frontend/node_modules/.bin:$PATH

# install dependencies
COPY package.json .
COPY package-lock.json .
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

# copy ctfconsole frontend source code
COPY . ./

# start ctfconsole frontend
CMD ["npm", "start"]