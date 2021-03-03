FROM golang:1.15.6-alpine3.12 as builder

COPY . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o host2ipv4 .

FROM scratch
COPY --from=builder /app/host2ipv4 /host2ipv4
ENTRYPOINT [ "/host2ipv4" ]
