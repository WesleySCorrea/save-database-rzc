
FROM golang:1.20.10-bullseye


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN go build -v -o ./ ./app


EXPOSE 8080

CMD ["/usr/local/bin/app/project"]