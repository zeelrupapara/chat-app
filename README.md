# Chat Application

### Setup Guide in Local
#### 1) Start Redis
```
docker compose up -d
```

#### 2) Set ENV in veriables
```
cp .env.example .env
```

#### 3) Start chat server
```
go run main.go
```
now check localhost:8080