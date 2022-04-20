FROM golang:1.18.1-alpine3.15 as gobuild

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /build

COPY main.go .

RUN go env

RUN go build -o test-goapp


# ---
FROM scratch AS run

COPY --from=build /build/test-goapp /bin/

CMD ["test-goapp"]
