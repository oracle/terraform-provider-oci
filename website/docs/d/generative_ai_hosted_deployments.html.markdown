---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_hosted_deployments"
sidebar_current: "docs-oci-datasource-generative_ai-hosted_deployments"
description: |-
  Provides the list of Hosted Deployments in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_hosted_deployments
This data source provides the list of Hosted Deployments in Oracle Cloud Infrastructure Generative AI service.

Lists the hosted applications in a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_hosted_deployments" "test_hosted_deployments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	application_id = oci_dataflow_application.test_application.id
	display_name = var.hosted_deployment_display_name
	id = var.hosted_deployment_id
	state = var.hosted_deployment_state
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted deployment.
* `state` - (Optional) A filter to return only the hosted deployments that their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `hosted_deployment_collection` - The list of hosted_deployment_collection.

### HostedDeployment Reference

The following attributes are exported:

* `active_artifact` - Container/artifact configuration for the deployment.
	* `artifact_type` - The type of the artifact.
	* `container_uri` - image url.
	* `hosted_deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	* `id` - if put artifact to a table, the id is needed
	* `status` - The current status of the artifact.
	* `tag` - image tag.
	* `time_created` - The date and time the artifact was created.
* `artifacts` - array of Artifacts.
	* `artifact_type` - The type of the artifact.
	* `container_uri` - image url.
	* `hosted_deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	* `id` - if put artifact to a table, the id is needed
	* `status` - The current status of the artifact.
	* `tag` - image tag.
	* `time_created` - The date and time the artifact was created.
* `compartment_id` - The compartment OCID to create the hosted application in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `hosted_application_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted deployment.
* `state` - The current state of the hosted deployment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time the hosted deployment was created, in the format defined by RFC 3339
* `time_updated` - The date and time the hosted deployment was updated, in the format defined by RFC 3339
