FROM golang:1.18

WORKDIR /api
COPY ./ /api

RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]