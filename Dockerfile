FROM alpine:3.6

COPY bin/app /bin/app

RUN chmod +x /bin/app

ENTRYPOINT ["/bin/app"]