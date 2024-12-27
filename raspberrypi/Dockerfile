FROM debian

ARG RASPBERRYPI_VERSION=3

RUN apt-get clean && apt-get update
RUN apt-get install python3-pip python3.11-venv git -y

RUN python3 -m venv pi-venv

# DHT11
RUN pi-venv/bin/python3 -m pip install --upgrade pip setuptools wheel

RUN git clone https://github.com/torchiaf/Adafruit_Python_DHT.git

RUN cd Adafruit_Python_DHT && /pi-venv/bin/python3 setup.py install --force-pi${RASPBERRYPI_VERSION}
RUN /pi-venv/bin/pip3 install ./Adafruit_Python_DHT --config-settings="--build-option=--force-pi${RASPBERRYPI_VERSION}"
RUN rm -rf /Adafruit_Python_DHT

# RabbitMQ
RUN pi-venv/bin/python -m pip install pika --upgrade

ADD ./rpc_server /rpc_server

WORKDIR /rpc_server

ENV RASPBERRYPI_VERSION=${RASPBERRYPI_VERSION}
ENV DHT11_PIN=12

CMD ["/pi-venv/bin/python3","init.py"]
