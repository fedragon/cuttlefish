FROM amd64/alpine:3.21

RUN apk add --no-cache fish git openssh

RUN adduser nemo -h /home/nemo -s /usr/bin/fish -D
USER nemo
WORKDIR /home/nemo

RUN mkdir -p .config/fish/functions

ADD tests/cuttlefish /usr/bin/
ADD tests/__cuttlefish.fish .config/fish/functions/

ENV TMPDIR=/tmp

RUN echo "functions -c __cuttlefish cuttlefish" > .config/fish/config.fish

ENTRYPOINT [ "ssh-agent", "fish" ]