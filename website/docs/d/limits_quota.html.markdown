---
subcategory: "Limits"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_limits_quota"
sidebar_current: "docs-oci-datasource-limits-quota"
description: |-
  Provides details about a specific Quota in Oracle Cloud Infrastructure Limits service
---

# Data Source: oci_limits_quota
This data source provides details about a specific Quota resource in Oracle Cloud Infrastructure Limits service.

Gets the quota for the OCID specified.

## Example Usage

```hcl
data "oci_limits_quota" "test_quota" {
	#Required
	quota_id = oci_limits_quota.test_quota.id
}
```

## Argument Reference

The following arguments are supported:

* `quota_id` - (Required) The OCID of the quota.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the resource this quota applies to. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the quota.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the quota.
* `name` - The name you assign to the quota during creation. The name must be unique across all quotas in the tenancy and cannot be changed. 
* `state` - The quota's current state.
* `statements` - An array of one or more quota statements written in the declarative quota statement language.
* `time_created` - Date and time the quota was created, in the format defined by RFC 3339. Example: `2016-08-25T21:10:29.600Z` 

