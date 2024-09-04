#!/usr/local/bin/python3
import sys
import duo
import json

def extract_words() -> dict:
    words = {}
    with open('vocabulary.json', 'r') as f:
        data = json.load(f)
        for i in data:
            words[i['text']] = i['translations']

    return words


def write_to_file(file_name: str, words: dict):
    print("Writing to Anki import file...")
    with open(file_name, 'w') as f:
        for k,v in words.items():
            f.write(k)
            f.write(";")
            for i in range(len(v)):
                f.write(v[i])
                if i != len(v)-1:
                    f.write(", ")
            f.write("\n")


def main():
    JWT = sys.argv[1]
    duo.vocab_dump(JWT)
    w = extract_words()
    write_to_file("vocab.txt", w)
    print("Done!")


if __name__ == '__main__':
    main()
