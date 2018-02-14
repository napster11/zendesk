# workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
RUN mkdir -p /go/src/github.com/napster11/zendesk

COPY . /go/src/github.com/napster11/zendesk
WORKDIR /go/src/github.com/napster11/zendesk

RUN go-wrapper download
RUN go-wrapper install
ENTRYPOINT /go/bin/zendesk

# Service listen on port 8080.
EXPOSE 8080

