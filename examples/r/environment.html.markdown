---
layout: "CDA"
page_title: "CDA: cda_environment"
sidebar_current: "docs-cda-resource-environment"
description: |-
  The Environment Resource creates a CDA Environment Entity.
---

# cda_environment

The Environment Resource creates a CDA Environment Entity.

## Example Usage

```hcl
resource "cda_environment" "example" {
    name        = "my_environment_name"
    description = "description"
    type        = "environment_custom_type"
    folder      = "environment_folder"
    owner       = "environment_owner"
  
    deployment_targets = ["target_name1", ..., "target_nameX"]
  
    custom_properties = { 
        "prop1" = "value1" 
        "prop2" = "value2"
        "prop3" = "value3"
        "prop4" = "value4" 
    }
  
    dynamic_properties = {
        "name1" = "value1"
        "name2" = "value2"
        "name3" = "value3"
        "name4" = "value4"
    }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The environment's name
- `description` - (Optional) The environment's description
- `type` - (Required) The environment's custom type
- `folder` - (Optional) The environment's folder
- `owner` - (Optional) The environment's owner
- `deployment_targets` - (Optional) List of deployment targets which are assigned to environment
- `dynamic_properties` - (Optional) Map of name and value of environment's dynamic properties
- `custom_properties` - (Optional) Map of name and value of environment's custom properties
