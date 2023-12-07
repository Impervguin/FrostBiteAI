import g4f
import json
import flask


app = flask.Flask(__name__)
app.config['SECRET_KEY'] = 'Han6some_pe1verts'

I = 0

MAX_TRIES_COUNT = 10

PROVIDERS = [
    g4f.Provider.FakeGpt,
    g4f.Provider.GptGo,
    # g4f.Provider.ChatBase,
    g4f.Provider.ChatgptX,
    g4f.Provider.FreeGpt,
    g4f.Provider.GPTalk,
    g4f.Provider.GptForLove,
    g4f.Provider.GptGo,
    g4f.Provider.You,
]


def create_error(errstr : str):
    return json.dumps({"error" : errstr})

def check_message_data(data : dict) -> tuple:
    if "messages" not in data.keys():
        return 1, "Error: No messages in request"

    if not isinstance(data["messages"], list) or len(data["messages"]) == 0:
        return 2, "Error: Not correct format of messages array"

    for mess in data["messages"]:
        if "role" not in mess.keys() or "content" not in mess.keys():
            return 3, f"Error: Not correct format in message: {mess}"
    
    return 0, "ok"


@app.route('/message', methods=['POST'])
def get_message():
    global I
    
    tries = 0

    if not flask.request.is_json:
        return create_error("Error: request doesn't have json form.")
    
    data = flask.request.json  
    err = check_message_data(data)
    if err[0]:
        return create_error(err[1])

    
    while tries < MAX_TRIES_COUNT:
        try:
            response = g4f.ChatCompletion.create(
                model="gpt-3.5-turbo",
                provider= PROVIDERS[I],
                messages=data["messages"],
            )
            break
        except BaseException:
            print(f"Provider failed: {PROVIDERS[I]}")
            I = (I + 1) % len(PROVIDERS)
            tries += 1
    if tries == MAX_TRIES_COUNT:
        return create_error("Error: couldn't find any server")
    json_resp = {'role' : 'assistant', 'content' : response}
    return json.dumps(json_resp)


if __name__ == '__main__':
    app.run(port=8080, host='127.0.0.1')
