FROM lacion/docker-alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/divergentdave/cidr-test-service"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/cidr-test-service/bin

WORKDIR /opt/cidr-test-service/bin

COPY bin/cidr-test-service /opt/cidr-test-service/bin/
RUN chmod +x /opt/cidr-test-service/bin/cidr-test-service

CMD /opt/cidr-test-service/bin/cidr-test-service