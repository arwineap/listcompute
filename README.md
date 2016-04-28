## Getting started
1. install golang build tools
    * https://golang.org/doc/install
    * sudo pacman -S go
    * sudo apt-get install golang
2. go build listcompute.go
3. Setup credential files
    * http://docs.aws.amazon.com/aws-sdk-php/v2/guide/credentials.html#credential-profiles
    * I like to setup a default profile, and a dev profile


## using
```
$ listcompute ops prd consul
ops-prd-consul-uswest1-3 46.35.123.200 10.1.14.49
ops-prd-consul-uswest1-2 46.35.123.201 10.1.13.232
ops-prd-consul-uswest1-1 46.35.123.202 10.1.14.215

$ listcompute -e ops prd consul
46.35.123.200
46.35.123.201
46.35.123.202

$ listcompute -i ops prd consul
10.1.14.49
10.1.13.232
10.1.14.215

$ AWS_PROFILE=default listcompute -n ops prd consul
ops-prd-consul-uswest1-3
ops-prd-consul-uswest1-2
ops-prd-consul-uswest1-1
```
