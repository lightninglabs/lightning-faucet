FROM golang:1.20-alpine

WORKDIR /app
RUN apk add --no-cache git
RUN git clone https://github.com/lightninglabs/lightning-faucet.git .
RUN go build -o lightning-faucet .
EXPOSE 8080

CMD ["./lightning-faucet"]
