# HTTP-RUNNER

Launch custom SH scripts from an HTTP api.

## Installation
```
Soon
```

## Usage

Add your own scripts in the `~/.http-runner/scripts` folder.

Call your script from HTTP with her file name. Example : `http://localhost/api/run/go.sh` will call `~/.http-server/scripts/go.sh`

GET parameters are passed as $1 and BODY as $2

## Configuration
Edit `~/.http-runner/config.yaml` :
```
port: 80
host: 127.0.0.1
security:
  # auth_type can be: NONE, BASIC_BATH
  auth_type: BASIC_AUTH
  basic_auth:
    login: admin
    password: ''
  # Array of IP authorised. Set as "*" for wildcard. You can use mask IP
  ip_authorised: [127.0.0.1]
```

Restart the service after editing :
```
service http-runner restart
```

## Logs

Logs are written into the `~/.http-runner/logs` folder.

There are organised like so :
```./logs/<script_name>/YYYY-MM-DD_HH-MM-SS.log```

You can access from the HTTP api on `/api/logs/<script_name>/<file.log>`

## Web interface

A web interface is available on http://localhost/admin

## Development

This project use [dep](https://github.com/golang/dep) for dependencies

## TODO
- Queue system
- Reluanch scripts
