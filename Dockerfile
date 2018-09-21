FROM scratch
EXPOSE 8080
ENTRYPOINT ["/meow"]
COPY ./bin/ /