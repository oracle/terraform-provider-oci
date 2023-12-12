---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage_carbon_emissions_config"
sidebar_current: "docs-oci-datasource-metering_computation-usage_carbon_emissions_config"
description: |-
  Provides details about a specific Usage Carbon Emissions Config in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_usage_carbon_emissions_config
This data source provides details about a specific Usage Carbon Emissions Config resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the configuration list for the UI drop-down list of carbon emission console.


## Example Usage

```hcl
data "oci_metering_computation_usage_carbon_emissions_config" "test_usage_carbon_emissions_config" {
	#Required
	tenant_id = oci_metering_computation_tenant.test_tenant.id
}
```

## Argument Reference

The following arguments are supported:

* `tenant_id` - (Required) tenant id


## Attributes Reference

The following attributes are exported:

* `items` - The list of available configurations.
	* `key` - The configuration key.
	* `values` - The configuration value.

