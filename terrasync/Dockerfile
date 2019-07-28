FROM golang:1.12 as builder

COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 go build -o terrasync .

FROM hashicorp/terraform:0.12.5

COPY --from=builder /src/terrasync /bin/terrasync
ENTRYPOINT ["terrasync"]
