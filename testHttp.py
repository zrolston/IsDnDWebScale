import requests
import json

#login = requests.get('https://idnddb-195923.appspot.com/api/signup', auth=('trev', 'vroom'))

login = requests.get('https://idnddb-195923.appspot.com/api/login', auth=('trev', 'vroom'))
print(login.text)

authToken = login.json()['str']

gerald = requests.get('https://idnddb-195923.appspot.com/api/emptyChar', auth=(authToken,''))

print(gerald.text)


put = requests.post('https://idnddb-195923.appspot.com/api/putChar', auth=(authToken,''), json=gerald.json())
print(put.text)

gGet = requests.get('https://idnddb-195923.appspot.com/api/getChars', auth=(authToken,''))
print(gGet.text)
"""
remove = requests.post('https://idnddb-195923.appspot.com/api/deleteChars', auth=(authToken,''))
print(remove.text)

gGet = requests.get('https://idnddb-195923.appspot.com/api/getChars', auth=(authToken,''))
print(gGet.text)
"""
