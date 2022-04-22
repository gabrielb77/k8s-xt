FROM golang:1.18.1-alpine3.15 as gobuild

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=auto
#ENV GOPATH=/go

#WORKDIR /build

#RUN ls -lh

COPY . .

RUN ls -lh
RUN go env
RUN pwd
#RUN go mod tidy
#RUN go get github.com/sirupsen/logrus
#RUN go get github.com/slackhq/simple-kubernetes-webhook/pkg/admission#
#RUN go get

#RUN go run main.go

RUN go build -o test-goapp app/main.go


# ---
FROM scratch AS run

COPY --from=gobuild /go/test-goapp /bin/

CMD ["test-goapp"]
