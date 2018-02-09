installed *dep* with `brew install dep` 

After the first lines of *main.go*

`$ dep init`

the dependency tree got updated.

and modified the gopkg.toml to select the desired verson

then ran `$ dep ensure` to re-install the package with the desired version

Added a *new* package with `$ dep ensure --add <sourceWeb/username/package>

Ex: `dep ensure -add github.com/leekchan/accounting`

A warning message appear to alert that the package should be imported before the next `dep ensure`

updated the main.go with new lines and ran it

`$ go run main.go`

