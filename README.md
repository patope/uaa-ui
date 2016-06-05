#uaa-ui: A front end for UAA


## Setup

Depends on a running UAA instance. You can easily run a local instance of UAA with minimal setup required by following [these instructions](https://github.com/cloudfoundry/uaa#quick-start).

The CLIENT_ID used to connect to that running UAA server requires the following authorities: `uaa.admin clients.secret scim.read zones.read`


Requires the following environment variables:

* UAA_URL - i.e. http://localhost:8080/uaa, http://uaa.local.pcfdev.io
* CLIENT_ID
* CLIENT_SECRET

To run locally, use the convenience script `./runlocal`



##
