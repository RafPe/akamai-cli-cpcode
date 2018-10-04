
# Akamai CLI - CPCode manager
The AkamaiCPCode manager wraps Akamai's {OPEN} APIs to let you manage CP codes.

Should you miss something we *gladly accept patches* :)

CLI uses custom [Akamai API client](https://github.com/apiheat/go-edgegrid)

# Configuration & Installation

## Credentials
Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Tools expect proper format of sections in edgerc file which example is shown below

>*NOTE:* Default file location is *~/.edgerc*

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

In order to change section which is being actively used you can
* change it via `--config parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_CONFIG=/Users/jsmitsh/.edgerc`

In order to change section which is being actively used you can
* change it via `--section parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_SECTION=mycustomsection`

>*NOTE:* Make sure your API client do have appropriate scopes enabled

## Installation
The tool can be used as completly standalone binary or in conjuction with akamai-cli 

### Akamai-cli ( recommended )

1.  Execute the following from console
    `akamai install https://github.com/apiheat/akamai-cli-cpcodes`

### Standalone
As part of automated releases/builds you can download latest version from the project release page

# Actions

```shell

NAME:
   akamai-akamai-cpcodes - A CLI to manage Akamai CPcodes

USAGE:
   akamai-akamai-cpcodes [global options] command [command options] [arguments...]

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     create   Creates new cpcode
     list     List cpcodes and account info
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/rpieniazek/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --debug value            Debug Level [$AKAMAI_EDGERC_DEBUGLEVEL]
   --help, -h               show help
   --version, -v            print the version


```

# Development
In order to develop the tool with us do the following:
1. Fork repository
1. Clone it to your folder ( within *GO* path )
1. Ensure you can restore dependencies by running 
   ```shell
   dep ensure
   ```
1. Make necessary changes
1. Make sure solution builds properly ( feel free to add tests )
   ```shell
   go build -ldflags="-s -w -X main.appVer=v1.2.3 -X main.appName=akamai-cpcodes" -o akamai-cpcodes
   ```