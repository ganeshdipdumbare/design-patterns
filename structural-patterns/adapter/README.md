# Adapter
**Adapter** is a structural design pattern that allows objects with incompatible interfaces to collaborate.
## Steps
1. `Client` is a class that contains business logic.
2. `Client interface` that other class must follow to collaborate with client code.
3. `Service` useful class, usually 3rd party or legacy code.
4. `Adapter` is a class which contains Service object. It implements client interface to make sure that client is able to communicate with the service.
5. Client code doesnt get coupled with concrete adapter as it just defines the interface which should be implemented by adapter. 