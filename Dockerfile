FROM scratch
EXPOSE 4000
COPY ./teapot /teapot
ENTRYPOINT ["/teapot"]
