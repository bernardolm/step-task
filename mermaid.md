# Mermaid

<http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html>

## By Iman Tumorang inspired by Bob’s Clean Architecture Concept

* <https://hackernoon.com/golang-clean-archithecture-efd6d7c43047>
* <https://hackernoon.com/trying-clean-architecture-on-golang-2-44d615bf8fdf>
* <https://github.com/bxcodec/go-clean-arch>

If Uncle Bob’s Architecture, has 4 layer :

* Entities
* Usecase
* Controller
* Framework & Driver

In my projects, I’m using 4 too :

* Models
* Repository
* Usecase
* Delivery

```mermaid
flowchart RL;
    UseCase--interface-->Repository
    Delivery--interface-->UseCase
    UseCase-->Models

    subgraph Models
    User-->Task
    end

    subgraph Repository ["Repository (db, microservices...)"]
    UserRepository
    TaskRepository
    end

    subgraph UseCase ["UseCase (business)"]
    UserUseCase
    TaskUseCase
    end

    subgraph Delivery ["Delivery (presenter)"]
    subgraph Router
        UserController
        TaskController
    end
    subgraph Console
        UserCommand
        TaskCommand
    end
    end
```

---

## By Ruangyot Nanchiang

* <https://orenrose.medium.com/clean-architecture-in-golang-with-go-kit-e5b716a3b881>
* <https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html>

<img src="assets/CleanArchitecture.jpg" alt="drawing" width="500" style="margin: 25px 0 35px 0"/>


```mermaid
flowchart RL;
    Infrastructures-->Adapeters
    Adapeters-->UseCases
    UseCases-->Entities

    subgraph Entities ["Enterprise Business Rules"]
        Entitiy
    end

    subgraph UseCases ["Application Business Rules"]
        UseCase
    end

    subgraph Adapeters ["Interface Adapeters"]
        Controller
        Presenter
        Gateways
    end

    subgraph Infrastructures ["Frameworks & Drivers (infrastructure)"]
        Device
        Web
        DB
        UI
        Console
        ExternalInterface
    end
```

```mermaid
flowchart RL;
    subgraph "Flow of Control"
    Presenter-->UseCaseOutputPort
    UseCaseInteractor-->UseCaseOutputPort
    UseCaseInteractor-->UseCaseInputPort
    UseCaseInteractor-->Presenter
    UseCaseInteractor-->Controller
    Controller-->UseCaseInputPort
    end
```
