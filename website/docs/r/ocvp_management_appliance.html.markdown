---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_management_appliance"
sidebar_current: "docs-oci-resource-ocvp-management_appliance"
description: |-
  Provides the Management Appliance resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# oci_ocvp_management_appliance
This resource provides the Management Appliance resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Creates a management appliance.

## Example Usage

```hcl
resource "oci_ocvp_management_appliance" "test_management_appliance" {
	#Required
	configuration {
		#Required
		is_log_ingestion_enabled = var.management_appliance_configuration_is_log_ingestion_enabled
		is_metrics_collection_enabled = var.management_appliance_configuration_is_metrics_collection_enabled

		#Optional
		metrics = var.management_appliance_configuration_metrics
		support_bundle_bucket_id = oci_objectstorage_bucket.test_bucket.id
	}
	connections {
		#Required
		credentials_secret_id = oci_vault_secret.test_secret.id
		type = var.management_appliance_connections_type
	}
	display_name = var.management_appliance_display_name
	sddc_id = oci_ocvp_sddc.test_sddc.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	public_ssh_keys = var.management_appliance_public_ssh_keys
}
```

## Argument Reference

The following arguments are supported:

* `configuration` - (Required) (Updatable) Configuration of management appliance.
	* `is_log_ingestion_enabled` - (Required) (Updatable) Is log ingestion from SDDC to Oracle Cloud Infrastructure enabled.
	* `is_metrics_collection_enabled` - (Required) (Updatable) Is metrics collection and publishing is enabled for appliance.
	* `metrics` - (Optional) (Updatable) Array of metrics ids to collect.
	* `support_bundle_bucket_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of support bundle Object Storage bucket. 
* `connections` - (Required) (Updatable) Array of connections for management appliance.
	* `credentials_secret_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of secret in Oracle Cloud Infrastructure vault, that is used for storage of username and password in JSON format. 
	* `type` - (Required) (Updatable) Type of connection.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A descriptive name for the management appliance. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `public_ssh_keys` - (Optional) One or more public SSH keys to be included in `~/.ssh/authorized_keys` file for Management Appliance compute instance. Several public SSH keys must be separate by newline character. 
* `sddc_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of SDDC in OCI, that this appliance is going to be registered in. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment in OCI, that this appliance is going to be created in. 
* `compute_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compute instance of management appliance in OCI. 
* `configuration` - Configuration of management appliance.
	* `is_log_ingestion_enabled` - Is log ingestion from SDDC to Oracle Cloud Infrastructure enabled.
	* `is_metrics_collection_enabled` - Is metrics collection and publishing is enabled for appliance.
	* `metrics` - Array of metrics ids to collect.
	* `support_bundle_bucket_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of support bundle Object Storage bucket. 
* `connections` - Array of connections for management appliance.
	* `credentials_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of secret in Oracle Cloud Infrastructure vault, that is used for storage of username and password in JSON format. 
	* `type` - Type of connection.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the management appliance. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `heartbeat_connection_states` - Current states of connections.
	* `details` - Information about current connection status.
	* `state` - Current connection status.
	* `type` - Type of connection.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of management appliance. 
* `lifecycle_details` - Information about current lifecycleState. For FAILED and NEEDS_ATTENTION contains explanations. For other states may contain some details about their progress.
* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of management agent, that this appliance is running in. 
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of SDDC in OCI, that this appliance is going to be registered in. 
* `state` - Current state of the management appliance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_configuration_updated` - The date and time the configuration of management appliance was last updated in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_created` - The date and time the management appliance was created in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_last_heartbeat` - The date and time the management appliance has last received heartbeat in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the management appliance was last updated in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Appliance
	* `update` - (Defaults to 20 minutes), when updating the Management Appliance
	* `delete` - (Defaults to 20 minutes), when destroying the Management Appliance


## Import

ManagementAppliances can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_management_appliance.test_management_appliance "id"
```

