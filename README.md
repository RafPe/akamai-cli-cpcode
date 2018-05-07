
# Akamai CLI - CPCode manager
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

The AkamaiCPCode manager wraps Akamai's {OPEN} APIs to let you manage CP codes. You can create/list.

This tool have been created with idea of using *akamai cli*. It supports most of the methods provided by API of Akamai. Should you miss something we gladly *accept patches* :)

It uses custom [Akamai API client](https://github.com/RafPe/go-edgegrid)

<!--ts-->
   * [Akamai CLI - CPCode manager](#akamai-cli---cpcode-manager)
   * [Configuration &amp; Installation](#configuration--installation)
      * [API Credentials and sections](#api-credentials-and-sections)
      * [Installation](#installation)
         * [Via akamai-cli ( recommended )](#via-akamai-cli--recommended-)
         * [Standalone](#standalone)
      * [App overview](#app-overview)
         * [General](#general)
         * [List](#list)
         * [Create](#create)
      * [Credits](#credits)

<!--te-->
# Configuration & Installation

## API Credentials and sections
Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Tools expect proper format of sections in .edgerc which looks as follow

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

In order to change section which is being actively used you can
* change via `--section parameter` of the tool itself
* change via env variable `export AKAMAI_EDGERC_SECTION=mycustomsection`

Make sure your API client do have approiate scopes enabled to manage network lists

## Installation
Available in two different ways.With akamai-cli toolkit or as a standalone version

### Via akamai-cli ( recommended )

1.  Execute the following from console
    `akamai install https://github.com/RafPe/akamai-cli-cpcodes`

### Standalone
To compile it from source, you will need Go 1.9 or later, and the [Glide](https://glide.sh) package manager installed:
1. Fetch the package:
   `go get https://github.com/RafPe/akamai-cli-cpcodes`
1. Change to the package directory:
   `cd $GOPATH/src/github.com/RafPe/akamai-cli-cpcodes`
1. Install dependencies using Glide:
   `glide install`
1. Compile the binary:
   `go build -ldflags="-s -w -X main.version=X.X.X" -o akamai-cpcodes`



## App overview

### General

```shell
NAME:
   akamai-cpcodes - A CLI to interact with Akamai CP codes

USAGE:
   akamai-cpcodes [global options] command [command options] [arguments...]

VERSION:
   1.0.0

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
   --output value           Defines output type ( json | table )  (default: "table")
   --help, -h               show help
   --version, -v            print the version
```


### List
This main command allows you to exexute `get` actions on Akamai.

* List account available contracts

    ```
    > [SHELL]  RafPe $ akamai cpcodes list contracts
    # ID           Name
    ctr_G-123POEPE   INDIRECT_CUSTOMER
    ctr_3-45678PEP   INDIRECT_CUSTOMER
    ```

* List products under account

    ```
    > [SHELL]  RafPe $  akamai cpcodes list products -contractID ctr_3-45678PEP
    # ID               Name
    prd_Site_Defender  Site_Defender
    ```

* List groups

    ```
    > [SHELL]  RafPe $ akamai cpcodes list groups
    # ID        ContractID       Name
    grp_123   [ctr_3-45678PEP]   Some Bank N.V.-3-45678PEP
    grp_111   [ctr_3-45678PEP]   some_prop
    grp_716   [ctr_3-45678PEP]   Some Group
    grp_715   [ctr_3-45678PEP]   Some Bank
    grp_117   [ctr_G-123POEPE]   Some Bank N.V.-G-29QEPME
    ```

* List CPcodes

    ```
    > [INSERT] RafPe $ akamai cpcodes --contractID ctr_3-ZYI41V --groupID grp_123

    # ID        ProductIDs           Created                        Name
    cpc_123  [prd_Site_Defender]  2013-04-15 07:05:28 +0000 UTC  acceptance.a.com
    cpc_456  [prd_Site_Defender]  2018-03-28 11:22:57 +0000 UTC  acceptance.b.com

    ```

### Create
Used for creating new network CP codes

* Create new CP code
    ```shell
    > $ akamai cpcodes create --contractID ctr_3-ZYI41V --groupID grp_123  --CPcodeName demo
    ok
    ```

## Credits
* [Petr](https://github.com/partamonov) - for being the mentor on Golang :)




