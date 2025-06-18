FROM fedora:latest AS final
WORKDIR /app

RUN dnf update -y && \
    dnf install -y \
    curl \
    ffmpeg \
    unzip \
    go \
    && dnf clean all

RUN curl -fsSL https://bun.sh/install | bash
ENV PATH="/root/.bun/bin:${PATH}"

RUN go install github.com/a-h/templ/cmd/templ@latest
ENV PATH="/root/go/bin:${PATH}"

COPY package.json bun.lock ./
RUN bun install

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

RUN templ generate
RUN bunx tailwindcss -i style.css -o include_dir/output.css -m

EXPOSE 8080

RUN go build -o /bin/server

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]


