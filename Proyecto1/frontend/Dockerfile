FROM node:16-alpine as build

WORKDIR /app

COPY . .

RUN npm install

CMD [ "npm","run","build","--prod" ]

FROM nginx:alpine

COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80

