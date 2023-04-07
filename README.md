# Overview
### A web platform developed based on openai engine

![image](docs/img/8f0b4567-c96f-468a-a7c4-57ae3a62dbb4.gif)


# Requirements

* Go 1.19.7
* Node v16.13+
* MySQL 5.7+
* Antd Pro v5
* Antd v5
* umijs v4

# Installation

## Backend
Install using `go`...
    
    # cat config.yaml
        proxy:
          url: "http://0.0.0.0:0000/send_message"
          token: ""
        db:
          user: ""
          host: ""
          port:
          password: ""
          name: "openai"
    
    # cd migrations/ && load sql
    
    # run main 
    go run main.go

## Frontend
```bash 
cd openai-fe && tyarn && tyarn start
```
