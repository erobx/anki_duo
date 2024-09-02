# Duolingo -> Anki

Retrieve your studied vocab words from Duolingo and import them into Anki.

## Requirements

This project requires python to be installed. You also must provide your JWT from Duolingo to login.

The unofficial [Duolingo api](https://gitlab.com/JASchilz/duolingo/-/blob/master/USAGE.md?ref_type=heads) has a guide on how to retrieve your JWT.

You also need the [Anki client](https://apps.ankiweb.net/) installed on your system with the [AnkiConnect](https://ankiweb.net/shared/info/2055492159) add-on.

## Running the project

Run the Makefile to bulid the project, then run the exe with the required arguments.

```bash
make
./anki-duo {JWT} {file_name}
```


