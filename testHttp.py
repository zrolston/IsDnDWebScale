import requests
import json

#login = requests.get('https://idnddb-195923.appspot.com/api/signup', auth=('davin', 'p@ssword'))

login = requests.get('https://idnddb-195923.appspot.com/api/login', auth=('davin', 'p@ssword'))
print(login.text)

authToken = login.json()['str']

gerald = requests.get('https://idnddb-195923.appspot.com/api/getGerald', auth=(authToken,''))

print(gerald.text)

blah = {"name":"Gerald"}
put = requests.post('https://idnddb-195923.appspot.com/api/putChar', auth=(authToken,''), json=blah)
print(put.text)
"""
gGet = requests.get('https://idnddb-195923.appspot.com/api/getChars', auth=(authToken,''))
print(gGet.text)

remove = requests.post('https://idnddb-195923.appspot.com/api/deleteChars', auth=(authToken,''))
print(remove.text)

gGet = requests.get('https://idnddb-195923.appspot.com/api/getChars', auth=(authToken,''))
print(gGet.text)
"""
