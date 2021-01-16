FROM golang
RUN apt-get -y update

RUN git clone https://github.com/J0ey1iu/HushTell.git && cd HushTell && git checkout Vanilla
WORKDIR "/go/HushTell"

CMD ["./docker_start_script.sh"]