---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_limits_resource_availability"
sidebar_current: "docs-oci-datasource-limits-resource_availability"
description: |-
  Provides details about a specific Resource Availability in Oracle Cloud Infrastructure Limits service
---

# Data Source: oci_limits_resource_availability
This data source provides details about a specific Resource Availability resource in Oracle Cloud Infrastructure Limits service.

For a given compartmentId, resource limit name, and scope, returns the following:
  - the number of available resources associated with the given limit
  - the usage in the selected compartment for the given limit
  Note: not all resource limits support this API. If the value is not available, the API will return 404.


## Example Usage

```hcl
data "oci_limits_resource_availability" "test_resource_availability" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	limit_name = "${var.resource_availability_limit_name}"
	service_name = "${oci_limits_service.test_service.name}"

	#Optional
	availability_domain = "${var.resource_availability_availability_domain}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) This field is mandatory, if the scopeType of the target resource limit is AD. Otherwise, this field should be omitted. If the above are not respected, the API will return a 400 - InvalidParameter response. 
* `compartment_id` - (Required) The OCID of the compartment for which data is being fetched.
* `limit_name` - (Required) The limit name for which to fetch the data.
* `service_name` - (Required) The service name of the target quota.


## Attributes Reference

The following attributes are exported:

* `available` - The count of available resources. 
* `used` - The current usage in the given compartment. 

