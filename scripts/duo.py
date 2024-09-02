import sys
import json

# add package to system path
sys.path.insert(0,'/home/erob/Documents/projects/anki_duo/duolingo/')
from duolingo import duolingo

def vocab_dump(JWT: str):
    print("Logging into Duolingo...")
    lingo = duolingo.Duolingo(jwt=JWT)
    print("Getting vocab...")
    vocab = lingo.get_vocabulary('es')
    with open('vocabulary.json', 'w') as f:
        print("Dumping json...")
        json.dump(vocab, f)

