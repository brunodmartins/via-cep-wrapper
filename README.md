# via-cep-wrapper
A simple Via CEP Wrapper to demonstrate GoLang tests usage

## Purpose

Demonstrate how struct services could make 
easy to build and test a Go application.

## Use Case

A simple API that wraps [ViaCEP](http://viacep.com.br), a REST API to search CEP information.

```curl localhost:8080/via_cep_wrapper/03010-000```

## Code

The code is spplited on two branchs: `using_functions` and `using_service_struct`

### using_functions

- The code is built upon functions only
- There is only data structs
- All tests depend on inner layers 
- This leads to complex integration tests
- To test the handler layer, the code must mock the Http client of viaCEP

### using_service_struct

- The code is built upon abstractions (interfaces) and implementations (structs)
- All tests depend on mocks, built with [GoMock](https://github.com/golang/mock)
- This leads to simple unit tests
- To test the handler layer, the code must only mock the layer below, in that case, the Service layer

## Conclusion

Both code do their job but with different approaches. 

On my previous experiences, the functions approach worked quite well with small and simple code bases, but adding 
external dependencies like APIs or Databases,  or even growing the code, made the test to complex to maintain and understand.

Nowadays, I prefer the struct service approach, it has proved that for even large code base works quite well
