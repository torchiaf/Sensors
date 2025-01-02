FROM python:3.9.21-alpine

ARG RASPBERRYPI_VERSION=3

ADD ./app /app

RUN apk update
RUN apk add git build-base

RUN python3 -m venv pi-venv
RUN pi-venv/bin/python3 -m pip install --upgrade pip
RUN pi-venv/bin/pip install -U pyyaml pyinstaller

# DHT11
RUN pi-venv/bin/python3 -m pip install --upgrade pip setuptools wheel

RUN git clone https://github.com/torchiaf/Adafruit_Python_DHT.git

RUN cd Adafruit_Python_DHT && /pi-venv/bin/python3 setup.py install --force-pi${RASPBERRYPI_VERSION}
RUN /pi-venv/bin/pip3 install ./Adafruit_Python_DHT --config-settings="--build-option=--force-pi${RASPBERRYPI_VERSION}"

# RabbitMQ
RUN pi-venv/bin/python -m pip install pika --upgrade

# Build init.py
RUN pi-venv/bin/pyinstaller --onefile /app/init.py
