import os
import json
import Adafruit_DHT
from config import device

os.environ['RASPBERRYPI_VERSION'] = device.getConfig("RASPBERRYPI_VERSION")

sensor = Adafruit_DHT.DHT11
pin = int(device.getConfig("DHT11_PIN"))

try:
    humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
    
    d = dict();
    d['t'] = temperature
    d['h'] = humidity
    
    print(json.dumps(d))

except:
    print("read() error")
