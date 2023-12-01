---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_deployment_management"
sidebar_current: "docs-oci-resource-data_safe-security_policy_deployment_management"
description: |-
  Provides the Security Policy Deployment Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_security_policy_deployment_management
This resource provides the Security Policy Deployment Management resource in Oracle Cloud Infrastructure Data Safe service.

Updates the security policy deployment.

## Example Usage

```hcl
resource "oci_data_safe_security_policy_deployment_management" "test_security_policy_deployment_management" {
	#Required
	compartment_id = var.compartment_id
	target_id = oci_data_safe_target_database.test_target_database.id
	
	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.security_policy_deployment_management_description
	display_name = var.security_policy_deployment_management_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the security policy deployment.
* `target_id` - (Required) Unique target identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the security policy deployment.
* `display_name` - (Optional) (Updatable) The display name of the security policy deployment. The name does not have to be unique, and it is changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `security_policy_deployment_id` - (Required) The OCID of the security policy deployment resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the security policy deployment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security policy deployment.
* `display_name` - The display name of the security policy deployment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security policy deployment.
* `lifecycle_details` - Details about the current state of the security policy deployment in Data Safe.
* `security_policy_id` - The OCID of the security policy corresponding to the security policy deployment.
* `state` - The current state of the security policy deployment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target where the security policy is deployed.
* `time_created` - The time that the security policy deployment was created, in the format defined by RFC3339.
* `time_updated` - The last date and time the security policy deployment was updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Policy Deployment Management
	* `update` - (Defaults to 20 minutes), when updating the Security Policy Deployment Management
	* `delete` - (Defaults to 20 minutes), when destroying the Security Policy Deployment Management


## Import

Import is not supported for this resource.

