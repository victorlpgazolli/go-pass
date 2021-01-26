# go-pass

Simple CLI Password Manager

🚧 **go-pass is under development** 🚧
## Examples 

```shell
 $ # Set password:
 $ pass -n facebook -p 123456
 $ # Get password:
 $ pass -n facebook
 $ # List passwords:
 $ pass -l
```

## Contents

* [Features](#features)
* [Installation](#installation)

## Features

* Save password
* Read password
* List passwords


### Installation

You can install this via the command-line with either `curl`, `wget` or another similar tool.

| Method    | Command                                                                                           |
|:----------|:--------------------------------------------------------------------------------------------------|
| **curl**  | `sh -c "$(curl -fsSL https://raw.githubusercontent.com/victorlpgazolli/go-pass/master/install.sh)"` |
| **wget**  | `sh -c "$(wget -O- https://raw.githubusercontent.com/victorlpgazolli/go-pass/master/install.sh)"`   |
| **fetch** | `sh -c "$(fetch -o - https://raw.githubusercontent.com/victorlpgazolli/go-pass/master/install.sh)"` |

#### Manual inspection

It's a good idea to inspect the install script from projects you don't yet know. You can do
that by downloading the install script first, looking through it so everything looks normal,
then running it:

```shell
wget https://raw.githubusercontent.com/victorlpgazolli/go-pass/master/install.sh
sh install.sh
```

