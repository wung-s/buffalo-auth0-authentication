This repository is set up as part of a [blog post](https://medium.com/@wung_s/integrate-auth0-with-buffalo-for-api-authentication-44e36140df7a) descriping how [Auth0](http://auth0.com) can be integrated into a [Buffalo](http://gobuffalo.io) app for API authentication

Head over to the blog to follow along. Happy coding !!

### Steps to run locally

Note: You will have to set up your Auth0 first

```
$ git clone git@github.com:wung-s/buffalo-auth0-authentication.git

$ cd buffalo-auth0-authentication && buffalo db create -a

$ dep ensure

$ buffalo dev
```

[Powered by Buffalo](http://gobuffalo.io)
