FROM golang:1.14-alpine AS builder

WORKDIR /tmp/build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/msa-containerapp-go-itemapi .

FROM scratch

COPY --from=builder /app/msa-containerapp-go-itemapi /bin/app

EXPOSE 8080

ENTRYPOINT ["/bin/app"]
