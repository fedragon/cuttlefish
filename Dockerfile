FROM amd64/alpine:3.12
RUN apk add --no-cache fish git openssh

RUN adduser swen -h /home/swen -s /usr/bin/fish -D
USER swen
WORKDIR /home/swen

RUN mkdir -p .config/fish/functions

ADD tests/fish-switchenv /usr/bin/
ADD tests/__switchenv.fish .config/fish/functions/

ENV TMPDIR /tmp

RUN echo "functions -c __switchenv switchenv" > .config/fish/config.fish

ENTRYPOINT [ "ssh-agent", "fish" ]