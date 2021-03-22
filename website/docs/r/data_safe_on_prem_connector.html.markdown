---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_on_prem_connector"
sidebar_current: "docs-oci-resource-data_safe-on_prem_connector"
description: |-
  Provides the On Prem Connector resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_on_prem_connector
This resource provides the On Prem Connector resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new on-premises connector.


## Example Usage

```hcl
resource "oci_data_safe_on_prem_connector" "test_on_prem_connector" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.on_prem_connector_description
	display_name = var.on_prem_connector_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the on-premises connector.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the on-premises connector.
* `display_name` - (Optional) (Updatable) The display name of the on-premises connector. The name does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `available_version` - Latest available version of the on-premises connector.
* `compartment_id` - The OCID of the compartment that contains the on-premises connector.
* `created_version` - Created version of the on-premises connector.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the on-premises connector.
* `display_name` - The display name of the on-premises connector.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the on-premises connector.
* `lifecycle_details` - Details about the current state of the on-premises connector.
* `state` - The current state of the on-premises connector.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the on-premises connector was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the On Prem Connector
	* `update` - (Defaults to 20 minutes), when updating the On Prem Connector
	* `delete` - (Defaults to 20 minutes), when destroying the On Prem Connector


## Import

OnPremConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_on_prem_connector.test_on_prem_connector "id"
```

