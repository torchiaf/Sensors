import Adafruit_DHT
from config import module

device = module.getDevice("dht11")

def read():
    sensor = Adafruit_DHT.DHT11
    pin = int(device.getConfig("DHT11_PIN"))

    try:
        humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
        
        d = dict();
        d['t'] = temperature
        d['h'] = humidity
        
        return d

    except:
        print("read() error")
