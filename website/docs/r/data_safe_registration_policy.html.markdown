---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_registration_policy"
sidebar_current: "docs-oci-resource-data_safe-registration_policy"
description: |-
  Provides the Registration Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_registration_policy
This resource provides the Registration Policy resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/RegistrationPolicy

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Creates a new OptIn/Registration Policy

## Example Usage

```hcl
resource "oci_data_safe_registration_policy" "test_registration_policy" {
	#Required
	compartment_id = var.compartment_id
	features = var.registration_policy_features
	resource_id = oci_cloud_guard_resource.test_resource.id

	#Optional
	can_override_features = var.registration_policy_can_override_features
	connection_option {
		#Required
		connection_type = var.registration_policy_connection_option_connection_type
		identifiers = var.registration_policy_connection_option_identifiers
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.registration_policy_description
	display_name = var.registration_policy_display_name
	freeform_tags = {"Department"= "Finance"}
	opc_dry_run = var.registration_policy_opc_dry_run
	x_cluster_id = var.registration_policy_x_cluster_id
}
```

## Argument Reference

The following arguments are supported:

* `can_override_features` - (Optional) (Updatable) Indicates whether features will be overridden for all targets.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the registration policy will be created.
* `connection_option` - (Optional) (Updatable) Types of connection supported by Data Safe.
	* `connection_type` - (Required) (Updatable) The connection type used to connect to the database. Allowed values:
		* PRIVATE_ENDPOINT - Represents connection through private endpoint in Data Safe.
		* ONPREM_CONNECTOR - Represents connection through on-premises connector in Data Safe.
	* `identifiers` - (Required) (Updatable) List of OCIDs required to establish the connection.
		* For `PRIVATE_ENDPOINT`, provide the OCID(s) of Data Safe private endpoint(s).
		* For `ONPREM_CONNECTOR`, provide the OCID(s) of on-premises connector(s).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) (Updatable) A description of the registration policy.
* `display_name` - (Optional) (Updatable) The display name of the registration policy.
* `features` - (Required) (Updatable) The Data Safe features granted to the databases registering under the registration policy.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `opc_dry_run` - (Optional) Indicates that the request is a dry run, if set to "true". A dry run request does not modify the configuration item details and is used only to perform validation on the submitted data.
* `resource_id` - (Required) The OCID of the resource used in the registration policy.
* `x_cluster_id` - (Optional) Identifier of the cluster of the CDB associated to the registration policy being created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `can_override_features` - Indicates whether features will be overridden.
* `compartment_id` - The OCID for the compartment containing the registration policy.
* `connection_option` - Types of connection supported by Data Safe.
	* `connection_type` - The connection type used to connect to the database. Allowed values:
		* PRIVATE_ENDPOINT - Represents connection through private endpoint in Data Safe.
		* ONPREM_CONNECTOR - Represents connection through on-premises connector in Data Safe.
	* `identifiers` - List of OCIDs required to establish the connection.
		* For `PRIVATE_ENDPOINT`, provide the OCID(s) of Data Safe private endpoint(s).
		* For `ONPREM_CONNECTOR`, provide the OCID(s) of on-premises connector(s).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}`
* `description` - A description of the registration policy.
* `display_name` - The display name of the registration policy.
* `enablement_level` - The resource type which has been opted in for the registration policy.
	* DATABASE - The registration policy will be opted in at the Container Database level.
* `features` - The Data Safe features granted to the databases registering under the registration policy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `id` - The OCID of the registration policy.
* `lifecycle_state_details` - Details about the lifecycle state of the registration policy
* `resource_id` - The OCID of the resource used in the registration policy.
* `state` - The lifecycle state of the registration policy.
	* CREATING - The registration policy is getting created.
	* ACTIVE  - The registration policy has been successfully created.
	* UPDATING - The registration policy is getting updated.
	* NEEDS_ATTENTION - The registration policy needs attention.
	* FAILED - The registration policy creation or update has failed.
	* DELETING - The registration policy is getting deleted.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time when the registration policy was created.
* `time_updated` - The date and time when the registration policy was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Registration Policy
	* `update` - (Defaults to 20 minutes), when updating the Registration Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Registration Policy


## Import

RegistrationPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_registration_policy.test_registration_policy "id"
```
