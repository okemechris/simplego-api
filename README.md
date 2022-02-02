# Go Keycloak
Simple API using Keycloak for Auth


## Installing

```bash
go get github.com/Joshswooft/go-keycloak
```

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


