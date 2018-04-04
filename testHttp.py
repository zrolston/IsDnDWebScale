import requests
import json

#r = requests.get('https://idnddb-195923.appspot.com/api/signup', auth=('benis','crempo'))
#print(r.text)
#r = requests.get('https://idnddb-195923.appspot.com/api/you/should/not/call/this', auth=('benis', 'root'))

r0 =  requests.get('https://idnddb-195923.appspot.com/api/login', auth=('benis','crempo'))
authToken = r0.json()['str']
#authToken = "eb3b043f-a654-445c-b3ed-e7874875bcdc"
print(r0.text)

r3 = requests.get('https://idnddb-195923.appspot.com/api/getGerald')

r2 = requests.post('https://idnddb-195923.appspot.com/api/putChar', auth=(authToken,''), json=r3.json())

print(r2.status_code)
print(r2.text)

r1 = requests.get('https://idnddb-195923.appspot.com/api/getChars', auth=(authToken,0))

print(r1.status_code)
print(r1.text)
