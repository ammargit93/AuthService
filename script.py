import requests
import sqlite3


data = {
    'Username':'Affan',
    'Password':'09876',
    'DBPassword': 'yourord',
    'DBTag': 'firstdatabae'
}

# response = requests.post('http://localhost:8080/register', json=data)

# print(response.json())

conn = sqlite3.connect("models/UserDB.db")
cur = conn.cursor()
cur.execute("select * from Users")
conn.commit()
row = cur.fetchall()
for r in row:
    print(r)