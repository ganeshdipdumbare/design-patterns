# Factory Method
## Steps
1. Declare `Product` interface.
2. Create `Concrete Products` using Product interface.
3. Create `Creator` interface with common operation for all the products and `createProduct():Product` method.
4. Create `Concrete Creator` for each product by implementing the Creator interface.