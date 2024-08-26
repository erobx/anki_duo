import duo
import words

def main():
    duo.vocab_dump()
    print("Writing words to Anki plaintext file...")
    w = words.extract_words()
    words.write_to_file(w)
    print("Done!")


if __name__ == '__main__':
    main()
