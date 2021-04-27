# Motivation
This project was built as part of a challenge tht required a showcase of a webpage that contained an api call that would retrieve data from a database in json format. Additionally the project should contain a small ci/cd pipeline that builds an image that subsequently gets uploaded.

I decided to build a very naive payment gateway with no security as a nice showcase that would still meet the challenge's criteria. I hope you enjoy messing around with it and possibly breaking it gloriously.

# Design
There isn't much to this project except a single backend service coupled with a frontend webpage, there are no message buses, no queues or anything. The intent was to keep it as simple as possible. The backend comes with additional tooling to make things easier, it contains the api server but also contains a migration manager that can be run before the container starts to make sure the database is up to date.

```
 ___________
|           |
|  Frontend |
|___________|
      ^
      |
 ___________         ___________
|           |       |           |
|  Backend  | <---> |  Database |
|___________|       |___________|
```

I decided to build this whole thing in go, mainly because its fast and I like it but also because its strongly typed nature leads to less unpredictability and less edge cases.

For the database I considered going with BoltDB because it seems to be all the rage and I wanted to try it but ended up going to old faithful postgres because its the db technology I'm the most comfortable with.

# CI/CD Pipeline
I chose github actions because its free, its fast and has yet to let me down. This project isnt going to get thousands of builds per day so github actions feels like a good fit.
As for the process itself there are 3 steps:
* Test - Runs unit tests for the project
* Build - Builds the docker image
* Upload to ECR - Uploads the docker image to AWS ECR

Since we dont have a deploy step we dont have to concern ourselves too much with dependencies within these jobs so there the upload step only waits for the build step and not the tests.

In a scenario where there are deployments and promotions to other environments, the test step becomes a dependency of upstream steps.

# Security concerns
With the full disclaimer that this projects is strictly a demo and not meant to be used in production, the process to actually get it on production would look something like this:
* The docker container runs in a controlled environment, within a closed VPC
* All configurations are done through environment variables or configuration files, nothing baked in to the container!
* External access is made through a proxy such as nginx, traefik or other such tool
* There is a load balancer in front of all external entrypoints
* Secrets are managed either in cluster with Kubernetes secrets or using an external secure tool like hashicorp vault or aws secrets.
* All connections are encrypted with TLS


# Scaling
There are multiple ways of scaling this project, some of them I already mentioned;

The most naive way of scaling is vertically, a beefy server will be able to answer more requests. This approach isn't viable in the long term though.

Discarding vertical scaling we can horizontaly scale the backend container (or the frontend container for that matter, although frontend technologies typically can stand much higher traffic). By putting a load balancer in front of a group of equal services, the load gets distributed between them and we are able to support a lot of traffic.

But what if we run very complex operations? Another scaling method that builds on top of horizontal scaling is splitting services into tinier microservices which are in charge of a very small amount of work. We can spin up hundreds of these which distributes traffic even more. This brings a lot more complexity and requires technologies like message buses or streaming processors but it is generally an effective scaling strategy.

I did all that but now the problem is the database (or redis or any other such tool) - In this case we can opt for redis clusters in the redis case, database sharding or even distributed databases like Cassandra. This bring a ton more complexity but massive scale requires complex solutions.
