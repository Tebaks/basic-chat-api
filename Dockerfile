FROM golang:1.17.1 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /chatapp

FROM alpine:3.11.3

WORKDIR /

COPY --from=build /chatapp /chatapp
COPY --from=build /app/.env ./.env

ENTRYPOINT [ "/chatapp" ]