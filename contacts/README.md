# Contacts Microservice

## Requirements

- Contacts MicroService
- A contact information to be maintained.(Name, Email,Address, ContactNo)
- It is a RESTful service
- methods to implement
  - Create(*models.Contact) (interface{}, error)
	- GetBy(id string) (*models.Contact, error)
	- UpdateBy(id string, data map[string]interface{}) (interface{}, error)
	- DeleteBy(id string) (interface{}, error)
- Database is postgres
  - The database design is interface based so change of database must not impact the project
  - Messagebroker, kafka to be used
  - Other microservice features can be added.
  - Dockerize the Microservice.
  
  ## Connecting to database

  - Step-0:install docker from docker.com
  - https://docs.docker.com/engine/install/ubuntu/
  - step-1 ```docker pull postgres```
  - step-2 ```docker run -d --name pg -e POSTGRES_PASSWORD=admin123 -e POSTGRES_USER=postgres -e POSTGRES_DB=contactsdb -p 5432:5432 postgres```
  - step-3: ensure postgres container is running ```docker ps```

- To up and run kafka

```cd kafka && docker-compose up -d```

- Create a topic in kafka. The below command must run inside the kafka container.
  
  ```docker exec -it kafka_kafka_1 bash```
  ```kafka-topics --create --bootstrap-server  localhost:9092 --replication-factor 1 --partitions 1 --topic contacts-create```

- to run consumer
- ```go run consumer.go```

- This application uses both gin and echo framewors
- To use gin
- ```go run main.go --webframework=gin``` or ```go run main.go```
- To use echo
- ```go run main.go --webframework=echo```
- To use both
- When using both it takes default port give it to gin and for echo it increments the port number by 1. example port=58080; For gin it is 58080 and for echo it is 58081
- ```go run main.go --webframework=both```