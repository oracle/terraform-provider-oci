---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_configuration"
sidebar_current: "docs-oci-datasource-metering_computation-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the list of config for UI dropdown list


## Example Usage

```hcl
data "oci_metering_computation_configuration" "test_configuration" {
	#Required
	tenant_id = "${oci_metering_computation_tenant.test_tenant.id}"
}
```

## Argument Reference

The following arguments are supported:

* `tenant_id` - (Required) tenant id


## Attributes Reference

The following attributes are exported:

* `items` - The list of available configurations
	* `key` - The key of the config
	* `values` - The value of the config

