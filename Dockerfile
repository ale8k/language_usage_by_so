#########
# BUILD #
#########
FROM golang:1.17.3 AS build

WORKDIR /usr/src/app

COPY ./go.mod ./
COPY ./go.sum ./
COPY ./*.go ./

RUN go mod download
RUN mkdir ./build
RUN go build -o ./app

#######
# RUN #
#######
FROM alpine:3.15

WORKDIR /app

EXPOSE 8000

COPY --from=build /usr/src/app /app

RUN chmod +x /app/app

ENTRYPOINT ["/app/app"]