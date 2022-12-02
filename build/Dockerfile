# Build stage
FROM chainberry/golang-buster:1.19.0 AS build-env
WORKDIR /src
COPY api/ api/
COPY cmd/ cmd/
COPY internal/ internal/
COPY Makefile Makefile
COPY go.mod go.mod
#ARG goproxy=direct
#ENV GOPROXY=$goproxy
#ENV CGO_ENABLED=0
#ENV PATH="$PATH:$(go env GOPATH)/bin"
RUN make build

## Production stage
FROM chainberry/golang-buster:1.19.0 as run-env
COPY --from=build-env /src/service /service/work_view
WORKDIR /service

CMD ["./work_view", "migrate", "up", "&&", "./work_view", "server" ]
