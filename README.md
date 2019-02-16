# Synology CLI

Synology CLI provides easy access to your Synology applications using the official Synology API.

## What is the purpose of Synology CLI?

Since Synology itself does not provide a cli tool, it is very hard to integrate Synology applications into scripts. 

As an example: You want to download a lot of files using your Synology DownloadStation.
You have to manage the Synology authentication and the correct calls to DownloadStation API on your own. The simple task downloading a lot of files becomes very complicated. 

Using Synology CLI, all this can be done with a single call:
```
$> cat download-links.txt | synology-cli ds task add -
``` 

## How to use Synology CLI?

The Synology CLI provides on each level the `help` command with examples. 

Just try:
```
$> synology-cli
# or
$> synology-cli help
```

## Issues

If you find any bug or lack in the documentation please submit an issue or even a pull request.


## Contributing

You are very welcome to contribute to this project.
