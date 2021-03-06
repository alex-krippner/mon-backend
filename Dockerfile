# syntax=docker/dockerfile:1

FROM golang:1.18 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app/api
CMD ["air"]


# Build an artifact that can be deployed to a server

# FROM base as built

# WORKDIR /go/app/api
# COPY . .

# ENV CGO_ENABLED=0

# RUN go get -d -v ./...
# RUN go build -o /tmp/api-server ./*.go

# FROM busybox

# COPY --from=built /tmp/api-server /usr/bin/api-server
# CMD ["api-server", "start"]