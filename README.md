# Base Code Golang DDD Pattern

# How to deploy
1. setup env.json, open file and fill
    - Database (setup db connection)
    - MainSetup (setup server host, name and type)
    - logging (setup logging files)
    - Pubsub (setup publisher and subcriber)
    - internal routing (setup endpoint)
    - external routing (setup host and endpoint for partner)
2. build binary `./build.sh`
3. run binary


# how to use swagger
1. Download swag by using: 

`$ go get -u github.com/swaggo/swag/cmd/swag`

**1.16 or newer**

`$ go install github.com/swaggo/swag/cmd/swag@latest`

2. add param while creating controller function on it
3. add example in each struct model that is created
4. run swag init









