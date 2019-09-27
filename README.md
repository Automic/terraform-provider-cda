Terraform Continuous Deliver Automation Provider
==================
The CDA provider is used to create and configure Entities in your CDA system. See the Usage and Documentation sections for an introduction to the provider.

Requirements
------------
- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Installation
---------------------

Usage
----------------------

To use a released provider in your Terraform environment, run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and Terraform will automatically install the provider. To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

To instead use a custom-built provider in your Terraform environment (e.g. the provider binary from the build instructions above), follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

For either installation method, documentation about the provider specific configuration options can be found on the [provider's website](https://www.terraform.io/docs/providers/aws/index.html).

## Provider Documentation

The provider is usefull in creating, updating Environment entity of CDA.

### Example
```hcl
# Configuring the provider
provider "cda" {
  cda_server = "${var.cda_server}"
  user       = "${var.cda_user}"
  password   = "${var.cda_password}"  
}

# Creating Environment
resource "cda_environment" "firstEnvironment" {
  name               = "environment_name"
  folder             = "DEFAULT"
  custom_type        = "Generic"
  dynamic_properties = {}
  custom_properties  = {}
  deployment_targets = []
  description        = "Description Update"
  owner              = "100/AUTOMIC/AUTOMIC"  
}

# Creating Deployment Target
resource "cda_environment" "firstEnvironment" {
}

# Creating Deployment Profile
resource "cda_environment" "firstEnvironment" {
}

# Creating Log-in Object
resource "cda_environment" "firstEnvironment" {
}

# Creating Workflow Execution
resource "cda_environment" "firstEnvironment" {
}
```
