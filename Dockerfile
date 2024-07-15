FROM golang:1.20-alpine3.19

LABEL creator="qywok"

ENV APP_DIR=/tebakaja_proxy \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    DOTENV=.env \
    HOST=0.0.0.0 \
    PORT=7860

WORKDIR ${APP_DIR}

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main . && \
    go clean -modcache && \
    rm -rf /var/cache/apk/* \
      /root/.cache/go-build /root/go/pkg

RUN cat > ${DOTENV} <<EOF
TEBAKAJA_PROXY_HOST=${HOST}
TEBAKAJA_PROXY_PORT=${PORT}

TEBAKAJA_CORS_ALLOW_ORIGINS=https://huggingface.co,https://qywok-tebakaja-proxy-space-0.hf.space,https://qywok-tebakaja-proxy-space-1.hf.space,https://qywok-tebakaja-proxy-space-2.hf.space,https://qywok-tebakaja-proxy-space-3.hf.space,https://qywok-tebakaja-proxy-space-4.hf.space
TEBAKAJA_CORS_ALLOW_HEADERS=*
TEBAKAJA_CORS_ALLOW_METHODS=GET,POST
EOF

EXPOSE ${PORT}

CMD ["./main"]
