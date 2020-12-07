# Introduction
The application is a backend for an Ecommerce website. It designed using Microservices architecture and has 3 modules: Customers, Orders and Products.   
The architecture is designed to scale for very high web traffic.  

# Architecture
![Architecture](https://user-images.githubusercontent.com/26188281/101307782-44fe3300-386e-11eb-9495-06cdd4875557.jpg)

- Ingress acts as a reverse proxy and load balances HTTP traffic from the web.  
- APIs for interservice communication are using gRPC. gRPC ensures faster communication since it is in binary format and avoids unnecessary marshaling and unmarshalling.  
-  Backend for frontend exposes HTTP REST APIs to the Ingress. It maps the REST APIs to gRPC APIs exposed by individual services.  
- Backend for frontend ensures minimal changes in individual services for changing requirements of frontend applications.  
- Services are written in Golang. Each service maintains its data in a localized database.  
- SQL database is used for Customer service since it has structured data.  
- NoSQL database is used for Products since the data is largely unstructured. A caching mechanism is used as the data for products mostly remains unchanged.  
- Order service has complex processing, which could result in a delay in response. A queuing mechanism is used to avoid potential problems caused by heavy web traffic.   

# Features/ Technology stack

- Golang
- Redis
- gRPC
- HTTP REST
- MongoDB
- MariaDB
- Queuing
- Terraform
- Monitoring
- Unit testing
- Benchmarking
- OAuth 2.0


# Tasks
