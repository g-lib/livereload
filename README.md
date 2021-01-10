# livereload

Reload webpages on changes, without hitting refresh in your browser.

## Install

```shell
go get -u github.com/g-lib/livereload
```

## Usage of Command Line Interface

LiveReload provides a command line utility, `livereload`, for starting a server in a directory.

By default, it will listen to port 35729, the common port for [LiveReload browser extensions](http://feedback.livereload.com/knowledgebase/articles/86242-how-do-i-install-and-use-the-browser-extensions-).

```shell
$ livereload --help
NAME:
   livereload - Start a `livereload` server

USAGE:
   livereload [global options] command [command options] [arguments...]

DESCRIPTION:
   Reload webpages on changes, without hitting refresh in your browser

AUTHOR:
   Tacey Wong <xinyong.wang@qq.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host Hostname                   Hostname to run livereload server on (default: "127.0.0.1")
   --port Port, -p Port              Port to run livereload server on (default: 35729)
   --directory Directory             Directory to serve files from (default: ".")
   --target File, -t File            File or `directory` to watch for changes
   --wait seconds, -w seconds        Time delay in seconds before reloading (default: 0)
   --open-url-delay value, -o value  If set, triggers browser opening <D> seconds after starting (default: 0)
   --debug, -d                       Enable  pretty logging (default: false)
   --help, -h                        show help (default: false)

COPYRIGHT:
   2020 - Now Tacey Wong , BSD-3-Clause License
```

## Usage of Package