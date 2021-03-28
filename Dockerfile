FROM golang:1.13
RUN go get golang.org/x/tools/cmd/present
COPY . $GOPATH/src/golang.org/x/tools/cmd/present/content
WORKDIR $GOPATH/src/golang.org/x/tools/cmd/present/
ENV GAE_ENV "standard"
CMD ["present", "-notes"]
