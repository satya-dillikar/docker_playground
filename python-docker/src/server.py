from flask import Flask
import os
server = Flask(__name__)

@server.route("/")
def hello():
    return "Hello World!"

@server.route("/hi")
def hi():
    return "Hi!"

if __name__ == "__main__":
    port = int(os.environ.get("PORT", 5000))
    server.run(debug=True,host='0.0.0.0',port=port)    
