import json

with open('known_words.json', 'r') as f:
    known_words = json.load(f)

print(len(known_words))