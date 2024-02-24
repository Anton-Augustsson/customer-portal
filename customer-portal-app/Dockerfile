FROM node:19.5.0-alpine3.17 AS builder


WORKDIR /usr/src/app


COPY package*.json ./


RUN npm install


COPY . .


RUN npm run build --prod


FROM nginx:1.19-alpine


RUN rm -rf /usr/share/nginx/html/*


COPY --from=builder /usr/src/app/dist/customer-portal-app /usr/share/nginx/html


EXPOSE 80
EXPOSE 443