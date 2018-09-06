FROM golang:1.10.1
LABEL Maintainer="Carlos Guzmán <cguzman@xmltravelgate.com>"
LABEL Description="golang" Vendor="XML Travelgate" Version="1.0.0"

# ENV VARIABLES
ENV DEPLOY_MODE="prod"
ENV SRC_PATH=/go/src/changelog/

# COPY SRC PROJECT
COPY . $SRC_PATH

# INSTALL DEPENDENCY MANAGEMENT FOR GO (DEP)  
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# UPDATE PROJECT DEPENDENCIES
RUN cd $SRC_PATH \
    && dep ensure -update && dep ensure -vendor-only

# BUILD PROJECT AND LAUNCH
RUN cd $SRC_PATH \
    && go build -o /go/bin/changelog

EXPOSE 8350

CMD ["bash", "-c", "/go/bin/changelog"]
