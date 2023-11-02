# Build stage
FROM golang:1.21.3-alpine3.18 AS builder
LABEL org.opencontainers.image.authors="f.troia@davinci.care"

WORKDIR /app
ADD . .

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
EXPOSE 5000
CMD ["/app/proxyservice"]
