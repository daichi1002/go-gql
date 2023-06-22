FROM golang:latest

ENV ROOT=/go/src
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

# RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod tidy
EXPOSE 8080

CMD ["go", "run", "main.go"]