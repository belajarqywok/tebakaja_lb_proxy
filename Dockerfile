# --- Build stage ---
FROM golang:1.20 as builder

LABEL creator="qywok"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# --- Deploy stage ---
FROM alpine:latest

WORKDIR /app

RUN adduser -D -u 1000 user
USER user
ENV PATH="/home/user/.local/bin:$PATH"

COPY --from=builder /app/main /app/main

EXPOSE 7860

CMD ["./main"]
