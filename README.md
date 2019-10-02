Terraform Continuous Deliver Automation Provider
==================
The CDA provider is used to create and configure Entities in your CDA system. See the Usage sections for an introduction to the provider.

Requirements
------------
- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Installation
---------------------
 1. Download the latest compiled binary from [GitHub releases](https://github.com/Automic/terraform-provider-cda/releases).
 2. Unzip/untar the archive.
 3. Move it into 
 
      <b>Linux</b>
    
      ```$HOME/.terraform.d/plugins:```
 
      ```
      $ mkdir -p $HOME/.terraform.d/plugins/linux_amd64
      $ mv terraform-provider-cda_v1.0.0 $HOME/.terraform.d/plugins/terraform-provider-cda_v1.0.0
      ```
    
      ### Windows
    
      ```%APPDATA%\terraform.d\plugins```
      
	
    Or put the binary in the terraform installtion folder.
    
 4. Create your Terraform configurations as normal, and run terraform init:

    ```$ terraform init```
    
    This will find the plugin locally.
    
Authentication
----------------------    
The CDA provider offers the following methods for authentication, in this order:

* Static credentials: Add ```user``` and ```password``` in-line in the CDA provider block.
  
  ```
  provider "cda" {
    cda_server = "${var.cda_server}"
    user       = "${var.cda_user}"
    password   = "${var.cda_password}"
  }
  
* Environment variables: Set the following environment variables ```CDA_SERVER``` ```CDA_USER``` ```CDA_PASSWORD```.

  ```provider "cda" {}```


Example Usage
----------------------
To use the provider you need an Automic CDA account. The following example demonstrates a current basic usage of the provider to create an Environment, Deployment Target, Log-in Object, Deployment Profile in CDA. It also shown how to trigger a Wofklow execution. For detailed resources reference see the ```examples``` folder.

```hcl
# Configuring the provider
provider "cda" {
  cda_server = "${var.cda_server}"
  user       = "${var.cda_user}"
  password   = "${var.cda_password}"  
  
  default_attributes = { // default_attributes can be used to set the 'folder' and 'owner' attributes globally for the template.
    folder = "DEFAULT"
    owner  = "100/AUTOMIC/AUTOMIC"
  }
}

# Creating Environment
resource "cda_environment" "firstEnvironment" {
  name               = "environment_name"
  type               = "Generic"
  folder             = "DEFAULT" // If not specified the value provided in the default_attributes will be used.
  owner              = "100/AUTOMIC/AUTOMIC" // If not specified the value provided in the default_attributes will be used.
  description        = "Description"
  
  dynamic_properties = {
      "property" = "value"
  }
  
  custom_properties  = {}
  
  deployment_targets = ["DB", "Tomcat"] 
}

# Creating Deployment Target
resource "cda_deployment_target" "firstTarget" {
  name        = "target_name"
  type        = "Database JDBC"
  folder      = "DEFAULT" // If not specified the value provided in the default_attributes will be used.
  owner       = "100/AUTOMIC/AUTOMIC" // If not specified the value provided in the default_attributes will be used.
  description = "Description"
  agent       = "WIN01"
}

# Creating Log-in Object
resource "cda_login_object" "firstLoginObject" {
  name        = "login_object_name"
  folder      = "DEFAULT" // If not specified the value provided in the default_attributes will be used.
  owner       = "100/AUTOMIC/AUTOMIC" // If not specified the value provided in the default_attributes will be used.

  credentials = [
    {
      agent      = "*"
      type       = "WINDOWS"
      username   = "Agent_User"
      password   = "automic"
    }
  ]
}

# Creating Deployment Profile
resource "cda_deployment_profile" "firstDeploymentProfile" {
  name         = "deployment_profile_name"
  folder       = "DEFAULT" // If not specified the value provided in the default_attributes will be used.
  owner        = "100/AUTOMIC/AUTOMIC" // If not specified the value provided in the default_attributes will be used.
  application  = "My Application"
  environment  = "${cda_environment.firstEnvironment.name}"
  login_object = "${cda_login_object.firstLoginObject.name}"

  deployment_map = {
    "Component A" = "${cda_deployment_target.firstTarget.name}"
  }
}

# Creating Workflow Execution
resource "cda_workflow_execution" "firstExecution" {
  triggers                     = true
  application                  = "DemoApp" 
  workflow                     = "deploy" 
  package                      = "1" 
  deployment_profile           = "${cda_deployment_profile.firstDeploymentProfile.name}" 
  manual_approval              = "true" 
  approver                     = "100/AUTOMIC/AUTOMIC"
  schedule                     = "2019-12-28T13:44:00Z"  
  override_existing_components = "true"
  
   overrides_component = [
	    {
	       component_name = "webapp"
	       "tomcat/property_name" = "new_value"
	    } 
   ]	
}
```

Building the Provider
----------------------
If you would like to work on the CDA provider, you need [Go](http://www.golang.org) installed on your machine.

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e `$HOME/development/terraform-providers/`).

Clone repository to: `$HOME/development/terraform-providers/`

```sh
$ mkdir -p $HOME/development/terraform-providers/; cd $HOME/development/terraform-providers/
$ git clone https://github.com/Automic/terraform-provider-cda.git
...
```

Enter the provider directory and run `go build` to build the provider

```sh
$ go build
```
