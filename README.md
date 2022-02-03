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


### Client with a client secret
First create a client on the console inside the realm you have made. 

Inside the settings tab you want to change "Access Type" to `confidential` then hit save. Refresh the page and you will see new tabs. Click on credentials and you should see a client secret which you can then copy into your `.env`.

Note: set the "Valid redirect URIs" to `*`.

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
    "ID": 1,
    "CreatedAt": "2022-02-03T10:18:40.797376Z",
    "UpdatedAt": "2022-02-03T10:18:40.797376Z",
    "DeletedAt": null,
    "title": "1st event",
    "description": "this is my 1st event with keycloak"
}
```

To make any of this endpoints work we need to be an authenticated user. The quickest way to do that is first create a user on your realm in the admin console. Then go over to postman and create a new request. Inside the authorization tab set the `Grant Type` to `Password credentials`. Then fill out the details including your clientId, clientSecret, username and password for the user you just created. 

The access token url should look like this: `http://localhost:8080/auth/realms/example/protocol/openid-connect/token`. 

Set the `Client Authentication` to be "send as Basic Auth header" then generate your token. 

Note: On the admin console in the client settings make you turn on the setting: `Direct Access Grants Enabled` to enable this flow. 


- `POST` `/event` creates an event
- `GET` `/event/{id}` get an event
- `GET` `/events/` gets all the events


