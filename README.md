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
- [x] [Object Pool](/object_pool/)  
Instantiates and maintains a group of objects instances of the same type.

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
The observer pattern allows a type instance to publish events to other type instances(observers) who wish to be updated when a particular event occurs.
- [x] [Template Method](/template_method/)  
Defines a skeleton class which defers some methods to subclasses.
- [x] [Strategy](/strategy/)  
Enables an algorithm's behavior to be selected at runtime.
- [x] [Chain of Responsibility](/chain_of_responsibility/)  
Avoids coupling a sender to receiver by giving more than object a chance to handle the request.
- [x] [State](/state/)  
Encapsulates varying behavior for the same object based on its internal state;
The main point is create objects representing various states, 
and a context object whose behavior changes as its internal state object changes.
- [x] [Iterator](/iterator/)  
Split complex traversal operations into iterator objects.
- [ ] [Visitor](/visitor/)  
Separates an algorithm from an object on which it operates;
The element object accepts the visitor object so that the visitor object can handle operations on the element object.
The intention is to separate data structure from data manipulation.
To solve the coupling problem between stable data structure and volatile operation.
- [ ] [Memento](/memento/)  
Capture the internal state of an object and save this state outside the object, 
so that the object can be restored to the original state in the future.
- [ ] [Command](/command/)  
Bundles a command and arguments to call later.
- [ ] [Interpreter](/interpreter/)  
Defines a grammatical representation for a language and provides an interpreter to deal with this grammar.
- [ ] [Mediator](/mediator/)  
Mediator connects objects and acts as a proxy.

## Other Patterns

### Synchronization Patterns

- [ ] [Semaphore](/semaphore/)  
Allows controlling access to a common resource.

### Concurrency Patterns

- [ ] [Parallelism](/parallelism/)  
Completes large number of independent tasks.
- [ ] [Bounded Parallelism](/bounded_parallelism/)  
Completes large number of independent tasks with resource limits.
- [ ] [Generators](/generators/)  
Yields a sequence of values one at a time.

### Messaging Patterns

- [ ] [Funnel In](/funnel_in/)  
Funnels tasks to a work sink (e.g. server).
- [ ] [Funnel Out](/funnel_out/)  
Distributes tasks among workers (e.g. producer).

### Stability Patterns

- [ ] [Circuit Breaker](/circuit_breaker/)  
Stops the flow of the requests when requests are likely to fail.

### Profiling Patterns

- [ ] [Profile Timing](/profile_timing/)  
Measure the execution time of some functions for code optimization.
