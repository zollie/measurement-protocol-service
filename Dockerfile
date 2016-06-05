FROM centos

ADD measurement-protocol-service  measurement-protocol-service

ADD run.sh run.sh

EXPOSE 2020

ENTRYPOINT ["/measurement-protocol-service"]
