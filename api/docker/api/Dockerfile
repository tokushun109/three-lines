FROM golang:1.18

WORKDIR /api
COPY ./ /api

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air", "-c", ".air.toml"]