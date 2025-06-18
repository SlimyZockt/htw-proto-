# syntax=docker/dockerfile:1
FROM alpine:3.21 AS final
WORKDIR /app

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
    curl \
    tzdata \
    ffmpeg \
    sqlite \
    go \
    nodejs \
    pnpm

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add templ --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/

RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh 

COPY package.json package-lock.json ./
RUN pnpm install

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

RUN templ generate
RUN pnpm exec tailwindcss -o include_dir/output.css -m
RUN goose -dir=assets/migrations/ sqlite3 app.db up

# Expose the port that the application listens on.
EXPOSE 8080

RUN GOOS=linux go build -o /bin/server

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]

