# vessel-service/Dockerfile
FROM golang:1.11.5 as builder

WORKDIR /go/src/
COPY ./github.com/ ./github.com/
COPY ./golang.org/ ./golang.org/
COPY ./gopkg.in/ ./gopkg.in/
COPY ./micros/ ./micros/
#RUN ls -la ./*
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init && dep ensur
WORKDIR /go/src/micros/vessel-service
#RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o vessel-service -a -installsuffix cgo .

FROM scratch
COPY --from=builder /go/src/micros/vessel-service/vessel-service /app/vessel-service
WORKDIR /app

CMD ["./vessel-service"]
