#stage 1
FROM node:22 AS node
WORKDIR /app
RUN npm install -g @angular/cli@18
COPY package.json package-lock.json ./
RUN npm i
COPY . .
EXPOSE 4200
ENTRYPOINT [ "ng", "serve", "--host", "0.0.0.0", "--poll", "2000"]