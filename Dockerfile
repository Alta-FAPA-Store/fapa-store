FROM golang:1.16-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main ./app/
# EXPOSE 8080
# CMD ["/app/main"]

# stage 2
FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/.env .
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]