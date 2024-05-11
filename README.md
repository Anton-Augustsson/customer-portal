# Customer portal
An example application of a scalable customer portal developed in Golang using gin for the backend and Angular for the frontend.

## Quick start
### Requirements
- Docker 
- Protobuf and gRPC
- Kind
- Kubectl
- Helm


### Generate Client and server stub
```bash
cd backend
./update-grpc.sh
```

### Build images
All microservices needs currently needs to be build manually.
An example is shown how it works for subscription-service.
```bash
cd backend/subscription-service
docker build . --tag subscription-service
kind load docker-image subscription-service:latest
```

### Deploy
```bash
helm install customer-portal-test customer-portal
```

Currently, when deploying with kind the port is not exposed to the 
host, so you need to run an exposure command manually.
You need to first find the name of the pod you want to expose.
```bash
kubectl get pods
```
Then you can use the pod to expose the port.
```bash
kubectl port-forward robot-remote-controller-deployment-7b775fd47c-n9x59 8182:50051
```
Note: exchange `robot-remote-controller-deployment-7b775fd47c-n9x59` with 
your pod.



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


#### gRPC
- [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Quick start](https://grpc.io/docs/languages/go/quickstart/)


### API design
- [Better API Design with OpenAPI (Cloud Next '18)](https://www.youtube.com/watch?v=uBs6dfUgxcI)
- [RESTful web API design](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)
- [Google Cloud API design](https://cloud.google.com/apis/design)


### GO
- [Effective Go](https://go.dev/doc/effective_go#interfaces_and_types)


### NGINX
- [NGINX: Full Example Configuration](https://www.nginx.com/resources/wiki/start/topics/examples/full/)
- [NGINX Tutorial for Beginners](https://www.youtube.com/watch?v=9t9Mp0BGnyI)
- [CREATE AN ANGULAR APP AND DEPLOY USING IN NGINX AND DOCKER.](https://www.arunyadav.in/codehacks/blogs/post/19/create-an-angular-app-and-deploy-using-in-nginx-and-docker)


### Angular
- [Angular Basics: How To Use HttpClient in Angular](https://www.telerik.com/blogs/angular-basics-how-to-use-httpclient) 
- [ng-openapi-gen: An OpenAPI 3 code generator for Angular](https://www.npmjs.com/package/ng-openapi-gen)
- [HTTP: Setup for server communication](https://angular.io/guide/http-setup-server-communication)


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
- [Set up a LiveSwitch Angular Project](https://developer.liveswitch.io/liveswitch-server/get-started/js/set-up-a-liveswitch-angular-project.html)


### Further reading

### OpenAPI
- [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator)
- [Documentation for the spring Generator](https://openapi-generator.tech/docs/generators/spring/)
- [Open API Server Implementation Using OpenAPI Generator](https://www.baeldung.com/java-openapi-generator-server)
- [Spring Documentation Overview](https://docs.spring.io/spring-boot/docs/current/reference/html/documentation.html#documentation)
- [OpenAPI data types](https://swagger.io/docs/specification/data-models/data-types/)
- [OpenAPI Guide](https://swagger.io/docs/specification/about/)
- [OpenAPI 3.0 Tutorial](https://support.smartbear.com/swaggerhub/docs/en/get-started/openapi-3-0-tutorial.html)
- [OpenAPI Specification v3.1.0](https://spec.openapis.org/oas/v3.1.0)
- Swagger Viewer for VSCode: Allows you to view the OpenAPI specification in a more intuitive way. It also provides linting.


### Spring boot
An alternative to Golang for microservices is JAVA using the spring boot framework.
However, it is slower and have become unessesaraly more complex over the year,
and arguably Golang is more suited for microservices.

- [Spring Boot Tutorial | Full Course [2023] [NEW]](https://www.youtube.com/watch?v=9SGDpanrc8U)
