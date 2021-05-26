---
subcategory: "Limits"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_limits_services"
sidebar_current: "docs-oci-datasource-limits-services"
description: |-
  Provides the list of Services in Oracle Cloud Infrastructure Limits service
---

# Data Source: oci_limits_services
This data source provides the list of Services in Oracle Cloud Infrastructure Limits service.

Returns the list of supported services.
This includes the programmatic service name, along with the friendly service name.


## Example Usage

```hcl
data "oci_limits_services" "test_services" {
	#Required
	compartment_id = var.tenancy_ocid
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the parent compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `services` - The list of services.

### Service Reference

The following attributes are exported:

* `description` - The friendly service name.
* `name` - The service name. Use this when calling other APIs.

