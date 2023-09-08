FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o suitmedia-go .

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/suitmedia-go .

ENTRYPOINT [ "./suitmedia-go" ]
EXPOSE 8080
