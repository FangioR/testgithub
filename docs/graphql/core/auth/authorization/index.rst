.. meta::
   :description: Manage authorization with Hasura
   :keywords: hasura, docs, authorization, auth

.. _authorization:

Authorization / Access control
==============================

.. contents:: Table of contents
  :backlinks: none
  :depth: 1
  :local:

Overview
--------

Hasura supports **role-based** authorization where access control is done by creating rules for each role,
table and operation (*insert*, *update*, etc.). These access control rules use dynamic session
variables that are passed to the GraphQL engine from your :ref:`authentication service <authentication>`
with every request. Role information is inferred from the ``X-Hasura-Role`` and ``X-Hasura-Allowed-Roles``
session variables. Other session variables can be passed by your auth service as per your requirements.

**For example:**

.. thumbnail:: /img/graphql/core/auth/hasura-perms.png
   :width: 80 %
   :alt: Create a permission rule

.. note::

  Hasura roles and permissions are implemented at the Hasura layer.
  They have nothing to do with Postgres roles and users.

Trying out access control
-------------------------

If you just want to see role-based access control in action, you need not set up or integrate your
auth service with GraphQL engine. You can just:

* Define permission rules for a table for a role.

* Use the GraphiQL interface in the console to make a request and send the session variables as
  request headers (*send a* ``X-Hasura-Role`` *key, with its value as the name of the role you've
  defined rules for*). The data in the response will be restricted as per your configuration.

Follow the example at :ref:`access control basics <auth_basics>`.


**See:**

.. toctree::
  :maxdepth: 1

  basics
  roles-variables
  inherited-roles
  permission-rules
  common-roles-auth-examples
  role-multiple-rules

.. admonition:: API limits and access controls in Hasura Cloud
  :class: dhc
  
  Additional access controls and API limits like maximum query depth can be found in Hasura Cloud. See more at :ref:`API limits with Hasura Cloud <api_limits>`.

.. admonition:: Additional Resources

  Enterprise Grade Authorization - `Watch Webinar <https://hasura.io/events/webinar/authorization-modeling-hasura/?pg=docs&plcmt=body&cta=watch-webinar&tech=>`__.
