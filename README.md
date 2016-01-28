# WASP CLI
wasp-go is a command line utility for WASP. You can get / browse WASP configurations directly from your terminal without any complicated `cURL` commands. 

## Build
You need to have golang setup in your system to build it locally. 

```bash
$ git clone https://github.com/ashwanthkumar/wasp-go
$ cd wasp-go
$ make setup
$ make build
```

## Configuration file
You need to create a file under `~/.indix/wasp.json` which has the following contents
```json
{
  "wasp": {
    "host": "http://localhost:9000",
    "token": "auth-token-to-make-puts-and-deletes"
  }
}
```

You can still make `ls` and `get` calls without the `token` config but the `put` and `delete` might fail. 

## Usage

```
$ wasp -h
Command line client to WASP

Usage:
  wasp [command]

Available Commands:
  get         Get the configuration in current path
  ls          List keys in the current path
  put         Put a value against a path

Use "wasp [command] --help" for more information about a command.
```
