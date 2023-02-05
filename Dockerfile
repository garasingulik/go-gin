# Build
FROM golang:1.20 AS builder
ARG APP_NAME=app

WORKDIR /app
COPY . /app

# Test
WORKDIR /app
RUN go mod vendor && go test -run ''

# Package
WORKDIR /app
RUN go mod vendor
RUN go build -o /app/$APP_NAME . && chmod 755 /app/$APP_NAME

FROM golang:1.20
ARG APP_NAME=app

WORKDIR /app

COPY --from=builder /app/$APP_NAME  /app/$APP_NAME
COPY --from=builder /app/config/ /app/config/

EXPOSE 80
CMD ["/app/$APP_NAME"]
