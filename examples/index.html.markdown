---
layout: "cda"
page_title: "Provider: CDA"
sidebar_current: "docs-cda-index"
description: |-
  The Continuous Delivery Automation (CDA) provider is used to interact with the many resources supported by CDA. The provider needs to be configured with the proper credentials before it can be used.
---

# CDA Provider

The Continuous Delivery Automation (CDA) provider is used to interact with the
many resources supported by CDA. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the CDA Provider
provider "cda" {
  user = "client/username/department" // "100/AUTOMIC/AUTOMIC
  password  = "password"
  cda_server = "cda_enpoint"  // "http://192.168.1.100/CDA
  
  default_attributes = {
    folder = "Default"
    owner  = "${var.cda_user}"
  }

}

# Create an Environment
resource "cda_environment" "my_environment" {
  name        = "my_environment_name"
  description = "description"
  type        = "environment_custom_type"
  folder      = "environment_folder"
  owner       = "environment_owner"
  deployment_targets = ["target_name1", ..., "target_nameX"]
}
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the CDA
 `provider` block:

* `user` - (Required) This is the fully qualified username needed to perform CDA API operations. E.g. 100/Admin/IT.

* `password` - (Required) This is the password for CDA API operations.

* `cda_server` - (Required) This is the CDA server name for CDA API operations..

* `default_attributes` - (Optional) This the default value for the "folder", "owner" attributes which are propagated to all resources in case the resource doesn't override them.
