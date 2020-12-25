import backtrader as bt

class PrintClose(bt.Strategy):
    def __init__(self):
        self.dataclose = self.datas[0].close
    
    def log(self, text, data):
        dt = data.datetime.date(0).isoformat()
        print(f'{dt}, {text} {data.close[0]}')

    def next(self):
        self.log('close: ', self.datas[0])


#instiate the cerebro engine
cerebro = bt.Cerebro()

data = bt.feeds.YahooFinanceCSVData(dataname='backtrader/AMZN.csv')

cerebro.adddata(data)

# Add strategy to cerebro 
cerebro.addstrategy(PrintClose)

#Run cerebro
cerebro.run()