FROM golang:1.22 AS backend

RUN apt-get update && apt-get install -y git gcc make libc-dev

COPY ./src /app

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod download

RUN make build

FROM node:20 AS frontend

COPY ./src/web /app

WORKDIR /app

RUN npm install
RUN npm run build

FROM ubuntu:latest

RUN apt-get update && apt-get install -y sqlite3

COPY --from=backend /app/configs /app/configs
COPY --from=backend /app/db /app/db
COPY --from=backend /app/build/cloudsdale /app/cloudsdale
COPY --from=frontend /app/dist /app/dist
COPY ./service/docker-entrypoint.sh /

WORKDIR /app

EXPOSE 8888

ENTRYPOINT ["/bin/bash","/docker-entrypoint.sh"]