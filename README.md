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

If you're not my type, you can instead [download a binary](https://github.com/ashwanthkumar/wasp-cli/releases) directly from releases. 

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

### Updating config keys
You can use the `wasp put` command to update things, but you can do more than just update a value. You can use unix pipes `(|)` to update certain values. 

For example if you want to pass a JSON string as a config value, 
```
$ cat foo.json | wasp put foo.bar.baz --stdin
```

If you want to set the JSON as a nested configuration instead of a JSON value,
```
$ cat foo.json | wasp put foo.bar.baz --stdin --raw
```
