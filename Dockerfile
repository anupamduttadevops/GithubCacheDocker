

FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["./myapp"]




# FROM golang:1.22.3 AS builder

# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# FROM alpine:latest

# WORKDIR /root/

# COPY --from=builder /app/myapp .

# EXPOSE 8080

# CMD ["./myapp"]


