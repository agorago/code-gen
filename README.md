# code-gen

## Getting started
* git clone [git@github.com:agorago/wego-gen.git](git@github.com:agorago/wego-gen.git)
* git clone [git@github.com:agorago/wego.git](git@github.com:agorago/wego.git)
* cd wego-gen
* make all
* cd ..
* use `wego-gen/bin/gen.sh` to generate the code (input required values)

## Example

``` sh
PWD: $HOME/sample
RUN: ls
OUPUT: bplus		code-gen	sample.go

CHECK: code-gen/sample.go for the reference
```
- cd sample 
- make generate-error-codes
- make copy-bundles
- make build
- make run
