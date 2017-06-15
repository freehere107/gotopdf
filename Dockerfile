FROM daocloud.io/library/golang:1.7

RUN \
	apt-get update && \
	DEBIAN_FRONTEND=noninteractive \
	    apt-get upgrade -y && \
		apt-get install -y \
		        locales \
			unoconv \
			gcc \
			supervisor


COPY src/unoconv/unoconv /opt

RUN \
    apt-get install -y xfonts-utils

CMD chmod +x /opt/unoconv

CMD ["/opt/unoconv"]

EXPOSE 1323