import g4f
import json
import flask


app = flask.Flask(__name__)
app.config['SECRET_KEY'] = 'Han6some_pe1verts'

@app.route('/message', methods=['POST'])
def get_message():
    data = flask.request.json
    # return data
    print(data)
    response = g4f.ChatCompletion.create(
        model=data["info"][0]["model"],
        provider= g4f.Provider.Hashnode,
        messages=data["messages"],
    )
    print(response)
    json_resp = {'role' : 'assistant', 'content' : response}
    return json.dumps(json_resp)


if __name__ == '__main__':
    app.run(port=8080, host='127.0.0.1')