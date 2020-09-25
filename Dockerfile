FROM golang:1.15

RUN apt-get update && apt-get upgrade
RUN apt-get install -y git supervisor 
RUN git clone --depth 1 --branch portalv3 https://github.com/incognitochain/incognito-chain
RUN git clone https://github.com/incognitochain/incognito-highway && \
    cd incognito-highway && \
    git checkout 20200629_3 && \
    go build -o highway

COPY supervisord.conf supervisord.conf
COPY run.sh run.sh
COPY blockchain.txt ./incognito-chain/blockchain/blockchain.go
RUN chmod a+x run.sh
RUN mkdir /var/log/supervisord

EXPOSE 9334 9338

ENTRYPOINT ["/go/run.sh"]
# CMD ["supervisord"]

