# sls-go-template

## Delete this section

  * Edit these files to your project:
    * `go.mod`
    * `serverless.yml`
      * modify `org` at the very beginning 
    * `package.json`
    * this `README.md`

## Deploying

(requires make, podman-docker and yarn):

```
# aws ecr get-login-password --region us-west-2 | podman login -u AWS --password-stdin "235368163414.dkr.ecr.us-west-2.amazonaws.com"
$ go get -v ./...
$ yarn install --frozen-lockfile
$ make && yarn sls deploy
```

## Details

