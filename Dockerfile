# syntax = docker/dockerfile:1.4

FROM golang:1.21.3-bookworm as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /rss_service


FROM build as release
ENV PORT="2345"
WORKDIR /

COPY --from=build /rss_service /rss_service

RUN groupadd -f -g 1000 rss && \
    useradd -u 1000 -g 1000 rss --create-home --shell /bin/bash && \
    mkdir /data && \
    chown -R rss:rss /data
USER rss:rss

EXPOSE ${PORT}

ENTRYPOINT ["/rss_service"]

# CMD ["/rss_service"]

