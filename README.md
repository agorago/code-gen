# code-gen

## HOW TO GET STARTED
- git clone [git@github.com:MenaEnergyVentures/code-gen.git](https://github.com/MenaEnergyVentures/code-gen)
- git clone [git@github.com:MenaEnergyVentures/bplus.git](https://github.com/MenaEnergyVentures/bplus)
- cd code-gen
- make all
- use `code-gen/bin/gen.sh` inside your directory to generate the code
```
Example:
PWD: $HOME/sample
RUN: ls
OUPUT: bplus		code-gen	sample.go
HOW: sample.go looks?
CHECK: code-gen/sample.go
```
- cd `sample` 
- make generate-error-codes
- make build
- RUN bin/main
