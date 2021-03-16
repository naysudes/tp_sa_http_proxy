# Build
FROM golang:1.16-buster AS build

WORKDIR /app
ADD . .

ENV CGO_ENABLED=0

RUN go mod tidy
RUN go mod vendor
RUN go build -o  ./main.go

EXPOSE 8080

CMD ["./main.go"]