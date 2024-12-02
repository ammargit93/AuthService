from flask import Flask, request, jsonify, render_template, redirect, url_for
import requests

app = Flask(__name__)

app.secret_key = 'your_secret_key'

GO_API_URL = 'http://localhost:8080' 


@app.route('/login', methods=['GET','POST'])
def login():
    if request.method == "POST":
        username = request.form['username']
        password = request.form['password']

        data = {
            'Username': username,
            'Password': password,
            'DBTag': 'test_db', 
            'DBPassword': 'securepassword123'
        }
        response = requests.post(f'{GO_API_URL}/login', json=data)
        print(response.json()['token'])
        session['username'] = token
        return redirect('/dashboard')
    return render_template("login.html")

@app.route('/dashboard', methods=['GET','POST'])
def dashboard():
    


app.run(debug=True)