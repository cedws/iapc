FROM cgr.dev/chainguard/static:latest
ARG TARGETPLATFORM
ARG PROJECT_NAME
COPY ${TARGETPLATFORM}/${PROJECT_NAME} /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
