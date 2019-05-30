FROM golang:1.12 AS modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.12 as builder


RUN useradd -u 10001 tenerife
COPY --from=modules /go/pkg/mod /go/pkg/mod

RUN mkdir -p /tenerife
ADD . /tenerife
WORKDIR /tenerife

RUN make build

# Final stage: Run the binary
FROM scratch AS Runner

ENV PORT 8090

# certificates to interact with other services
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# don't forget /etc/passwd from previous stage
COPY --from=builder /etc/passwd /etc/passwd
USER tenerife

# and finally the binary
COPY --from=builder /tenerife/bin/tenerife /tenerife
EXPOSE $PORT

CMD ["tenerife"]
