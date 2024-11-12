# Duolingo -> Anki

Retrieve your studied vocab words from Duolingo and import them into Anki.

## Requirements

This project requires python to be installed. You also must provide your JWT from Duolingo to login.

The unofficial [Duolingo api](https://github.com/erobx/golingo) has a guide on how to retrieve your JWT.

You also need the [Anki client](https://apps.ankiweb.net/) installed on your system with the [AnkiConnect](https://ankiweb.net/shared/info/2055492159) add-on.

## Running the project

Make sure Anki is running before running the Makefile.

Run the Makefile to bulid the project, then provide your JWT.

```bash
make
```

