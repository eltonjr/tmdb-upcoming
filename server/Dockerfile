# ---- Base ----
FROM golang:alpine AS base
RUN apk add --update git
WORKDIR /app
COPY . .

# ---- Deps ----
FROM base AS dependencies
RUN GO111MODULE=on go get -insecure ./...

# ---- Build ----
FROM dependencies as build
RUN cd cmd/server && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -installsuffix cgo -o server

# ---- Release ----
FROM alpine:3.10 AS release
COPY --from=build /app/cmd/server/server /opt/

EXPOSE 9000

CMD ./opt/server
