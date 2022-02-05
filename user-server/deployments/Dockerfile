FROM golang:1.17-alpine AS build-env
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o user-server ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=build-env /app/user-server ./
COPY configs ./configs
ENTRYPOINT ./user-server