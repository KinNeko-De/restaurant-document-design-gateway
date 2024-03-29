# restaurant-document-design-gateway
This service is a api gatewy for designing restaurant documents.
It is part of the [restaraunt project](https://github.com/KinNeko-De/restaurant).

# Automation
Every time you push something a [ci pipeline](.github/workflows/ci.yml) will run. As developers, we are working to automate processes and thus eliminate manual work. There's no excuse why we have to do things manually.

[![restaurant-document-design-gateway-ci](https://github.com/KinNeko-De/restaurant-document-design-gateway/actions/workflows/ci.yml/badge.svg)](https://github.com/KinNeko-De/restaurant-document-design-gateway/actions/workflows/ci.yml)

# Logging
The application uses structured logging. It logs to the console for usage in kubernetes.

# Metrics
Is out of scope for now.

# Tests
Write tests whenever they are needed and are useful.

Never try to spare time here. You have to invest more time later.

Never write useless tests to get high code coverage.

[![codecov](https://codecov.io/gh/KinNeko-De/restaurant-document-design-gateway/branch/main/graph/badge.svg?token=BoDmQQ8ol7)](https://codecov.io/gh/KinNeko-De/restaurant-document-design-gateway)

Tests need to be maintained as normal code. Keep that in mind while you write tests.

# Rate limit for access for demonstration
As it is planned to provide a public demonstration, a github account is needed to access the endpoints. Each github account can generate up to three document every hour.

The application needs access to your github user email to identify you.