# File: backend/backend.py

from flask import Flask, request, jsonify
import json

app = Flask(__name__)

# In-memory storage (not suitable for production; use a database like SQLite or PostgreSQL)
users = []  # List to store users: e.g., [{"username": "user1", "password": "pass"}]
articles = []  # List to store articles: e.g., [{"title": "Article 1", "content": "Details", "author": "user1"}]

@app.route('/register', methods=['POST'])
def register():
    data = request.json  # Expect JSON like {"username": "user1", "password": "pass"}
    if any(user['username'] == data.get('username') for user in users):
        return jsonify({"error": "User already exists"}), 400
    users.append(data)
    return jsonify({"message": "User registered successfully"}), 201

@app.route('/login', methods=['POST'])
def login():
    data = request.json  # Expect JSON like {"username": "user1", "password": "pass"}
    for user in users:
        if user['username'] == data.get('username') and user['password'] == data.get('password'):
            return jsonify({"message": "Login successful", "user": user['username']}), 200
    return jsonify({"error": "Invalid credentials"}), 401

@app.route('/submit-article', methods=['POST'])
def submit_article():
    data = request.json  # Expect JSON like {"title": "Article Title", "content": "Article content", "author": "user1"}
    if not data.get('author') or not any(user['username'] == data.get('author') for user in users):
        return jsonify({"error": "Author not registered"}), 403
    articles.append(data)
    return jsonify({"message": "Article submitted successfully", "article": data}), 201

@app.route('/articles', methods=['GET'])
def get_articles():
    return jsonify(articles)  # Returns a list of all articles

if __name__ == '__main__':
    app.run(debug=True, port=5000)  # Run on port 5000
