FROM golang:1.17-alpine AS build-env
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o consumer ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=build-env /app/consumer ./
COPY configs ./configs
ENTRYPOINT ./consumer