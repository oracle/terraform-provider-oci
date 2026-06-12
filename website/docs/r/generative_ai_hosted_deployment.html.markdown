---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_hosted_deployment"
sidebar_current: "docs-oci-resource-generative_ai-hosted_deployment"
description: |-
  Provides the Hosted Deployment resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_hosted_deployment
This resource provides the Hosted Deployment resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/HostedDeployment

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates a hosted deployment.

## Example Usage

```hcl
resource "oci_generative_ai_hosted_deployment" "test_hosted_deployment" {
	#Required
	active_artifact {

		#Optional
		artifact_type = var.hosted_deployment_active_artifact_artifact_type
		container_uri = var.hosted_deployment_active_artifact_container_uri
		hosted_deployment_id = oci_generative_ai_hosted_deployment.test_hosted_deployment.id
		id = var.hosted_deployment_active_artifact_id
		status = var.hosted_deployment_active_artifact_status
		tag = var.hosted_deployment_active_artifact_tag
		time_created = var.hosted_deployment_active_artifact_time_created
	}
	hosted_application_id = oci_generative_ai_hosted_application.test_hosted_application.id

	#Optional
	compartment_id = var.compartment_id
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.hosted_deployment_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `active_artifact` - (Required) (Updatable) Container/artifact configuration for the deployment.
	* `artifact_type` - (Optional) (Updatable) The type of the artifact.
	* `container_uri` - (Applicable when artifact_type=SIMPLE_DOCKER_ARTIFACT) (Updatable) image url.
	* `hosted_deployment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	* `id` - (Optional) (Updatable) if put artifact to a table, the id is needed
	* `status` - (Optional) (Updatable) The current status of the artifact.
	* `tag` - (Applicable when artifact_type=SIMPLE_DOCKER_ARTIFACT) (Updatable) image tag.
	* `time_created` - (Optional) (Updatable) The date and time the artifact was created.
* `compartment_id` - (Optional) The compartment OCID to create the hosted deployment in.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `hosted_application_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Hosted Deployment
	* `update` - (Defaults to 20 minutes), when updating the Hosted Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Hosted Deployment


## Import

HostedDeployments can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_hosted_deployment.test_hosted_deployment "id"
```
