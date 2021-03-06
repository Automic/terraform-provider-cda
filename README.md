Terraform Continuous Delivery Automation Provider
==================
The CDA provider is used to create and configure Entities in your CDA system. See the Usage sections for an introduction to the provider.

Requirements
------------
- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Installation
---------------------
 1. Download the latest compiled binary from [releases](https://github.com/Automic/terraform-provider-cda/releases).
 2. Unzip/untar the archive.
 3. Move it into 
 
      <b>Linux</b>
    
      ```$HOME/.terraform.d/plugins:```
 
      ```
      $ mkdir -p $HOME/.terraform.d/plugins/linux_amd64
      $ mv terraform-provider-cda_v1.0.0 $HOME/.terraform.d/plugins/terraform-provider-cda_v1.0.0
      ```
    
      <b>Windows</b>
    
      ```%APPDATA%\terraform.d\plugins```
      
	
    Or put the binary in the terraform installation folder.
    
 4. Create your Terraform configurations as normal, and run terraform init:

    ```$ terraform init```
    
    This will find the plugin locally.
    
Authentication
----------------------    
The CDA provider offers the following methods for authentication, in this order:

* Static credentials: Add ```cda_server```, ```user``` and ```password``` in-line in the CDA provider block.
  
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
The following example demonstrates a current basic usage of the provider to create an Environment, Deployment Target, Log-in Object, Deployment Profile in CDA. It also shows how to trigger a Wofklow execution. For detailed resource reference see the ```examples``` folder.

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

# Terraform Automic Agent Install Provisioner
Automic Agent Install is a Terraform Provisioner which can install CDA Agent and Service Manager on a Red Hat, Ubuntu or Windows instance.

Installation
---------------------
1. Download Terraform Automic Agent Install Provisioner from [releases](https://github.com/Automic/terraform-provider-cda/releases)
2. Extract the downloaded artifact and copy the terraform-provisioner-automic_agent_install binary in the terraform installation folder.
3. The Agent and Service Manager binaries are contained in the Provisioner artifact. They can be downloaded also from the [Automic Marketplace](https://downloads.automic.com/marketplace/browse?search=agent)

Argument Reference
---------------------
|Attribute|Type|Mandatory/Optional|Default Value|Description|
|--|--|--|--|--|
|destination|String|M||The target installation folder.|
|source|String|M||The local folder which contains the Agent and Service Manager artifacts and the install script.|
|agent_name|String|M||The Agent name.|
|agent_port|String|O|2300|The Agent port.|
|ae_host|String|M||The Automation Engine host.|
|ae_port|String|M||The Automation Engine port.|
|ae_system_name|String|M||The Automation Engine name.|
|sm_name|String|M||The Service Manager name|
|sm_host|String|O|8871|The Service Manager host.|
|variables|Map|O|{}|A Terraform map containing override & custom variables which can be used to configure the Agent.|

Example Usage
---------------------

### Precondition
There must be an user already created on the instance where the the Agent and Service Manager will be installed. The user must have sufficient rights to execute the install script.

The following example creates an AWS Unix instance and installs a CDA agent and Service Manager on it.
```
provider "aws" {
  region     = "${var.aws_region}"
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
}

resource "aws_instance" "cda_instance" {
  ami                    = "${var.aws_ami}"
  instance_type          = "${var.instance_type}"
  vpc_security_group_ids = ["${var.aws_security_group_id}"]
  key_name               = "${var.aws_key_name}"


  tags = {
    Name = "example_instance"
  }

  provisioner "automic_agent_install" {
    destination = ""${var.installation_dir}"" //The target directory on the instance, where the installation script and the Agent and Service Manager binaries will be copied.
    source = "C:\\Users\\example_user\\provisoner\\linux" //The folder where the installation script, the Agent and Service Manager binaries are downloaded.

    agent_name = "${random_string.cda_entity_name.result}"
    agent_port = "${var.agent_port}"
    ae_system_name = "${var.ae_system_name}"
    ae_host = "${var.ae_host}"
    ae_port = "${var.ae_port}"
    sm_port = "${var.sm_port}"
    sm_name = "${var.sm_name}"

    variables = {
      UC_EX_PATH_JOBREPORT = "../tmp"   //overrides existing variable
      UC_EX_PATH_TEMP = "../tmp"    //overrides existing variable
      UC_EX_NEW_VARIABLE = "VALUE"    //adds new variable to agent configuration
      UC_EX_IP_ADDR = "${self.public_ip}"   //adds new variable to the agent configuration
    }

    connection {
      host = self.public_ip
      type = "ssh"
      user = "ubuntu"
      private_key = "${file("${var.private_key_file}")}"
    }
  }
}
```

Sample output in case of successful execution:

```
aws_instance.cda_instance (automic_agent_install): Installed Directory: /home/ubuntu/AE
aws_instance.cda_instance (automic_agent_install): Agent Name: 45FIOE1ZXH
aws_instance.cda_instance (automic_agent_install): Agent Port: 2300
aws_instance.cda_instance (automic_agent_install): Service Manager Name: sm_45FIOE1ZXH
aws_instance.cda_instance (automic_agent_install): Service Manager Port: 8871
```

How to install a downloaded version of an Agent and Service Manager
---------------------
Create a new artifact following exactly the below structure:
### Linux
#### Agent
- Copy the files from another x64 Agent installation version and create the following structure:

```bash
ucxjlx6
├── backup
├── bin
│   ├── ucx.msl
│   ├── ucxjlx6
│   ├── ucxjlx6m
│   ├── ucxjlx6s
│   └── ucxjxxx.ori.ini
├── out
├── resources
└── temp
```
- Archive the folder with the name: ucxjlx6.tar.gz
- Replace the existing ucxjlx6.tar.gz in the source folder

#### Service Manager
- Copy the files from aother x64 Service Manager installation version and create the following structure:

```bash
ucsmgrlx6
├── bin
│   ├── automicsm
│   ├── libgcc_s.so
│   ├── libgcc_s.so.1
│   ├── libssp.so
│   ├── libssp.so.0
│   ├── libssp.so.0.0.0
│   ├── libstdc++.so
│   ├── libstdc++.so.6
│   ├── libstdc++.so.6.0.3
│   ├── libsysapi.so
│   ├── libzu00132.so
│   ├── uc.msl
│   ├── uc4.ori.smd
│   ├── ucybsmcl
│   ├── ucybsmgr
│   └── ucybsmgr.ori.ini
└── temp
```
- Archive the folder with the name: ucsmgrlx6.tar.gz
- Replace the existing ucsmgrlx6.tar.gz in the source folder

### Windows
#### Agent
- Copy the files from aother x64 Agent installation version and create the following structure:
```bash
agent
├── trp3.cab
├── SETUP.EXE
├── Setup.lst
├── uc.msl
├── ucsmtp.dll
├── ucxjwx6.exe
├── UCXJWX6.ini
├── ucxjwx6m.exe
├── zu00132.dll
└── zuksp.dll
```
- Archive the folder with the name: agent.zip
- Replace the existing agent.zip in the source folder

#### Service Manager
- Copy the files from aother x64 Service Manager installation version and create the following structure:
```bash
servicemanager
├── bootstrp3.cab
├── SETUP.EXE
├── SETUP.LST
├── uc.msl
├── UC4.smd
├── UCYBSMgr.exe
├── ucybsmgr.ini
└── zu00132.dll
```
- Archive the folder with the name: servicemanager.zip
- Replace the existing servicemanager.zip in the source folder
