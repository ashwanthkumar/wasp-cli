# WASP CLI
wasp-cli is a command line utility for WASP. You can get / browse WASP configurations directly from your terminal without any complicated `cURL` commands. 

## Build
You need to have golang setup in your system to build it locally. 

```bash
$ git clone https://github.com/ashwanthkumar/wasp-cli
$ cd wasp-cli
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
$ wasp
Command line client to WASP

Usage:
  wasp [command]

Available Commands:
  gen         Generate files using Go templates from WASP configs
  get         Get the configuration in current path
  ls          List keys in the current path
  put         Put a value against a path
  rmr         Delete a config key (recursively)

Flags:
  -h, --help   help for wasp

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

### Deleting config keys
You can delete a configuration key using
```
$ wasp rmr foo.bar.baz
Do you want to delete everything under foo.bar.baz? [yN]: y
```

If you're running deletes in a batch mode (strictly not advisable),
```
$ wasp rmr foo.bar.baz --yes
```

### Generate Configuration files for your application from WASP Settings
You can generate configuration files for your application based on the settings in WASP
```
// configuration.tmpl
{
 "wasp": {
  "host": "{{ .wasp.host }}"
 }
}

$ wasp gen dev.golang /path/to/configuration.tmpl > /path/to/config.json
```

We use the excellent [Go Templates](https://golang.org/pkg/text/template/) for rendering the files.

### Tips
If you've `jq` installed you can prettify the Configuration outputs much more by just piping the output to jq. Don't use this if you're writing it to a file.
```
$ wasp get production.foo.bar.baz | jq --color-output
```
