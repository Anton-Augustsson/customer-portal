# Customer portal
An example application of a scalable customer portal developed in Golang using gin for the backend and Angular for the frontend.


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


## Generate server stub
```bash
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/open-api-customer-portal.yaml \
    -g go-gin-server \
    -o /local/out/gin
```
read more [go-gin-server.md](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go-gin-server.md)

## Build
The build and start up is handled with docker compose, so in the root project directory type:
```bash
docker compose up
```
Then you can test by typing `http://localhost:8080/camera` in your web browser.


## Resources
### VS Code
- Swagger Viewer: Allows you to view the OpenAPI specification in a more intuitive way. It also provides linting.


### OpenAPI
- [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator)
- [Documentation for the spring Generator](https://openapi-generator.tech/docs/generators/spring/)
- [Open API Server Implementation Using OpenAPI Generator](https://www.baeldung.com/java-openapi-generator-server)
- [Spring Documentation Overview](https://docs.spring.io/spring-boot/docs/current/reference/html/documentation.html#documentation)
- [OpenAPI data types](https://swagger.io/docs/specification/data-models/data-types/)
- [OpenAPI Guide](https://swagger.io/docs/specification/about/)
- [OpenAPI 3.0 Tutorial](https://support.smartbear.com/swaggerhub/docs/en/get-started/openapi-3-0-tutorial.html)
- [OpenAPI Specification v3.1.0](https://spec.openapis.org/oas/v3.1.0)


### API design
- [Better API Design with OpenAPI (Cloud Next '18)](https://www.youtube.com/watch?v=uBs6dfUgxcI)
- [RESTful web API design](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)


### GO
- [Effective Go](https://go.dev/doc/effective_go#interfaces_and_types)


### NGINX
- [NGINX: Full Example Configuration](https://www.nginx.com/resources/wiki/start/topics/examples/full/)
- [NGINX Tutorial for Beginners](https://www.youtube.com/watch?v=9t9Mp0BGnyI)


### Live streaming service
- [How Does Live Streaming Platform Work? (YouTube live, Twitch, TikTok Live)](https://www.youtube.com/watch?v=7AMRfNKwuYo)
- [Streams API](https://developer.mozilla.org/en-US/docs/Web/API/Streams_API)
- [libav RTMP streaming](https://github.com/juniorxsound/libav-RTMP-Streaming)
- [Go Video Transcoder - Govitra](https://github.com/bqqbarbhg/go-video-transcoder)
- [transcoder](https://pkg.go.dev/cloud.google.com/go/video/transcoder/apiv1beta1)
- [goDASH — GO Accelerated HAS Framework for Rapid Prototyping](https://ieeexplore.ieee.org/abstract/document/9123103)
- [go-segment](https://github.com/cwinging/go-segment)
- [go-ts-segmenter](https://github.com/jordicenzano/go-ts-segmenter)
- [livego](https://github.com/qieangel2013/livego)
- [goDash Application](https://github.com/uccmisl/godash)
- [Angular video transformations](https://cloudinary.com/documentation/angular_video_transformations)


### Further reading
#### gRPC
AN alternatively to a REST API is gRPC which is said to provide better performance then REST API. It is much more modern then REST API and was initial created by google in 2015 but is now managed by Cloud Native Computing Foundation.

### Spring boot
An alternative to Golang for microservices is JAVA using the spring boot framework.
However, it is slower and have become unessesaraly more complex over the year,
and arguably Golang is more suited for microservices.

- [Spring Boot Tutorial | Full Course [2023] [NEW]](https://www.youtube.com/watch?v=9SGDpanrc8U)
