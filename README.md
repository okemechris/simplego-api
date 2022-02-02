# Go Keycloak
Simple API using Keycloak for Auth


## Installing

```bash
go get github.com/Joshswooft/go-keycloak
```

## Getting Started

To get started 1st copy the example environment file into a private .env file using the below command.

```bash
cp .env.example .env
```

This project uses keycloak in docker to help you setup easily. The underlying database can be changed but currently its set to postgres.

```bash
docker-compose up -d
```

This should set up a server running keycloak on `http://localhost:8080/`.

Next navigate to: `http://localhost:8080/auth/admin` and login using the values you set in `.env` for `KEYCLOAK_USER` and `KEYCLOAK_PASSWORD`.

Follow along the steps on this [page](https://www.keycloak.org/getting-started/getting-started-docker) to set up your 1st realm, add a user and then add a client.

Remember to set the following in the console to what you have in your `.env` file: 

- `CLIENT_ID`
- `CLIENT_SECRET`
- `REALM`

## Auth

Authorization is handled by keycloak via an interceptor. Keycloak is an open source library which provides many useful features such as: 
- Single sign on
- Standard Protocols (OpenID Connect, OAuth 2.0 and SAML 2.0)
- Centralized management
- Social login
- User federation and more. 

Check out the full list here: https://www.keycloak.org/


## Domains
This project follows a domain driven design approach.

### Event

Example event object: 
```json
{
    "id": "1",
    "title": "my event",
    "description": "test event"
}
```

- `POST` `/event` creates an event
- `GET` `/event/{id}` get an event
- `GET` `/events/` gets all the events


