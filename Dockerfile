FROM skyzyx/alpine-pandoc:1.2.0


WORKDIR /src
ADD . /src
RUN cd /src && go build -o app
EXPOSE 8095
ENTRYPOINT ./app