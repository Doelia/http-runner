

HTTP-RUNNER permet de lancer des commandes unix depuis une API http.

## Installation
```
apt install http-runner
service http-runner start
```

## Configuration
Edit the file `~/.http-runner/config.yaml` :
```
port: 80
host: 0.0.0.0
security:
    auth:
        type: BASIC_AUTH
        basic_auth:
            login: admin
            password: ''
    # Array of IP authorised. Set as "*" for wildcard
    ip_authorised: []
```

Restart the service after editing :
```
service http-runner restart
```

## Add script

Add your own scripts in the `~/.http-runner/scripts` folder.

Call your script from HTTP with her file name. Example : `http://localhost/api/call/go.sh` will call `~/.http-server/scripts/go.sh`

## Logs

Logs are generated into the `~/.http-runner/logs` folder.

There are organised like so :
```./logs/<script_name>/YYYY-MM-DD-HH-MM-SS-%d.log```

You can access from the HTTP api on /api/logs/<script_name>/<file.log>

## Web interface

You can access to the web interface on http://localhost/admin