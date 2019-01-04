import requests
import json
import random
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
        if (initRequest['total_count'] - (requestsToDo * 100)) > 0:
            remainderItems = requestItems(appid, initRequest['total_count'] % 100, requestsToDo*100)
            items.append(remainderItems['results'])
    return items

def getItemNameId(appid, hash_name):
    targetString = "Market_LoadOrderSpread("
    r = requests.get(f'https://steamcommunity.com/market/listings/{appid}/{hash_name}')
    if r.status_code == 429:
        sleepTime = 60 * 3
        print(f"Too many requests detected, sleeping for {sleepTime/60} minutes...")
        sleep(sleepTime)
        r = requests.get(f'https://steamcommunity.com/market/listings/{appid}/{hash_name}')
    target = r.text.find(targetString)
    return r.text[target+24:target+33]

itemRequests = getAllItems("583950")
total = 0
allItems = []

reset = 0
for request in itemRequests:
    for item in request:
        storeItem = {
            "name": item['name'],
            "hash_name": item['hash_name'],
            "item_name_id": getItemNameId(item['asset_description']['appid'], item['hash_name']),
            "app_name": item['app_name'],
            "appid": item['asset_description']['appid']
        } 
        allItems.append(storeItem)
        print(f"{storeItem['name']} - {storeItem['item_name_id']}: Completed")
        #sleep(random.randint(2, 6))
        total += 1

itemJSON = json.dumps(allItems)
f = open("item_data.json", "w")
f.write(itemJSON)
f.close()
print("Process complete")