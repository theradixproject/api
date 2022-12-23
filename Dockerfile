FROM debian:bullseye
COPY api /usr/bin/api
ENTRYPOINT ["/usr/bin/api"]
EXPOSE 3000