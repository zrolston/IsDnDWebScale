import requests

r = requests.get('https://idnddb-195923.appspot.com/api/you/should/not/call/this', auth=('benis', 'root'))
#r = requests.get('https://idnddb-195923.appspot.com/api/login', auth=('benis', 'crempo'))

print(r.status_code)

print(r.json())
