FROM golang:1.21-alpine

WORKDIR /app

COPY . . 

RUN apk add gcc musl-dev

ENV CGO_ENABLED=1 

RUN go build /app/server/main.go 

EXPOSE 8080

CMD ["/app/main"] 