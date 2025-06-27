
FROM golang:1.21

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o edge-gateway

EXPOSE 8080
CMD ["./edge-gateway"]