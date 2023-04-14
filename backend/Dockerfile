FROM golang as base

FROM base as dev

RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v golang.org/x/tools/cmd/goimports@latest