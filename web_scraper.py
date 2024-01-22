import sys
from bs4 import BeautifulSoup
import requests

arg1 = sys.argv[1]

def get_stock_data(stock_sym):

    url = f'https://www.google.com/finance/quote/{stock_sym}:NSE?hl=en'
    response =requests.get(url)
    soup =BeautifulSoup(response.text,"html.parser")
    class1="YMlKec fxKbKc"
    price =float(soup.find(class_=class1).text.strip()[1:].replace(",",""))
    return price

if __name__ == "__main__":
    price =get_stock_data(arg1)
    print(price)
