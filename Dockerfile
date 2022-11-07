FROM golang:1.19.2-alpine3.16 as builder

WORKDIR /src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o url-to-qr main.go

FROM scratch

WORKDIR /bin/app

COPY --from=builder /src/app/url-to-qr .

CMD [ "./url-to-qr" ]