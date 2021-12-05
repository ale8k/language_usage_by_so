########
# BASE #
########
FROM golang:1.17.3 AS base
WORKDIR /usr/src/app

#######
# DEV #
#######
FROM base AS dev
WORKDIR /usr/src/app
RUN go get -u github.com/cosmtrek/air && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 9000
EXPOSE 2345
ENTRYPOINT ["air"]

#########
# BUILD #
#########
FROM base AS build 
WORKDIR /usr/src/app
RUN go mod download
RUN mkdir ./build
RUN go build -o ./statconsumer

########
# PROD #
########
FROM alpine:3.15
WORKDIR /usr/src/app
EXPOSE 9000
COPY --from=build /usr/src/app/statconsumer /usr/src/app/statconsumer
RUN chmod +x /usr/src/app/statconsumer
ENTRYPOINT ["/usr/src/app/statconsumer"]