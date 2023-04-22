import os
import string
import secrets
from flask import Flask, jsonify, request

app = Flask(__name__)


@app.route('/passwords/<int:length>', methods=['GET'])
def generate_password(length):
    if length < 8:
        return jsonify({'error': 'Password length must be at least 8'}), 400

    alphabet = string.ascii_letters + string.digits + string.punctuation
    password = ''.join(secrets.choice(alphabet) for _ in range(length))

    return jsonify({'password': password})


if __name__ == '__main__':
    port = int(os.environ.get('PORT', 5000))
    app.run(host='0.0.0.0', port=port)
