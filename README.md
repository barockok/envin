envin ![License MIT](https://img.shields.io/badge/license-MIT-blue.svg) [![Build Statue](https://travis-ci.org/barockok/envin.svg?branch=master)](https://travis-ci.org/barockok/envin)
=============

this is an utility to generate or overwrite JSON and YAML file used as your application configuration through environment variables. the initial idea is come from how manage the configuration file in docker container, most of us will create a template file then we can customize it through some environment variables. by use envin we no longer need to write a template file to have a customize configuration, we just need the original configuration and envin will do the job to customize it throught environment variables.

## Installation

todo

## Usage

envin can be used inside your script where your docker image will run it through ENTRYPOINT or CMD directive, on top of that it also can wrap your command  

inside a bash script

```bash
#!/bin/env bash
envin -prefix APP -src /etc/your_app_name/app.json -filetype JSON
```
this command will try to collect all environment vars started with prefix `APP_`. for example you have variables shown below.

```bash
APP_http_proxy=http://127.0.0.2:8181
APP_https_proxy=https:///127.0.0.2:8282
APP_mysql___host=127.0.0.3
APP_mysql___port=3308
APP_ip_whitelists___0=127.0.0.4
APP_ip_whitelists___9999=127.0.0.19
```

and your `app.json` content is like so 

```json
{
  "secret_key" : "abcdeftgh1gasdasgdu176231723",
  "host" : "http://example.com",
  "http_proxy" : "http://127.0.0.2:8080",
  "mysql" : {
    "username" : "root",
    "password" : "secret",
    "host" : "127.0.0.1",
    "port" : 3306
  },
  "ip_whitelists" : [
    "127.0.0.5",
    "127.0.0.6",
  ],
  "report_exception" : true
}
```
envin will overwrite it and it'll looks like this

```json
{
  "secret_key" : "abcdeftgh1gasdasgdu176231723",
  "host" : "http://example.com",
  "http_proxy" : "http://127.0.0.2:8181",
  "https_proxy" : "https://127.0.0.2:8282",
  "mysql" : {
    "username" : "root",
    "password" : "secret",
    "host" : "127.0.0.3",
    "port" : 3308
  },
  "ip_whitelists" : [
    "127.0.0.4",
    "127.0.0.6",
    "127.0.0.19"
  ],
  "report_exception" : true
}
```


## License

MIT

