FROM golang
ADD . /go/src/servertime
RUN cd /go/src/servertime && \
    go install servertime
ENTRYPOINT /go/bin/servertime
EXPOSE 8080
