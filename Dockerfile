FROM skyzyx/alpine-pandoc:1.2.0


WORKDIR /
COPY app /
EXPOSE 8095
ENTRYPOINT ./app