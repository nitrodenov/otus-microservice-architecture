FROM golang:1.16
WORKDIR ~/GoWorkspace/src/otus
COPY *.go .
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go build Otus2.go
EXPOSE 8000
CMD ["./Otus2"]