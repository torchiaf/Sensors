import os
from time import sleep

import Adafruit_DHT

def read(pin):
    sensor = Adafruit_DHT.DHT11
    pin = int(os.environ.get('DHT11_PIN'))

    try:
        humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
        
        d = dict();
        d['t'] = temperature
        d['h'] = humidity
        
        return d

    except:
        print("read() error")
