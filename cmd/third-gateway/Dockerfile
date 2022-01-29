FROM jrei/systemd-centos:7

COPY ThirdGateway.viper.yaml /etc/ThirdGateway/ThirdGateway.viper.yaml
COPY third-gateway /usr/local/bin/third-gateway

CMD ["third-gateway", "run"]
