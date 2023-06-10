# ./Dockerfile
FROM golang:1.20.5-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download

# build version for linux (docker builder os)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix togo -o main ./cmd/api/

FROM golang:1.20.5-alpine

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

ENTRYPOINT ["./main"]
