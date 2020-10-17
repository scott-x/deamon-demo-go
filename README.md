# deamon-demo-go

This demo is base on [cobra](https://github.com/spf13/cobra)


### errors 

-  `start error: exec: "demo": executable file not found in $PATH `: to enable the code running in deamon, you should mv the excuable binary to `$GOPATH`


### notes

- You should ovoid using existing command, such as `test`(a python build-in command)