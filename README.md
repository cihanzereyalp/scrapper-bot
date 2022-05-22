# scrapper-bot

Simple scrapper sending extracted data to the discord server/channel using discord's webhook function.

Create a config.json file and place it at the root of your project
```
|-- main.go
|-- main_test.go
|-- config.json    <---
|-- ...
```
Example config.json:
```
{
    "urlList": ["yourFirtURL", ...],
    "webHookURL": "yourDiscordWebHookURL",
    "pattern": "<your> regular expression group here (\\d+)</your>",
}
```

<sub>Make sure your URLs are allowed before using it since the scrapper bot cannot check disallowed files, directories, or web pages.</sub>