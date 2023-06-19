FROM golang:1.20

## UPDATE THE OS
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y  && \
    apt-get install -y tzdata  

WORKDIR /go/src

## SET ENVIRONMENT
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on
ENV TZ America/Sao_Paulo

## COPY NECESSARY FILES
COPY go.mod go.* ./
RUN go mod download && \
    go mod tidy
COPY . ./

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]
