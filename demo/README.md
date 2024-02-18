First build the jar file.
```bash
mvn clean package
```

Then build the docker image.
```bash
docker build --build-arg JAR_FILE=target/*.jar -t myorg/myapp .
```

Lastly create and run the docker container.
```bash
docker run -p 8080:8080 myorg/myapp
```

See more details on [Spring Boot Docker](https://spring.io/guides/topicals/spring-boot-docker).