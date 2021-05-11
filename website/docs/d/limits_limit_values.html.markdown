---
subcategory: "Limits"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_limits_limit_values"
sidebar_current: "docs-oci-datasource-limits-limit_values"
description: |-
  Provides the list of Limit Values in Oracle Cloud Infrastructure Limits service
---

# Data Source: oci_limits_limit_values
This data source provides the list of Limit Values in Oracle Cloud Infrastructure Limits service.

Includes a full list of resource limits belonging to a given service.


## Example Usage

```hcl
data "oci_limits_limit_values" "test_limit_values" {
	#Required
	compartment_id = var.tenancy_ocid
	service_name = oci_limits_service.test_service.name

	#Optional
	availability_domain = var.limit_value_availability_domain
	name = var.limit_value_name
	scope_type = var.limit_value_scope_type
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) Filter entries by availability domain. This implies that only AD-specific values are returned. 
* `compartment_id` - (Required) The OCID of the parent compartment (remember that the tenancy is simply the root compartment). 
* `name` - (Optional) Optional field, can be used to see a specific resource limit value.
* `scope_type` - (Optional) Filter entries by scope type.
* `service_name` - (Required) The target service name.


## Attributes Reference

The following attributes are exported:

* `limit_values` - The list of limit_values.

### LimitValue Reference

The following attributes are exported:

* `availability_domain` - If present, the returned value is only specific to this availability domain.
* `name` - The resource limit name. To be used for writing policies (in case of quotas) or other programmatic calls. 
* `scope_type` - The scope type of the limit. 
* `value` - The resource limit value.

