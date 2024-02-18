# Customer portal
An example application of a scalable customer portal developed in Java Spring Boot for the backend and Angular for the frontend.


## The example
The functionality that the system provides is viewing a camera on a remote device, e.g. a 
raspberry pi. Each customer has a subscriber account and they have a subscription to 
each of the cameras they want to view, i.e. one subscription for one camera.
Both using the remote camera viewer and managing subscriptions is accessible through 
the customer portal.


## Domain concepts (resources)
- Subscriber 
- Subscription
- Camera


## Resources
### VS Code
- Swagger Viewer: Allows you to view the OpenAPI specification in a more intuitive way. It also provides linting.


### OpenAPI
- [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator)
- [Open API Server Implementation Using OpenAPI Generator](https://www.baeldung.com/java-openapi-generator-server)
- [Spring Documentation Overview](https://docs.spring.io/spring-boot/docs/current/reference/html/documentation.html#documentation)
- [OpenAPI data types](https://swagger.io/docs/specification/data-models/data-types/)
- [OpenAPI Guide](https://swagger.io/docs/specification/about/)
- [OpenAPI 3.0 Tutorial](https://support.smartbear.com/swaggerhub/docs/en/get-started/openapi-3-0-tutorial.html)
- [OpenAPI Specification v3.1.0](https://spec.openapis.org/oas/v3.1.0)


### API design
- [Better API Design with OpenAPI (Cloud Next '18)](https://www.youtube.com/watch?v=uBs6dfUgxcI)
- [RESTful web API design](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)


### Further reading
#### gRPC
AN alternatively to a REST API is gRPC which is said to provide better performance then REST API. It is much more modern then REST API and was initial created by google in 2015 but is now managed by Cloud Native Computing Foundation.