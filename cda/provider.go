package cda

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			cdaServer: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CDA_URL", ""),
				Description: descriptions[cdaServer],
			},
			cdaUser: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CDA_USER", ""),
				Description: descriptions[cdaUser],
			},
			password: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CDA_PASSWORD", ""),
				Description: descriptions[password],
			},
			defaultAttributes: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: descriptions[defaultAttributes],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cda_environment":        resourceCdaEnvironment(),
			"cda_deployment_target":  resourceCdaDeploymentTarget(),
			"cda_deployment_profile": resourceCdaDeploymentProfile(),
			"cda_login_object":       resourceCdaLoginObject(),
			"cda_workflow_execution": resourceCdaWorkflowExecution(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var folderStr string
	var ownerStr string
	if attributes, found := d.GetOk(defaultAttributes); found {
		var attributesMap = attributes.(map[string]interface{})
		if attributesMap[folder] != nil {
			folderStr = attributesMap[folder].(string)
		}
		if attributesMap[owner] != nil {
			ownerStr = attributesMap[owner].(string)
		}
	}

	config := Config{
		CdaServer: d.Get(cdaServer).(string),
		User:      d.Get(cdaUser).(string),
		Password:  d.Get(password).(string),
		Folder:    folderStr,
		Owner:     ownerStr,
	}

	return config.checkConnection()
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		cdaServer:         "This is the CDA server name for CDA API operations.",
		cdaUser:           "This is the fully qualified username needed to perform CDA API operations. E.g. 100/Admin/IT",
		password:          "This is the password for CDA API operations.",
		defaultAttributes: "This is the default attribute for CDA resource: folder and owner",
	}
}
