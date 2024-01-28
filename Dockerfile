# syntax=docker/dockerfile:1

FROM golang:1.21.5 as base

FROM base as dev

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app/api
CMD ["air"]