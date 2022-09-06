---
title: Connect with Auth0 through SAML
---

This guide shows how to enable login with ZITADEL on Auth0.

It covers how to:

- create and configure the application in your project
- create and configure the connection in your Auth0 tenant

Prerequisites:

- existing ZITADEL Instance, if not present follow [this guide](../../guides/start/quickstart)
- existing ZITADEL Organization, if not present follow [this guide](../../guides/manage/console/organizations)
- existing ZITADEL project, if not present follow the first 3 steps [here](../../guides/manage/console/projects)
- existing Auth0 tenant as described [here](https://auth0.com/docs/get-started/auth0-overview/create-tenants)

> We have to switch between ZITADEL and a Auth0. If the headings begin with "ZITADEL" switch to the ZITADEL Console and
> if the headings start with "Auth0" please switch to the Auth0 GUI.

## **Auth0**: Create a new connection

In Authentication -> Enterprise
![Navigation Authentication Enterprise](/img/saml/auth0/auth_enterprise.png)

1. Press the "+" button right to "SAML"  
   ![Enterprise Connections](/img/saml/auth0/enterprise_connections.png)
2. Fill out the fields as follows in the SAML Connection:
   ![New SAML Connection](/img/saml/auth0/connection.png)

This includes:

- a unique "Connection name"
- the "Sign In URL"
- the "Sign Out URL"
- used "User ID Attribute"
- the definition how the request should be signed
- which binding should be used to call ZITADEL

All the information is filled out as an example, and to connect with any other environment you only have to change the
used domain, for example "example.com" with "zitadel.cloud".
Lastly, upload the certificate used to sign the reponses, provided for you under the
URL [https://accounts.zitadel.cloud/saml/certificate](https://accounts.zitadel.cloud/saml/certificate).

Then just press the button "Create" and the connection on Auth0 is configured.

## **ZITADEL**: Create the application

So that ZITADEL also recognises the now created SAML
connection, [here](https://auth0.com/docs/authenticate/protocols/saml/saml-identity-provider-configuration-settings) are
all necessary information to correctly fill out the metadata or download the metadata-file directly under the
URL https://YOUR_DOMAIN/samlp/metadata?connection=YOUR_CONNECTION_NAME, which in this example would
be https://example.auth0.com/samlp/metadata?connection=SAML-ZITADEL .

WIP - create SAML application on you instance with the metadata from auth0

## **Auth0**: Try the connection

To then test the connection you only have to press "Try" on the created connection in the Authentication -> Enterprise
screen.

![Authentication Enterprise Try](/img/saml/auth0/auth_enterprise_try.png)

To further customize the requests you can also customize the SAML communication as
described [here](https://auth0.com/docs/authenticate/protocols/saml/saml-configuration/customize-saml-assertions)