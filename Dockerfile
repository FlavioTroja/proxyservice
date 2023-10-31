# Build stage
FROM golang:1.18.3-alpine3.16 AS builder
LABEL org.opencontainers.image.authors="f.troia@davinci.care"

WORKDIR /app
ADD . .

# DAVIDE 20220203 - applicato fix per errore build
RUN go mod download ;\
	go build -o proxyservice || go get github.com/prometheus/common/log@v0.26.0 ;\
	test -f proxyservice || go build -o proxyservice

# Production stage
FROM alpine:3.14.1

# Adjust timezone
RUN apk add ca-certificates tzdata \
	&& cp /usr/share/zoneinfo/Europe/Rome /etc/localtime \
	&& echo "Europe/Rome" > /etc/timezone \
    && apk del tzdata

WORKDIR /app
COPY --from=builder /app/proxyservice .

CMD ["/app/proxyservice"]
