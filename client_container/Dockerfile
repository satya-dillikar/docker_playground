FROM ubuntu:18.04

RUN apt-get update -y
RUN apt-get install wget -y
RUN apt-get install curl -y
RUN wget https://github.com/fullstorydev/grpcurl/releases/download/v1.1.0/grpcurl_1.1.0_linux_x86_64.tar.gz

RUN tar -xvzf grpcurl_1.1.0_linux_x86_64.tar.gz
RUN chmod +x grpcurl
RUN mv grpcurl /usr/local/bin/grpcurl
