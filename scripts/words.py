import json

def extract_words():
    words = {}
    with open('vocabulary.json', 'r') as f:
        data = json.load(f)
        for i in data:
            words[i['text']] = i['translations']

    return words


def write_to_file(words):
    with open('import.txt', 'w') as f:
        for k,v in words.items():
            f.write(k)
            f.write(";")
            for i in range(len(v)):
                f.write(v[i])
                if i != len(v)-1:
                    f.write(", ")
            f.write("\n")

