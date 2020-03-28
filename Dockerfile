FROM golang
ENV SRC_DIR=/go/src
WORKDIR $SRC_DIR
ENV GOPATH $SRC_DIR
COPY src/ ./src/
RUN ls -R
RUN go build ./src/hello
ENTRYPOINT [ "./hello" ]
