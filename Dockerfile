FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN swag init -g router/swaggerRouter.go --outputTypes go --requiredByDefault
RUN go build -o /app/server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/mail/templates ./mail/templates
EXPOSE 8080
CMD ["/app/server"]