FROM zenika/alpine-golang
EXPOSE 4000
COPY teapot /teapot
CMD ["/teapot"]
