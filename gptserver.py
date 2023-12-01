import g4f
import json
import flask


app = flask.Flask(__name__)
app.config['SECRET_KEY'] = 'Han6some_pe1verts'

I = 0

PROVIDERS = [
    g4f.Provider.GptGo,
    # g4f.Provider.ChatBase,
    g4f.Provider.ChatgptX,
    g4f.Provider.FakeGpt,
    g4f.Provider.FreeGpt,
    g4f.Provider.GPTalk,
    g4f.Provider.GptForLove,
    g4f.Provider.GptGo,
    g4f.Provider.You,
]


@app.route('/message', methods=['POST'])
def get_message():
    global I
    data = flask.request.json
    # return data
    # print(data)
    while True:
        try:
            response = g4f.ChatCompletion.create(
                model=data["info"][0]["model"],
                provider= PROVIDERS[I],
                messages=data["messages"],
            )
            print(PROVIDERS[I])
            break
        except BaseException as e:
            I = (I + 1) % len(PROVIDERS)
    print(response)
    json_resp = {'role' : 'assistant', 'content' : response}
    return json.dumps(json_resp)


if __name__ == '__main__':
    app.run(port=8080, host='127.0.0.1')