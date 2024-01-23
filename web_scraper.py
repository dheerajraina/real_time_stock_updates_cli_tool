import sys
from nsetools import Nse
# from pprint import pprint

arg1 = sys.argv[1]


def get_stock_data(stock_sym):

    nse = Nse()
    response = nse.get_quote(stock_sym)
    base_price = round(response['basePrice'], 2)
    last_price = round(response['lastPrice'], 2)
    previous_close = round(response['previousClose'], 2)
    price_change = round(response['change'], 2)
    percent_price_change = round(response['pChange'], 2)
    if (last_price > base_price):
        price_change = f"+{price_change}"
        percent_price_change = f"+{percent_price_change}"
    elif (last_price < base_price):
        price_change = f"-{price_change}"
        percent_price_change = f"-{percent_price_change}"
    else:
        pass

    data = f"{base_price},{price_change},{percent_price_change},{last_price},{previous_close}"
    return data


if __name__ == "__main__":
    price = get_stock_data(arg1)
    print(price)
