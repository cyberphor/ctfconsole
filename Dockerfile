# Dockerfile metadata
FROM node:latest

WORKDIR /opt/ctfconsole/
RUN npx create-react-app . &&\
    rm src/App.js
COPY App.js src/App.js
CMD [ "npm", "start" ]
