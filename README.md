# code-gen

## HOW TO GET STARTED
- git clone [git@ssh.intelligentb.com:devops/code-gen.git](https://gitlab.intelligentb.com/devops/code-gen.git)
- git clone [git@ssh.intelligentb.com:devops/bplus.git](https://gitlab.intelligentb.com/devops/bplus.git)
- cd code-gen
- make all
- cd ..
- use `code-gen/bin/gen.sh` to generate the code (input required values)

##Example:

```
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
