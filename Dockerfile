FROM scratch
EXPOSE 80
COPY ./teapot /teapot 
ENTRYPOINT ["/teapot"]
