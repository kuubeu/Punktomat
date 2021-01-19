FROM node:alpine3.12 as build

WORKDIR /code

COPY package.json yarn.lock ./

RUN yarn

COPY . .

RUN yarn build

FROM nginx:1.19.6-alpine

COPY --from=build /code/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
