FROM golang:1.17-alpine as base_build

ENV GO111MODULE=on
RUN apk update && apk add --no-cache gcc openssh git bash libc-dev ca-certificates make g++
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p /app/src/handler
WORKDIR /app/src/handler

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/src/handler/fetchapp /app/src/handler/main.go

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates tzdata && rm -rf /var/cache/apk/*
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app/src/handler
COPY --from=base_build /app/src/handler/fetchapp .

EXPOSE 7799
ENTRYPOINT [ "./fetchapp" ]