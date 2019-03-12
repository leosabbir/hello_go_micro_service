FROM golang:1.11.5 as builder

WORKDIR /go/src/
COPY ./github.com/ ./github.com/
COPY ./golang.org/ ./golang.org/
COPY ./gopkg.in/ ./gopkg.in/
COPY ./micros/ ./micros/

WORKDIR /go/src/micros/consignment-service
#RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consignment-service -a -installsuffix cgo .

FROM scratch
COPY --from=builder /go/src/micros/consignment-service/consignment-service /app/consignment-service
WORKDIR /app

CMD ["./consignment-service"]
