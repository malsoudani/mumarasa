import backtrader as bt

class PrintClose(bt.Strategy):
    def __init__(self):
        self.dataclose = self.datas[0].close
    
    def log(self, text, data):
        dt = data.datetime.date(0).isoformat()
        print(f'{dt}, {text} {data.close[0]}')

    def next(self):
        self.log('close: ', self.datas[0])