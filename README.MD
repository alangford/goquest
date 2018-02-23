# goquest

This is a slackbot built with go, that will create jira issues for you


## Getting Started

### Prerequisites

Must have [Go](https://golang.org/doc/install)


### Installing

Clone the repo

```
git clone https://github.com/bolty2142/goquest.git

```

## Development

To run this app in development you will need to create a slack api token

If you perfer to run this outside of Docker you need to export these variables in your .bash_profile

``` 
export SLACK_API_TOKEN=change me // your slack api token
export JIRA_URI=https://yourcompany.atlassian.net/ // your companys jira uri
export JIRA_USER=change me // jira username to login
export JIRA_PASS=change me // jira password to login

```

And then run

``` go build main.go ```

And then run the main file

Or to run inside a docker container run 

```
docker build -t goslackbot .

```

And then to run it in a contontainer use the commande 

```

docker run -e SLACK_API_TOKEN=changeme -e JIRA_URI=https://yourcompany.atlassian.net/ -e JIRA_USER=changeme -e JIRA_PASS=changeme goslackbot

```

## Production

To run this app in production you can use various methods of deployment

I love using AWS ECS to run all of my docker containers


## Built With

* [Go](https://golang.org/) - The language used
* [Docker](https://www.docker.com/) - The container used



## Authors

* **Andrew Langford** - *initial work* - [GitHub Link](https://github.com/bolty2142)

See also the list of [contributors](https://github.com/bolty2142/goquest/graphs/contributors) who participated in this project

