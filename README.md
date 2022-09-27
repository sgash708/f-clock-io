# f-clock-io

automate clock in/out for a certain site....

## TODO

```:bash
make cpcfg
```

## Require
Chromedriver
### linux
```:bash
CHROMEDRIVER_VERSION=`curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE`
curl -sS -o /tmp/chromedriver_linux64.zip http://chromedriver.storage.googleapis.com/$CHROMEDRIVER_VERSION/chromedriver_linux64.zip
unzip /tmp/chromedriver_linux64.zip
mv chromedriver /usr/local/bin/
```

https://chromedriver.chromium.org/downloads

### mac
```:bash
brew install chromedriver
```

## Run

To run `fcio` binaries, you need `config.yml` in the same directory.

```go
make m-build
go install ./bin/fcio
fcio clockin
```
