## Builder
FROM golang:1.24.2-alpine AS builder

WORKDIR /home/app

COPY go.mod go.sum ./

RUN go mod vendor

COPY . .

RUN go test ./...

RUN go build -mod=vendor -o main.exe ./cmd

## Runner
FROM alpine:3.21.3

WORKDIR /home/app

ARG PORT \
  THROTTLE_TTL \
  THROTTLE_LIMIT \
  GIN_MODE \
  ALLOWED_ORIGINS \
  POSTGRES_HOST \
  POSTGRES_PORT \
  POSTGRES_USER \
  POSTGRES_PASSWORD \
  POSTGRES_DB \
  POSTGRES_SSL \
  AZURE_URL \
  EMAIL_FROM \
  SMS_FROM

ENV PORT=$PORT \
  THROTTLE_TTL=$THROTTLE_TTL \
  THROTTLE_LIMIT=$THROTTLE_LIMIT \
  GIN_MODE=$GIN_MODE \
  ALLOWED_ORIGINS=$ALLOWED_ORIGINS \
  POSTGRES_HOST=$POSTGRES_HOST \
  POSTGRES_PORT=$POSTGRES_PORT \
  POSTGRES_USER=$POSTGRES_USER \
  POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  POSTGRES_DB=$POSTGRES_DB \
  POSTGRES_SSL=$POSTGRES_SSL \
  AZURE_URL=$AZURE_URL \
  EMAIL_FROM=$EMAIL_FROM \
  SMS_FROM=$SMS_FROM

RUN apk --no-cache add ca-certificates openssh-server && \
  mkdir -p /root/.ssh && \
  ssh-keygen -A

COPY --from=builder /home/app/ssh/authorized_keys /root/.ssh/authorized_keys
COPY --from=builder /home/app/ssh/sshd_config /etc/ssh/.
COPY --from=builder /home/app/main.exe ./main.exe
COPY --from=builder /home/app/migrations ./migrations

RUN chmod +x ./main.exe && \
  chmod u=rwx /root/.ssh && \
  chmod u=rw /root/.ssh/authorized_keys

EXPOSE $PORT 2222

CMD ["/bin/sh", "-c", "/usr/sbin/sshd -D & ./main.exe"]


# docker build -t go-communications:1.24.2-alpine --no-cache . `
# --build-arg PORT="5000" `
# --build-arg THROTTLE_TTL="60" `
# --build-arg THROTTLE_LIMIT="10" `
# --build-arg GIN_MODE="debug" `
# --build-arg ALLOWED_ORIGINS="http://localhost:3000" `
# --build-arg POSTGRES_HOST="localhost" `
# --build-arg POSTGRES_PORT="5432" `
# --build-arg POSTGRES_USER="postgres" `
# --build-arg POSTGRES_PASSWORD="asdfghjkl123" `
# --build-arg POSTGRES_DB="local_db" `
# --build-arg POSTGRES_SSL="false"
# --build-arg AZURE_URL=""
# --build-arg EMAIL_FROM=""
# --build-arg SMS_FROM=""

# docker run -d -p 5000:5000 -p 2222:2222 go-communications:1.24.2-alpine
# ssh -i "ssh/id_rsa" -p 2222 root@localhost
