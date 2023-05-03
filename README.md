# twiter


1. go mod init

2. dependencies
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get go.mongodb.org/mongo-driver/bson    
 go get go.mongodb.org/mongo-driver/bson/primitive
 go get golang.org/x/crypto/bcrypt
 go get github.com/gorilla/mux
 go get github.com/rs/cors

 3. Conexion a MOngo
    bd/ConexcionBD

4. handlers
    mux --> router para http
    listen and serv
    Port Definition