# Builder
**Builder** is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.
## Steps
1. Declare `Builder` interface consists of steps to build the product.
2. `Concrete Builders` implements builder to produce products.
3. `Products` are resulting object built by using concrete builders.
4. `Director` accepts builder and call construction steps.
5. `Client` must associate one builder with one director.