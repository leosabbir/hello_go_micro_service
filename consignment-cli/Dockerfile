FROM golang:1.11.5 as builder

WORKDIR /go/src/
COPY ./github.com/ ./github.com/
COPY ./golang.org/ ./golang.org/
COPY ./micros/ ./micros/
#RUN ls -la ./*
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init && dep ensur
WORKDIR /go/src/micros/consignment-cli
#RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consignment-cli -a -installsuffix cgo .

FROM scratch
COPY --from=builder /go/src/micros/consignment-cli/consignment-cli /app/consignment-cli
COPY --from=builder /go/src/micros/consignment-cli/consignment.json /app/consignment.json
WORKDIR /app

CMD ["./consignment-cli"]
