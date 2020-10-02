FROM ubuntu:latest

RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get -y install wget vim python

WORKDIR /tmp

RUN wget -q https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-312.0.0-linux-x86_64.tar.gz
RUN tar -xvf google-cloud-sdk-312.0.0-linux-x86_64.tar.gz
RUN ./google-cloud-sdk/install.sh
CMD ["./google-cloud-sdk/bin/gcloud","init", "--console-only"]
# RUN gcloud components update
# RUN gcloud components install app-engine-go

