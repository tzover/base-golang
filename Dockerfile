FROM golang:1.17.5-buster
RUN apt-get update && apt update 
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    ROOT=/go/src/app \
    GO111MODULE=on
WORKDIR ${ROOT}
# COPY ./app/go.mod ./app/go.sum ./
# RUN go mod download
EXPOSE 8080

# CMD ["go", "run", "main.go"]