# go-design-patterns [![Go Report Card](https://goreportcard.com/badge/github.com/windvalley/go-design-patterns)](https://goreportcard.com/report/github.com/windvalley/go-design-patterns) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=windvalley_go-design-patterns&metric=alert_status)](https://sonarcloud.io/dashboard?id=windvalley_go-design-patterns) [![codecov](https://codecov.io/gh/windvalley/go-design-patterns/branch/main/graph/badge.svg?token=UV7V4WC03R)](https://codecov.io/gh/windvalley/go-design-patterns)

Design patterns implemented in Go

## Creational Patterns

- [x] [Simple Factory](/simple_factory/)  
When you need an object, just pass a correct parameter to the function, and you can get the object you need without knowing the details of its creation.
- [x] [Factory Method](/factory_method/)  
Defers instantiation of an object to a specialized function for creating instances.
- [x] [Abstract Factory](/abstract_factory/)  
Provides an interface for creating families of releated objects.
- [x] [Singleton](/singleton/)  
Restricts instantiation of a type to one object.
- [x] [Builder](/builder/)  
Split a large object into multiple small objects, then assemble multiple small objects into large objects, and hide the construction process from the outside.
- [x] [Prototype](/prototype/)  
Return a new instance by copying an existing instance, rather than creating a new one, which is mostly used to create complex or time-consuming instances.

## Structural Patterns

- [x] [Proxy](/proxy/)  
Provides a surrogate for an object to control it's actions, for example, to delay its actions or to perform other processing before and after its actions.
- [x] [Bridge](/bridge/)  
Decouples an interface from its implementation so that the two can vary independently.
- [x] [Decorator](/decorator/)  
Adds behavior to an object, statically or dynamically.
- [x] [Adapter](/adapter/)  
Convert incompatible interfaces to compatible interfaces so that objects that cannot work together due to incompatible interfaces can work together.
- [x] [Facade](/facade/)  
Uses one type as an API to a number of others.
- [x] [Composite](/composite/)  
Sometimes it is more directly called the tree pattern, which is used to unify the access of leaf node objects and non-leaf node objects.
- [x] [Flyweight](/flyweight/)  
Reuses existing instances of objects with similar/identical state to minimize resource usage.

## Behavioral Patterns

- [x] [Observer](/behavioral/)  
Provide a callback for notification of events/changes to data.
- [ ] [Template Method](/template_method/)  
Defines a skeleton class which defers some methods to subclasses.
- [ ] [Strategy](/strategy/)  
Enables an algorithm's behavior to be selected at runtime.
- [ ] [Chain of Responsibility](/chain_of_responsibility/)  
Avoids coupling a sender to receiver by giving more than object a chance to handle the request.
- [ ] [State](/state/)  
Encapsulates varying behavior for the same object based on its internal state.
- [ ] [Iterator](/iterator/)  
Split complex traversal operations into iterator objects.
- [ ] [Visitor](/visitor/)  
Separates an algorithm from an object on which it operates.
- [ ] [Memento](/memento/)  
Generate an opaque token that can be used to go back to a previous state.
- [ ] [Command](/command/)  
Bundles a command and arguments to call later.
- [ ] [Interpreter](/interpreter/)  
Defines a grammatical representation for a language and provides an interpreter to deal with this grammar.
- [ ] [Mediator](/mediator/)  
Connects objects and acts as a proxy.

## Other Patterns

- [ ] [Profile Timing](/profile_timing/)  
Measure the execution time of some functions for code optimization.
