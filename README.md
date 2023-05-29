# What is it?

Just a project sample I am using for my talk on DDD in Go. Partially based on what guys from ThreeDotsLabs are talking about in their blogposts.

Uses a simple docker compose setup with a single Kafka service.

API definitions, be it openapi, or protofiles, are stored under *api* folder.

https://github.com/bufbuild/buf is used to execute proto related commands, so that all tools are same across all developers.

# Services

* **users-api** - emulates some user database
* **orders-api** - emulates storage for orders
* **email-notifier** - sends some email messages based on events emitted by other services
