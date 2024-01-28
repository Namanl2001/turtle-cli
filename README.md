# Turtle

The `turtle` CLI enables you to securely store your API keys locally in the directory `${HOME}/turtle-secrets`. This is achieved by encrypting the keys using a password provided by you.

# Installation

```
go build main.go
go install
```

# Usage

```
$turtle encrypt --name=github --key=qawsgty12w --password=Namanl2001
api-key named github saved successfully 

$turtle decrypt --name=github --password=Namanl2001
qawsgty12w

$turtle delete --name=github
successfully deleted api-key github 
```