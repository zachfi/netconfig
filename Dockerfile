FROM alpine:3.15 as certs
COPY ./netconfig /bin/netconfig
RUN chmod 0755 /bin/netconfig
RUN mkdir /var/netconfig
RUN apk --update add ca-certificates
RUN apk add libc6-compat
RUN apk add tzdata
ENTRYPOINT ["/bin/netconfig"]
