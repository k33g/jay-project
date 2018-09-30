### Documentation here

- https://gist.github.com/alexellis/d648d927c34c082bb5d965f06b818026

> ⚠️ you need to add the OpenFaaS integration

### Create a function

```shell
faas-cli new welcome --lang node --prefix k33g
```

Rename the `.yml` file to `stack.yml` so that OpenFaaS Cloud knows where to find your function(s)


### Deploy

```shell
git add .
git commit -s -m ":tada: Initial function"
git push origin master
```

### Add function

You can have one or many functions in your stack.yml file.

To append new functions type in:

```shell
faas new <name-of-new-function> --lang go --prefix k33g --append=./stack.yml 

faas new hello --lang go --prefix k33g --append=./stack.yml 

```


### DashBoard

- https://system.o6s.io/dashboard/k33g

### Function

- https://k33g.o6s.io/welcome

