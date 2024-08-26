import sys
import os
import json
from dotenv import load_dotenv

# add package to system path
sys.path.insert(0,'../duolingo/')
from duolingo import duolingo

def vocab_dump():
    load_dotenv()
    JWT = os.getenv('JWT')
    print("Logging into Duolingo...")
    lingo = duolingo.Duolingo(jwt=JWT)
    print("Getting vocab from Duolingo...")
    vocab = lingo.get_vocabulary('es')
    with open('vocabulary.json', 'w') as f:
        print("Dumping json...")
        json.dump(vocab, f)

