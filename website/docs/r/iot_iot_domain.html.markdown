---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_domain"
sidebar_current: "docs-oci-resource-iot-iot_domain"
description: |-
  Provides the Iot Domain resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_iot_domain
This resource provides the Iot Domain resource in Oracle Cloud Infrastructure Iot service.

Creates a new IoT domain.


## Example Usage

```hcl
resource "oci_iot_iot_domain" "test_iot_domain" {
	#Required
	compartment_id = var.compartment_id
	iot_domain_group_id = oci_iot_iot_domain_group.test_iot_domain_group.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.iot_domain_description
	display_name = var.iot_domain_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the resource. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `iot_domain_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
* `data_retention_periods_in_days` - Data Retention periods
	* `historized_data` - Number of days for which any normalized data sent to IoT devices would be retained for.
	* `raw_command_data` - Number of days for which any raw command data sent to IoT devices would be retained for.
	* `raw_data` - Number of days for which any raw data sent to IoT devices would be retained for.
	* `rejected_data` - Number of days for which any data sent to IoT devices would be retained for.
* `db_allow_listed_identity_group_names` - List of IAM groups of form described in [here](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/mnqmn/#GUID-3634D6C9-A7F1-4875-9925-BAEA2D3C5197) that are allowed to directly connect to the data host.
* `db_allowed_identity_domain_host` - Host name of identity domain that is used for authenticating connect to data host via ORDS.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. 
* `device_host` - Host name of an IoT domain, where IoT devices can connect to.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `iot_domain_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain group.
* `state` - The current state of the IoT domain.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Iot Domain
	* `update` - (Defaults to 20 minutes), when updating the Iot Domain
	* `delete` - (Defaults to 20 minutes), when destroying the Iot Domain


## Import

IotDomains can be imported using the `id`, e.g.

```
$ terraform import oci_iot_iot_domain.test_iot_domain "id"
```

