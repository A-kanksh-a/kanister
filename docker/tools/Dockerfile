FROM registry.access.redhat.com/ubi8/ubi-minimal:8.1
ARG kan_tools_version="test-version"
LABEL name="kanister-tools" \
    vendor="Kanister" \
    version="${kan_tools_version}" \
    release="${kan_tools_version}" \
    summary="Microservice for application-specific data management for Kubernetes" \
    maintainer="Tom Manville<tom@kasten.io>" \
    description="Kanister tools for application-specific data management"

ENV GODEBUG=asyncpreemptoff=1

COPY --from=restic/restic:0.11.0 /usr/bin/restic /usr/local/bin/restic
# kopia alpine-1a0875e image
COPY --from=kanisterio/kopia@sha256:6e331f8cc9079062078c262112b6644ab108596f9ff82d552cec67b1a218e1c0 \
  /kopia/kopia /usr/local/bin/kopia
COPY LICENSE /licenses/LICENSE

ADD kando /usr/local/bin/
RUN microdnf update && microdnf install shadow-utils httpd-tools && \
  adduser -U kanister -u 1000 && \
  microdnf remove shadow-utils && \
  microdnf clean all

CMD [ "/usr/bin/tail", "-f", "/dev/null" ]
