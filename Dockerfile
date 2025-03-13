FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

COPY . .

RUN mkdir -p ~/.ssh
RUN --mount=type=secret,id=mysecret cat /run/secrets/mysecret > ~/.ssh/id_rsa
RUN chmod 600 ~/.ssh/id_rsa

RUN apk add --no-cache --update; \
    apk add git openssh; \
    apk add tzdata;

RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"; \
    export GOPRIVATE=gitlab.com/nbdgocean6; \
    export GONOPROXY=gitlab.com/nbdgocean6;   \
    export GONOSUMDB=gitlab.com/nbdgocean6; \
    ssh-keyscan gitlab.com >> /root/.ssh/known_hosts; \
    go mod tidy; \
    go build -ldflags "-s -w" -o nobita-promo-program main.go

#Final Build Image
FROM alpine:latest

WORKDIR /app

ENV TZ=Asia/Jakarta

RUN mkdir -p /app/config

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /build/nobita-promo-program /app/nobita-promo-program

ENTRYPOINT ["/app/nobita-promo-program"]
