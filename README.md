<h1 align="center"><img src="aergie.svg" alt="Aergie: An easy alternative to makefile" title="An easy alternative to makefile"></h1>

[![Continuous Integration](https://github.com/IQ2i/aergie/workflows/Continuous%20Integration/badge.svg?branch=master)](https://github.com/IQ2i/aergie/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/IQ2i/aergie)](https://goreportcard.com/report/github.com/IQ2i/aergie)
[![Coverage Status](https://coveralls.io/repos/github/IQ2i/aergie/badge.svg?branch=master)](https://coveralls.io/github/IQ2i/aergie?branch=master)

# Installation

Run the following installer to download the Aergie binary:

```bash
curl -sS https://raw.githubusercontent.com/IQ2i/aergie/master/script/install | sudo bash
```

If you can, run the above command in sudo. This will allow the installer to put the script in a better directory and to install the autocompletion.

# How to use Aergie

The first thing to do is to create a file at the root of your project. The file can have several possible names:

* `.aergie.yml`
* `.aergie.yaml`

Then you just need to declare two blocks: 

### `commands` (require)

This is where you can declare the different commands that allow you to work on your project.  
A command is form with:

|       | Description                                 | Example (see example below)         |
|-------|---------------------------------------------|-------------------------------------|
| name  | This is the key of the command element      | `start`                             |
| help  | Sentence to describe what the command does  | `help: Start docker compose`        |
| steps | An array with each action to be carried out | `steps: [docker-compose up -d]`     |

Example:

```yaml
commands:
    start:
        help: Start docker compose
        steps:
            - docker-compose up -d

    stop:
        help: Stop docker compose
        steps:
            - docker-compose stop

    install:
        help: Install my application
        steps:
            - ${php} composer install
            - ...
```

### `variables` (optional)

You can declare variables to use in your steps:

```yaml
variables:
    php: docker exec -it php_container_name

commands:
    install:
        help: Install my application
        steps:
            - ${php} composer install
```
