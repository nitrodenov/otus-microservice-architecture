FROM golang:1.16
WORKDIR ~/GoWorkspace/src/Otus1
COPY *.go .
RUN go build Otus1.go
EXPOSE 8000
CMD ["./Otus1"]