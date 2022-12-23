FROM debian:bullseye
RUN apt-get update && apt-get install -y ca-certificates
COPY api /usr/bin/api
ENTRYPOINT ["/usr/bin/api"]
EXPOSE 3000