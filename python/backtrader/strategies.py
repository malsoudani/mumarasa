import backtrader as bt

class PrintClose(bt.Strategy):
    def __init__(self):
        self.dataclose = self.datas[0].close
    
    def log(self, text, data):
        dt = data.datetime.date(0).isoformat()
        print(f'{dt}, {text} {data.close[0]}')

    def next(self):
        self.log('close: ', self.datas[0])

class MACrossover(bt.Strategy):
    params = (('pfast', 20), ('pslow', 50),)

    def log(self, txt, dt=None):
        dt = dt or self.datas[0].datetime.date(0)
        print(f'{dt.isoformat()} {txt}')
    
    def __init__(self):
        self.dataclose = self.datas[0].close
        self.order = None
        self.slow_sma = bt.indicators.MovingAverageSimple(self.datas[0], period=self.p.pslow)
        self.fast_sma = bt.indicators.MovingAverageSimple(self.datas[0], period=self.p.pfast)

    def notify_order(self, order):
        if order.status in [order.Submitted, order.Accepted]:
            # an active buy / sell... there is nothing to do
            return

        if order.status in [order.Completed]:
            if order.isbuy():
                self.log(f'buy excuted {order.executed.price:.2f}')
            elif order.issell():
                self.log(f'sell excuted {order.executed.price:.2f}')
            self.bar_executed = len(self)
        
        elif order.status in [order.Margin, order.Canceled, order.Rejected]:
            self.log('order canceled/rejected/margin')

        # reset the order
        self.order = None

    def next(self):
        # check if there is an open order
        if self.order:
            return
        
        # we don't have an open order, great... 
        # lets check if we are already holding a position in the market
        if not self.position:
            # we have no open positions right now..
            # if the fast period just crossed above the slow period
            if self.fast_sma[0] > self.slow_sma[0] and self.fast_sma[-1] < self.slow_sma[-1]:
                self.log(f'just bought: {self.dataclose[0]:2f}')
                # set the order variable so that we don't run through the position again.
                self.order = self.buy()
            # otherwise if the fast period just crossed under the slow period
            elif self.fast_sma[0] < self.slow_sma[0] and self.fast_sma[-1] > self.slow_sma[-1]:
                self.log(f'just sold: {self.dataclose[0]:2f}')
                # set the order variable so that we don't run through the position again.
                self.order = self.sell()
        else:
            # we are already holding a position so we need 
            # to look for a signal to close out...
            if len(self) > (self.bar_executed + 5):
                self.log(f'just closed {self.dataclose[0]:2f}')
                self.order = self.close()
