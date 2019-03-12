FROM golang:1.11.5 as builder

WORKDIR /go/src/
COPY ./github.com/ ./github.com/
COPY ./golang.org/ ./golang.org/
COPY ./micros/ ./micros/
WORKDIR /go/src/micros/user-cli
#RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-cli -a -installsuffix cgo .

FROM scratch
COPY --from=builder /go/src/micros/user-cli/user-cli /app/user-cli
WORKDIR /app

CMD ["./user-cli"]
