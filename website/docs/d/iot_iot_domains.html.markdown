---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_domains"
sidebar_current: "docs-oci-datasource-iot-iot_domains"
description: |-
  Provides the list of Iot Domains in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_iot_domains
This data source provides the list of Iot Domains in Oracle Cloud Infrastructure Iot service.

Retrieves a list of IoT domains within the specified compartment.


## Example Usage

```hcl
data "oci_iot_iot_domains" "test_iot_domains" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.iot_domain_display_name
	id = var.iot_domain_id
	iot_domain_group_id = oci_iot_iot_domain_group.test_iot_domain_group.id
	state = var.iot_domain_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) Filter resources whose display name matches the specified value. 
* `id` - (Optional) Filter resources by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type. 
* `iot_domain_group_id` - (Optional) Filter resources that match the specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain group.
* `state` - (Optional) Filter resources whose lifecycleState matches the specified value. 


## Attributes Reference

The following attributes are exported:

* `iot_domain_collection` - The list of iot_domain_collection.

### IotDomain Reference

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

