FROM golang:1.16.0 AS builder
WORKDIR /myappstatus/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myappstatus .

FROM alpine:3.13.2
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /myappstatus/kate .
ENTRYPOINT ["./myappstatus"] 