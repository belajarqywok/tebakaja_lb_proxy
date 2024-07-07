FROM golang:1.20

LABEL creator="qywok"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 7860

CMD ["./main"]
