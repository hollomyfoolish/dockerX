FROM node:8
RUN mkdir /usr/app
ADD ./node_modules ./usr/app/node_modules
ADD ./src /usr/app/src
ADD ./package.json /usr/app
ADD ./static /usr/app/static
WORKDIR /usr/app
ENTRYPOINT [ "node", "./src/app.js" ]
EXPOSE 3003
