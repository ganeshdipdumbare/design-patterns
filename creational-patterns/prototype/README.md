# Prototype
**Prototype** is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.
## Steps
1. Declare `Prototype` interface which has `clone()` method only in most cases.
2. Create `Concrete Prototype` which implements cloning.
3. The `Client` produce a copy of any object which follows the prototype interface.
### Prototype Registry
1. Registry to store frequently accessed prototypes.
2. Usual prototype registry consists of mapping of `name -> prototype`.