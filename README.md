### require
    go 1.7
    unoconv

### run 
    $ gvm pkgset create --local
    
    $ gvm pkgset use --local
    
    $ cd src/unoconv
    
    $ go run server.go
    
    
### docker

    docker-compose build
    
    docker-compose up