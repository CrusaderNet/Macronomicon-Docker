FROM golang:1.22.4
WORKDIR /app
COPY . ./
RUN go build -o /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]