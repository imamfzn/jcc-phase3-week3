FROM alpine:3.14

RUN apk --no-cache add ca-certificates
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.30-r0/glibc-2.30-r0.apk && \
    apk --no-cache add glibc-2.30-r0.apk && \
    rm glibc-2.30-r0.apk
COPY _output/gosample /usr/local/bin/gosample

CMD ["gosample"]
