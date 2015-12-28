FROM busybox

COPY ./bin/project /project

ENTRYPOINT ["./project"]
