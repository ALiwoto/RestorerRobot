FROM golang:1.23.1-alpine3.20 AS builder
WORKDIR /build
COPY . .
RUN apk add --no-cache ca-certificates gcc musl-dev \
    && go mod tidy \
    && CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o /build/restorerrobot .

FROM alpine:3.20
WORKDIR /app
RUN apk add --no-cache bash ca-certificates curl postgresql-client tzdata
COPY --from=builder /build/restorerrobot .
CMD [ "./restorerrobot" ]
