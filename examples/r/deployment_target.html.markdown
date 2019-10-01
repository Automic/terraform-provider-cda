---
layout: "CDA"
page_title: "CDA: cda_deployment_target"
sidebar_current: "docs-cda-resource-deployment-target"
description: |-
  A Deployment Target Resource creates a CDA Deployment Target Entity.
---

# cda_deployment_target

A Deployment Target Resource creates a CDA Deployment Target Entity.

## Example Usage

```hcl
resource "cda_deployment_target" "my_target" {
  name        = "my_target_name"
  description = "description"
  type        = "custom type"
  folder      = "target_folder"
  owner       = "target_owner"
  agent       = "agent_name"

  custom_properties = { 
      prop1 = "value1" 
      prop2 = "value2"
      prop3 = "value3"
      prop4 = "value4" 
  }

  dynamic_properties = {
    prop1 = "value1"
    prop2 = "value2"
    prop3 = "value3"
    prop4 = "value4"
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The deployment target's name
- `description` - (Optional) The deployment target's description
- `type` - (Required) The deployment target's custom type
- `folder` - (Optional) The deployment target's folder
- `owner` - (Optional) The deployment target's owner
- `agent` - (Optional) The agent name
- `dynamic_properties` - (Optional) Map of name and value of deployment target's dynamic properties
- `custom_properties` - (Optional) Map of name and value of deployment target's custom properties