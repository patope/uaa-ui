#uaa-ui: A front end for UAA


## Setup

Depends on a running UAA instance. You can easily run a local instance of UAA with minimal setup required by following [these instructions](https://github.com/cloudfoundry/uaa#quick-start).

The CLIENT_ID used to connect to that running UAA server requires the following authorities: `uaa.admin clients.secret scim.read zones.read`. The best way to configure this client is to use the UAA CLI. For example:
```
# Target your running UAA server
uaac target http://localhost:8080/uaa
# Use the admin client to create your uaa-ui
uaac token client get admin -s adminsecret
uaac client add admin --authorities "zones.read clients.read clients.secret clients.write uaa.admin clients.admin scim.write scim.read"
uaac curl /identity-zones
```

The uaa-ui requires the following environment variables:

* UAA_URL - i.e. http://localhost:8080/uaa, https://uaa.local.pcfdev.io
* UAA_CLIENT_ID - i.e. uaa-ui
* UAA_CLIENT_SECRET - i.e. uaa-ui-secret

To run locally, use the convenience script `./runlocal`
