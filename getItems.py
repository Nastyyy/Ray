import requests
import json

def getCards(start):
    url = 'https://steamcommunity.com/market/search/render/?norender=1&appid=583950&count=100&start=' + str(start)

    resp = requests.get(url)
    data = resp.json()

    total = 0

    for item in data['results']:
        total += 1
        print(item['app_name'], item['asset_description']['appid']," - ", item['name'], item['hash_name'])
