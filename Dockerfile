FROM golang:1.15.6-alpine3.12 as builder

COPY . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o host2value .

FROM scratch
COPY --from=builder /app/host2value /host2value
ENTRYPOINT [ "/host2value" ]
