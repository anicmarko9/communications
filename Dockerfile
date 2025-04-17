## Builder
FROM golang:1.24.2-alpine AS builder

WORKDIR /home/app

COPY go.mod go.sum ./

RUN go mod vendor

COPY . .

RUN go build -mod=vendor -o main.exe ./cmd

## Runner
FROM alpine:3.21.3

WORKDIR /home/app

ARG PORT \
  GIN_MODE

ENV PORT=$PORT \
  GIN_MODE=$GIN_MODE

RUN apk --no-cache add ca-certificates

COPY --from=builder /home/app/main.exe ./main.exe

RUN chmod +x ./main.exe

EXPOSE $PORT

CMD ["./main.exe"]

# docker build --build-arg PORT="5000" `
# --build-arg GIN_MODE="release" `
# -t communications:go_1.24.2 --no-cache .

# docker run -d -p 5000:5000 communications:go_1.24.2
