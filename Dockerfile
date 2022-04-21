FROM golang:1.18.1-alpine3.15 as gobuild

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=auto
ENV GOPATH=/go

WORKDIR /build

COPY . .

RUN ls -lR pkg
RUN go env
# n
#RUN go get github.com/sirupsen/logrus
#RUN go get github.com/slackhq/simple-kubernetes-webhook/pkg/admission

RUN go build -o test-goapp .


# ---
FROM scratch AS run

COPY --from=gobuild /build/test-goapp /bin/

CMD ["test-goapp"]
