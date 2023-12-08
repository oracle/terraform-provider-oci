---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_clean_energy_usage"
sidebar_current: "docs-oci-datasource-metering_computation-clean_energy_usage"
description: |-
  Provides details about a specific Clean Energy Usage in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_clean_energy_usage
This data source provides details about a specific Clean Energy Usage resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the clean energy usage summary by region.


## Example Usage

```hcl
data "oci_metering_computation_clean_energy_usage" "test_clean_energy_usage" {
	#Required
	region = var.clean_energy_usage_region

	#Optional
	ad = var.clean_energy_usage_ad
}
```

## Argument Reference

The following arguments are supported:

* `ad` - (Optional) The availability domain.
* `region` - (Required) The region.


## Attributes Reference

The following attributes are exported:

* `ad` - The availability domain.
* `region` - The region.
* `usage` - The percentage of clean enery used.

