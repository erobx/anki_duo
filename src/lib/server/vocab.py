import json

file_path = 'vocabulary.json'

def main():
    vocab = read_from_file(file_path)
    print(vocab)


def read_from_file(file_path):
    with open(file_path, 'r') as f:
        vocab = json.load(f)

    return vocab


if __name__ == main():
    main()