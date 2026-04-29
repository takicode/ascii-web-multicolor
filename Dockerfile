FROM golang:1.22.2-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o myapp .

FROM alpine:latest

LABEL maintainer="Taki Pelumi Emmanuel"
LABEL version="1.0"
LABEL description="Ascii web multicolor"

WORKDIR /app
COPY --from=builder /app/myapp .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/banner ./banner
EXPOSE 8080
CMD ["./myapp"]

