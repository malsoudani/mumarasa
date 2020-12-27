import backtrader as bt
import datetime
from strategies import *

#instiate the cerebro engine
cerebro = bt.Cerebro(optreturn=False)

data = bt.feeds.YahooFinanceCSVData(
    dataname='backtrader/AMZN.csv', 
    fromdate=datetime.datetime(2016, 12, 24), 
    todate=datetime.datetime(2080, 6, 5),
)

cerebro.adddata(data)

# Add strategy to cerebro 
cerebro.addanalyzer(bt.analyzers.SharpeRatio, _name="sharpe_ratio")
cerebro.optstrategy(MACrossover, pfast=range(5, 20), pslow=range(50, 100))

#Default position size
cerebro.addsizer(bt.sizers.SizerFix, stake=1)

if __name__ == '__main__':
    # starting_protfolio_value = cerebro.broker.get_value()

    #Run cerebro
    optimized_runs = cerebro.run()

    final_results_list = []

    for run in optimized_runs:
        for strategy in run:
            profit = round(strategy.broker.get_value() - 10000, 2)
            sharpe = strategy.analyzers.sharpe_ratio.get_analysis()
            if sharpe['sharperatio']:
                final_results_list.append([
                    strategy.params.pfast, 
                    strategy.params.pslow, 
                    profit, 
                    sharpe['sharperatio']
                ])
    sort_by_sharpe = sorted(final_results_list, key=lambda x: x[3], reverse=True)
    for line in sort_by_sharpe[:5]: 
        print(line)


    # end_protfolio_value = cerebro.broker.get_value()

    # profit = end_protfolio_value - starting_protfolio_value

    # print(f'profits: {profit:.2f}')