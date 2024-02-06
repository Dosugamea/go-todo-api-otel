FROM golang:1.21-alpine AS build

WORKDIR /temp
COPY ./go.mod ./go.sum ./
RUN go mod download
RUN apk add --update gcc musl-dev
ADD . /temp
RUN CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" GOOS=linux GOARCH=amd64 go build -o api .

FROM alpine:3.19.1 AS final
LABEL maintainer="dosugamea"
WORKDIR /

COPY --from=build /temp/api /api

EXPOSE 8080/tcp
CMD ["/api"]