# syntax=docker/dockerfile:1.4
# Build stage
FROM golang:1.23.4-alpine AS build

WORKDIR /app

# Install build toolchain for cgo (gcc, musl-dev, make)
RUN apk add --no-cache build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Enable cgo so WEBP conversion works
ENV CGO_ENABLED=1 GOOS=linux
RUN go build -o /app/main ./api

# Final stage
FROM alpine:latest

ARG PORT=8080

WORKDIR /app

# Minimal runtime deps for cgo binaries
RUN apk add --no-cache ca-certificates tzdata libgcc

COPY --from=build /app/main .
COPY config ./config

# Ensure files directory exists (mounted as volume in compose)
RUN mkdir -p /app/files

ENV PORT=${PORT}
EXPOSE ${PORT}

ENTRYPOINT ["./main"]
