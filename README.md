go-api-server
-----
![](https://yyh-gl.github.io/tech-blog/img/tech-blog/2019/06/go_web_api/dependency_direction3.png)

> https://yyh-gl.github.io/tech-blog/blog/go_web_api/

`Handler Layer -> UseCase Layer -> Domain Layer <- Infra`
```
├── cmd
│   └── api
│       └── main.go
├── domain
│   ├── model
│   │   └── user.go
│   └── repository
│       └── user.go
├── handler
│   └── rest
│       └── user.go
├── infra
│   └── persistence
│       └── user.go
└── usecase
    └── user.go
```

