#stage 1
FROM node:latest as node
WORKDIR /app
COPY . .
RUN npm install
RUN npm install -g @angular/cli@7.3.9
EXPOSE 4200
ENTRYPOINT [ "ng", "serve", "--host", "0.0.0.0"]