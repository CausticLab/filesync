FROM busybox:ubuntu-14.04
MAINTAINER CausticLab

ENV FILESYNC_RELEASE v0.0.1

ADD https://github.com/CausticLab/filesync/releases/download/${FILESYNC_RELEASE}/filesync-linux-amd64.tar.gz /tmp/filesync.tar.gz
RUN mkdir -p /usr/local/bin \
  && tar -zxvf /tmp/filesync.tar.gz -C /usr/local/bin \
  && mv /usr/local/bin/filesync-linux-amd64 /usr/local/bin/filesync \
  && chmod +x /usr/local/bin/filesync \
  && rm /tmp/* \
  && mkdir /share

ENTRYPOINT ["/usr/local/bin/filesync"]
