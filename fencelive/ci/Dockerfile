FROM golang:1.20-alpine3.17

ADD ./fencelive/build/ /app/
ADD ./fencelive/configurations /app/configurations/
ADD ./fencelive/migrations /app/migrations


WORKDIR /app

ENTRYPOINT ./start_server
