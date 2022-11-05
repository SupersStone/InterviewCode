FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /dao

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /dao/cmd/scanner/main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /dao/main .

CMD ["./main","--config","/root/config.d/config.yaml"]