FROM scratch
ADD ./bin/linux/amd64/pokevision /
ENTRYPOINT ["/pokevision"]
