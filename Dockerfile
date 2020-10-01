FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/authenticator

FROM scratch
COPY --from=builder /app/authenticator /app/
ENTRYPOINT ["/app/authenticator"]