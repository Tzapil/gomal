FROM scratch
ADD bin/main /
VOLUME ["/etc/ssl/certs"]
CMD ["/main"]
EXPOSE 8080