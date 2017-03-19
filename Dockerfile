FROM alpine:edge
MAINTAINER CausticLab

ENV FILESYNC_RELEASE v0.0.1

ADD https://github.com/CausticLab/filesync/releases/download/${RGON_EXEC_RELEASE}/filesync-linux-amd64.tar.gz /tmp/filesync.tar.gz
RUN tar -zxvf /tmp/filesync -C /usr/local/bin \
  && mv /usr/local/bin/filesync-linux-amd64 /usr/local/bin/filesync \
  && chmod +x /usr/local/bin/filesync \
  && rm /tmp/filesync

ENTRYPOINT ["/usr/local/bin/filesync"]