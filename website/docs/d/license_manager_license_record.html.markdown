---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_license_record"
sidebar_current: "docs-oci-datasource-license_manager-license_record"
description: |-
  Provides details about a specific License Record in Oracle Cloud Infrastructure License Manager service
---

# Data Source: oci_license_manager_license_record
This data source provides details about a specific License Record resource in Oracle Cloud Infrastructure License Manager service.

Retrieves license record details by the license record ID in a given compartment.

## Example Usage

```hcl
data "oci_license_manager_license_record" "test_license_record" {
	#Required
	license_record_id = oci_license_manager_license_record.test_license_record.id
}
```

## Argument Reference

The following arguments are supported:

* `license_record_id` - (Required) Unique license record identifier.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the license record is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The license record display name. Avoid entering confidential information. 
* `expiration_date` - The license record end date in [RFC 3339](https://tools.ietf.org/html/rfc3339) date format. Example: `2018-09-12` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The license record [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `is_perpetual` - Specifies if the license record term is perpertual.
* `is_unlimited` - Specifies if the license count is unlimited.
* `license_count` - The number of license units added by the user for the given license record. Default 1 
* `license_unit` - The product license unit.
* `product_id` - The license record product ID.
* `product_license` - The product license name with which the license record is associated.
* `product_license_id` - The product license [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) with which the license record is associated.
* `state` - The current license record state.
* `support_end_date` - The license record support end date in [RFC 3339](https://tools.ietf.org/html/rfc3339) date format. Example: `2018-09-12` 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the license record was created. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.
* `time_updated` - The time the license record was updated. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.

