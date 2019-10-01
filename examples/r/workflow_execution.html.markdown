---
layout: "CDA"
page_title: "CDA: cda_workflow_execution"
sidebar_current: "docs-cda-resource-workflow-execution"
description: |-
  A Workflow Execution Resource execute a CDA Workflow.
---

# cda_environment

A Workflow Execution Resource execute a CDA Workflow.

## Example Usage

```hcl
resource "cda_workflow_execution" "my_execution" {
  triggers                     = true 
  application                  = "application" 
  workflow                     = "workflow name" 
  package                      = "package" 
  deployment_profile           = "deployment_profile" 
  manual_approval              = "true" 
  approver                     = "100/AUTOMIC/AUTOMIC"
  schedule                     = "2019-12-28T13:44:00Z"
  override_existing_components = "false"

  overrides_application = {
    "/Prompt_Dynamic_Name" = "Dynamic_Value"
  }

  overrides_workflow = {
    "Prompt_Dynamic_Float" = "123.12"
  }

  overrides_package = {
    "Promp_Dynamic_Workflow_Ref" = "{\"name\" : \"deploy\", \"application\" : \"DemoApp\"}"
  }

  overrides_component = [
    {
      component_name = "webapp"
      Prompt_Dynamic_Name2 = "1234.45"
    }
  ]
}
```

## Argument Reference

The following arguments are supported:

- `trigger` - (Optional) Default value True.
 
True: execute a CDA workflow when create the resource, or when any of this resource attributes changes

False: does not execute CDA workflow.

- `application` - (Required) The application contains workflow
- `workflow` - (Required) The workflow name
- `package` - (Required) The package name
- `deployment_profile` - (Required) The deployment profile name
- `manual_approval` - (Optional) Default value false.
- `approver` - (Optional) Input required when manual_approval = true
- `schedule` - (Optional) Specify time to start a CDA workflow
- `override_existing_components` - (Optional) default false

 True: override the components
 False: skips existing components
 
 - `overrides_application` - (Optional) overrides prompt dynamic property of application which inputted in the template
 - `overrides_package` - (Optional) overrides prompt dynamic property of package which inputted in the template
 - `overrides_workflow` - (Optional) overrides prompt dynamic property of workflow which inputted in the template
 - `overrides_component` - (Optional) overrides prompt dynamic property of components inside the application.
 
