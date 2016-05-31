

## UAA Client requirements

Must have the following authorities:

* zones.read

```
uaac client add uaa-fe --authorities "uaa.admin clients.secret scim.read zones.read" --authorized_grant_types implicit
uaac curl /identity-zones
```

### uaa on PCFDev

```
uaac target http://uaa.local.pcfdev.io
uaac token client get admin -s admin-client-secret
uaac client update admin --authorities "zones.read clients.read clients.secret clients.write uaa.admin clients.admin scim.write scim.read"
uaac curl /identity-zones
uaac curl /Users
uaac curl /Groups
uaac curl /oauth/clients
```
