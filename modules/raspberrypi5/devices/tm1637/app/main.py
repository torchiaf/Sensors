import sys
import tm1637
from config import print, clk, dio

tm = tm1637.TM1637(clk=clk, dio=dio)

try:
    print('Received: {}'.format(sys.argv))

    action = sys.argv[1]
    fn = getattr(tm, sys.argv[2])
    arg = sys.argv[3]
    
    t = int(arg)
    fn(t) # show temperature 't*C'
except Exception as error:
    print(error)
    pass
# finally:
    # tm.write([0, 0, 0, 0])