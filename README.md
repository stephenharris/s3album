# S3 Album Viewer

A very simple VueJS + Golang application to view and delete photos & videos from an S3 bucket


## Building UI

The UI (html, JavaScript, stylesheet) is converted into Go code using [go-bindata](https://github.com/go-bindata/go-bindata).

```
go-bindata public/
```

## Installation

    wget https://raw.githubusercontent.com/stephenharris/s3album/master/install.sh -O install.sh \
    && chmod +x install.sh
    && sudo install.sh <bucket-name> <aws-region>