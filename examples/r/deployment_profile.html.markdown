---
layout: "CDA"
page_title: "CDA: cda_deployment_profile"
sidebar_current: "docs-cda-resource-deployment-profile"
description: |-
  A Deployment Profile Resource creates a CDA Deployment Profile Entity.
---

# cda_deployment_target

A Deployment Profile Resource creates a CDA Deployment Profile Entity.

## Example Usage

```hcl
resource "cda_deployment_profile" "my_deployment_profile" {
  name         = "my_deployment_profile_name"
  description  = "description"
  folder       = "folder"
  owner        = "owner"
  application  = "application"
  environment  = "environment"
  login_object = "login_object"

  deployment_map = { 
    component1 = "target_name1, ..., target_nameN" 
    component2 = "target_name11, ..., target_nameN1"
    component3 = "target_name111, ..., target_nameN11"
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The deployment profile's name
- `description` - (Optional) The deployment profile's description
- `folder` - (Optional) The deployment profile's folder
- `owner` - (Optional) The deployment profile's owner
- `application` - (Required) The application name
- `environment` - (Required) The environment name
- `login_object` - (Required) The login object name
- `deployment_map` - (Optional) Specifies the component to target mapping to be used for the deployment