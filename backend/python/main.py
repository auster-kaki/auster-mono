from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/diary', methods=['POST'])
def diary():
    data = request.get_json()
    # Process the data as needed
    return jsonify({"message": "Diary entry received", "data": data}), 200

if __name__ == '__main__':
    app.run(debug=True)