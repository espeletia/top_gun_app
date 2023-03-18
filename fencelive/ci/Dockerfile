FROM golang:1.20-alpine3.17
ARG app="fencelive"

ADD ./${app}/configurations /app/configurations
ADD ./${app}/build /app

WORKDIR /app

ENTRYPOINT ./start_server
