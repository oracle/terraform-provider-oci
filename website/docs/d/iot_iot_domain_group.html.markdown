---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_domain_group"
sidebar_current: "docs-oci-datasource-iot-iot_domain_group"
description: |-
  Provides details about a specific Iot Domain Group in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_iot_domain_group
This data source provides details about a specific Iot Domain Group resource in Oracle Cloud Infrastructure Iot service.

Retrieves the IoT domain group identified by the specified OCID.

## Example Usage

```hcl
data "oci_iot_iot_domain_group" "test_iot_domain_group" {
	#Required
	iot_domain_group_id = oci_iot_iot_domain_group.test_iot_domain_group.id
}
```

## Argument Reference

The following arguments are supported:

* `iot_domain_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an IoT Domain Group.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
* `data_host` - The host name of the database corresponding to the IoT Domain group.
* `db_allow_listed_vcn_ids` - This is an array of VCN OCID (virtual cloud network Oracle Cloud ID) that is allowed to connect the data host.
* `db_connection_string` - The connection string used to connect to the data host associated with the IoT domain group.
* `db_token_scope` - The token scope used to connect to the data host associated with the IoT domain group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `state` - The current state of an IoT Domain Group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

