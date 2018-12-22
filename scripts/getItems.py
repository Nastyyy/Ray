import requests
import json
from math import floor
from time import sleep

# This script deals with requesting item information from Steam's Market API.
# It uses Steam's search endpoint. If you set a norender paramter to 1,
# the response will give you the raw json data.
# The response's data is related to all of the items in 

def requestItems(appid, count, start):
    url = f"https://steamcommunity.com/market/search/render/?norender=1&appid={appid}&count={count}&start={start}"

    resp = requests.get(url)
    data = resp.json()

    return data

def getAllItems(appid):
    items = []
    initRequest = requestItems(appid, 100, 0)
    items.append(initRequest['results'])
    requestsToDo = floor(initRequest['total_count'] / 100)
    if requestsToDo >= 2:
        i = 1
        while i < requestsToDo:
            reqItems = requestItems(appid, 100, i * 100)
            items.append(reqItems['results'])
            i+=1
            sleep(3)
    return items


def printItems(items):
    for item in items['results']:
        if item['asset_description']['marketable'] != 1:
                print(f"{item['app_name']} not marketable")

        print(item['app_name'], item['asset_description']['appid']," - ", item['name'], item['hash_name'])

print(getAllItems("583950"))