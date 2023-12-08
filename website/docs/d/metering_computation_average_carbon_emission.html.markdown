---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_average_carbon_emission"
sidebar_current: "docs-oci-datasource-metering_computation-average_carbon_emission"
description: |-
  Provides details about a specific Average Carbon Emission in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_average_carbon_emission
This data source provides details about a specific Average Carbon Emission resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the average carbon emissions summary by SKU.


## Example Usage

```hcl
data "oci_metering_computation_average_carbon_emission" "test_average_carbon_emission" {
	#Required
	sku_part_number = var.average_carbon_emission_sku_part_number
}
```

## Argument Reference

The following arguments are supported:

* `sku_part_number` - (Required) The SKU part number.


## Attributes Reference

The following attributes are exported:

* `average_carbon_emission` - The average carbon emissions by SKU.
* `sku_part_number` - The sku part number.

