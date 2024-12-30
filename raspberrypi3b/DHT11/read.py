import sys
from time import sleep

import Adafruit_DHT

sensor = Adafruit_DHT.DHT11
pin = int(sys.argv[1])

try:
    while True:
        humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
        print("{{\"t\":{},\"h\":{}}}".format(temperature, humidity))
        
        sleep(1)

except KeyboardInterrupt:
    print("Exit")
