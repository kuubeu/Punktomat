FROM node:alpine3.12

WORKDIR /code

COPY package.json yarn.lock ./

RUN yarn

COPY . .
