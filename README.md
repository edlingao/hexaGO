# HexaGO (A simple ports & adapter architecture example)

This is a simple example of a calculator app using the ports & adapter architecture.

## How to run

1. Clone the repository
2. Install air `go install github.com/cosmtrek/air@latest`
3. Install templ `go install github.com/a-h/templ/cmd/templ@latest`
4. Run `air` in the root of the project
5. Open your browser and go to `http://localhost:3000`

## Inspiration

This project was inspired by the book Hexagonal Architecture by Alistair Cockburn & Juan Manuel Garrido de Paz (RIP).

The prupose of this project is to show a simple example of how to implement the ports & adapter architecture in a Go project.

## License

Do with this project whatever you want. It's free and open source. No strings attached.

# Introduction

What is hexago?

This is a small implementation of hexagonal architechture (ports and adapters) in golang.

### What is hexagonal architecture?

Hexagonal architecture is a software design pattern that separates the application into different layers. This pattern is also known as ports and adapters pattern.

The main parts of the hexagonal architecture are:

- **Core**: This is the main part of the application. It contains the business logic and the domain model.
- **Ports**: These are the interfaces that the core uses to communicate with the outside world. They are defined in the core and implemented in the adapters.
  - **Drivers**: These are the implementations of the ports. They are responsible for interacting with the external systems (databases, APIs, etc).
  -- **Driving Adapters**: These are the implementations of the drivers. They are responsible for translating the data between the core and the external systems.
- **Adapters**: These are the implementations of the ports. They are responsible for translating the data between the core and the external systems.

The main advantage of the hexagonal architecture is that it makes the application more testable and maintainable. The core is decoupled from the external systems, so it can be tested in isolation. The ports and adapters can be easily replaced with different implementations, so the application can be easily adapted to new requirements.

## How to use hexago?

I like to separate the internal layer into multiple hexagon.
Each hexagon lives in its own folder and has a folder structure like this:

```
<HexagonName>
├── core
│   ├── <HexagonName>.go
│   └── <HexagonName>_test.go
├── ports
│   ├── driven
│   │   ├── for<Verb><Action>.go // e.g forCreatingUser.go
│   │   └── for<Verb><Action>_test.go
│   └── driving
│       ├── for<Verb><Action>.go // e.g forCreatingUser.go
│       └── for<Verb><Action>_test.go
└── adapters
    ├── db.go
    └── calculatorHTTP.go
```

The core folder contains the business logic and the domain model of the hexagon.

The ports folder contains the interfaces that the core uses to communicate with the outside world.

The adapters folder contains the implementations of the ports.

*The ports are a collection of interfaces that limits the interaction between the core and the external systems.*


## How to start

1. Clone the repository
2. Install air `go install github.com/cosmtrek/air@latest`
3. Install templ `go install github.com/a-h/templ/cmd/templ@latest`
4. Run `air` in the root of the project
5. Open your browser and go to `http://localhost:3000`

### Dockerization Todo steps

1. [ ] Change `compose.yml` image name to your image name
2. [ ] Change `nginx.conf` domain name to your domain name
3. [ ] Change `.github/workflows/deploy.yml` docker image name to this same name


