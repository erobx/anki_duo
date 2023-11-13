import duolingo
import inspect
import json
from dotenv import load_dotenv
import os

load_dotenv()

source = inspect.getsource(duolingo)
new_source = source.replace('jwt=None', 'jwt')
new_source = source.replace('self.jwt = None', ' ')
exec(new_source, duolingo.__dict__)

username = os.environ['LOGIN']
jwt = os.environ['JWT']

lingo = duolingo.Duolingo(username, jwt=jwt)
print(lingo.get_friends())

# vocabulary = lingo.get_vocabulary('es')

# with open('vocabulary.json', 'w') as f:
#     json.dump(vocabulary, f)