FROM golang:1.16 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/urecruit/book_service
WORKDIR $GOPATH/src/gitlab.udevs.io/urecruit/book_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/book_service /

FROM alpine
COPY --from=builder book_service .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent
ENTRYPOINT ["/book_service"]
