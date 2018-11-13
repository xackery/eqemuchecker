FROM alpine:latest
RUN apk add --no-cache make perl lua perl-dbi perl-switch perl-dbd-mysql
COPY checker /app/checker
WORKDIR /home
CMD ["/app/checker"]