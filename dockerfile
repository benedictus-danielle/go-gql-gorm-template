FROM golang:1.16-alpine

WORKDIR /go/src/github.com/benedictus-danielle/go-gql-template-project
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go-gql-template-project
EXPOSE 8000
CMD ["/go-gql-template-project"]