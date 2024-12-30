import board
import adafruit_dht
from config import module

device = module.getDevice("dht11")
pin = device.getConfig("DHT11_PIN")

# Initial the dht device, with data pin connected to:
_device = getattr(adafruit_dht, device.type)
_board = getattr(board, pin)
dhtDevice = _device(_board)

def read():
    d = dict();

    try:
        d['t'] = dhtDevice.temperature
        d['h'] = dhtDevice.humidity
        
        return d

    except Exception as error:
        dhtDevice.exit()
        print(error)
        return d
