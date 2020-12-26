import backtrader as bt
import datetime
from strategies import *

#instiate the cerebro engine
cerebro = bt.Cerebro()

data = bt.feeds.YahooFinanceCSVData(
    dataname='backtrader/AMZN.csv', 
    fromdate=datetime.datetime(2019, 12, 24), 
    todate=datetime.datetime(2020, 3, 5),
)

cerebro.adddata(data)

# Add strategy to cerebro 
cerebro.addstrategy(PrintClose)

cerebro.addsizer(bt.sizer.SizerFix, stake=3)

if __name__ == '__main__':
    starting_protfolio_value = cerebro.broker.get_value()

    #Run cerebro
    cerebro.run()

    end_protfolio_value = cerebro.broker.get_value()

    profit = end_protfolio_value - starting_protfolio_value

    print(f'PnL: {profit:.2f}')