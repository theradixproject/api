FROM debian:bullseye
COPY api /usr/bin/api
ENTRYPOINT ["/usr/bin/api"]