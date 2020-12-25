import backtrader as bt

class MyStrategy(bt.Strategy):
    def next(self):
        pass # Do Nothing


#instiate the cerebro engine
cerebro = bt.Cerebro()

data = bt.feeds.YahooFinanceCSVData(dataname='backtrader/AMZN.csv')

cerebro.adddata(data)

# Add strategy to cerebro 
cerebro.addstrategy(MyStrategy)

#Run cerebro
cerebro.run()