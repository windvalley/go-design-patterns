# go-design-patterns [![Go Report Card](https://goreportcard.com/badge/github.com/windvalley/go-design-patterns)](https://goreportcard.com/report/github.com/windvalley/go-design-patterns) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=windvalley_go-design-patterns&metric=alert_status)](https://sonarcloud.io/dashboard?id=windvalley_go-design-patterns) [![codecov](https://codecov.io/gh/windvalley/go-design-patterns/branch/main/graph/badge.svg?token=UV7V4WC03R)](https://codecov.io/gh/windvalley/go-design-patterns)

Design patterns implemented in Go

## Creational Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Simple Factory](/simple_factory/) | When you need an object, just pass a correct parameter to the function, and you can get the object you need without knowing the details of its creation | ✔ |
| [Factory Method](/fatory_method/) | Defers instantiation of an object to a specialized function for creating instances | ✔ |
| [Abstract Factory](/abstract_factory/) | Provides an interface for creating families of releated objects | ✔ |
| [Singleton](/singleton/) | Restricts instantiation of a type to one object | ✔ |
| [Builder](/builder/) | Split a large object into multiple small objects, then assemble multiple small objects into large objects, and hide the construction process from the outside | ✘ |
| [Prototype](/prototype/) | Return a new instance by copying an existing instance, rather than creating a new one, which is mostly used to create complex or time-consuming instances | ✘ |
