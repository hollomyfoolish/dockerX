FROM node:8
RUN mkdir -p /usr/app
COPY ./src/ /usr/app/src
COPY ./node_modules /usr/app/node_modules
WORKDIR /usr/app
ENTRYPOINT ["node", "./src/app.js"]
EXPOSE 5443:4430