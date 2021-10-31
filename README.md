# MVP Ecommerce Project with Golang API Server

Sample REST API build using echo server.

The code implementation was inspired by port and adapter pattern or known as [hexagonal](blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example):

-   **Business**<br/>Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
-   **Modules**<br/>Contains implementation of interfaces that defined at the business (also called as server-side adapters in hexagonal's term)
-   **Controller**<br/>Controller http handler or api (also called user-side adapters in hexagonal's term)

![golang clean architecture](https://github.com/favians/go-hexagonal-gorm/raw/master/Hexagonal.png)

# Data initialization

To describe about how port and adapter interaction (separation concerned), this example will have two databases supported. There are MySQL using gorm as library.

# How To Run Server

Just execute code below in your console

```console
go run app/main.go
```

# ERD Design
![golang ERD Design](https://github.com/Alta-FAPA-Store/fapa-store/blob/main/ERD%20Online%20Shop.jpeg)

## Sytem Design
![golang Sytem Design](https://github.com/Alta-FAPA-Store/fapa-store/blob/main/System%20design%20diagram.jpeg)

