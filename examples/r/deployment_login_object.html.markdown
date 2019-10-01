---
layout: "CDA"
page_title: "CDA: cda_login_object"
sidebar_current: "docs-cda-resource-login-object"
description: |-
  A Login Object Resource creates a CDA Login Object Entity.
---

# cda_deployment_target

A Login Object Resource creates a CDA Login Object Entity.

## Example Usage

```hcl
resource "cda_login_object" "my_login_object" {
  name        = "my_login_object_name"
  description = "description"
  folder      = "folder"
  owner       = "owner"

  credentials = [
    {
      agent      = "String"
      type       = "string"
      username   = "string"
      password   = "string"      
    },
    {
      agent      = "String"
      type       = "string"
      username   = "string"
      password   = "string"      
    }
  ]
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The login object's name
- `description` - (Optional) The login object's description
- `folder` - (Optional) The login object's folder
- `owner` - (Optional) The login object's owner
- `credentials` - (Required) List of login object credentials
- `credentials:agent` - (Required) Agent name, support wildcard, accepts '*' as a value e.g. "all Agents from type WINDOWS".
- `credentials:type` - (Required) credential type: example WINDOWS, UNIX, FTPAGENT, CIT, IA.
- `credentials:username` - (Optional) credential username to login agent.
- `credentials:password` - (Optional) credential password to login agent.