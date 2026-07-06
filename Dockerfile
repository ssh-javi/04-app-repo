FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/main /main
EXPOSE 8080
CMD ["/main"]
