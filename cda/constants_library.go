package cda

var propertiesMap = map[string]string{
	name:                       "name",
	customType:                 "custom_type",
	folder:                     "folder",
	deploymentTargets:          "deployment_targets",
	dynamicProperties:          "dynamic",
	customProperties:           "custom",
	description:                "description",
	owner:                      "owner",
	agent:                      "agent",
	application:                "application",
	environment:                "environment",
	loginObject:                "login",
	username:                   "login_infor",
	password:                   "password",
	workflow:                   "workflow",
	pack:                       "package",
	deploymentProfile:          "deployment_profile",
	manualApproval:             "needs_manual_start",
	approver:                   "manual_confirmer",
	schedule:                   "planned_from",
	overrideExistingComponents: "install_mode",
}

const name = "name"
const customType = "type"
const folder = "folder"
const deploymentTargets = "deployment_targets"
const dynamicProperties = "dynamic_properties"
const customProperties = "custom_properties"
const description = "description"
const owner = "owner"
const agent = "agent"
const application = "application"
const environment = "environment"
const loginObject = "login_object"
const deploymentMap = "deployment_map"
const credentials = "credentials"
const username = "username"
const password = "password"
const missingRequiredFieldMessage = "Missing required argument: The argument '%s' is required, but no definition was found."
const workflow = "workflow"
const deploymentProfile = "deployment_profile"
const manualApproval = "manual_approval"
const approver = "approver"
const schedule = "schedule"
const overrideExistingComponents = "override_existing_components"
const pack = "package"
const installationUrl = "installation_url"
const monitorUrl = "monitor_url"
const triggers = "triggers"
const defaultAttributes = "default_attributes"
const cdaServer = "cda_server"
const cdaUser = "user"
const overridesApplication = "overrides_application"
const overridesPackage = "overrides_package"
const overridesWorkflow = "overrides_workflow"
const overridesComponent = "overrides_component"

type EntityType int

const (
	DeploymentProfileType EntityType = 1 + iota
	EnvironmentType
	DeploymentTargetType
	LoginObjectType
)

var environmentVariableMaps = map[string]string{
	cdaServer: "CDA_SERVER",
	cdaUser:   "CDA_USER",
	password:  "CDA_PASSWORD",
}

var routerMaps = map[EntityType]string{
	DeploymentProfileType: "profiles",
	EnvironmentType:       "environments",
	DeploymentTargetType:  "deployment_targets",
	LoginObjectType:       "logins",
}

var entityTypeNameMaps = map[EntityType]string{
	DeploymentProfileType: "Deployment profile",
	EnvironmentType:       "Environment",
	DeploymentTargetType:  "Deployment target",
	LoginObjectType:       "Login object",
}
