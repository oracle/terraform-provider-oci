---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_license_metric"
sidebar_current: "docs-oci-datasource-license_manager-license_metric"
description: |-
  Provides details about a specific License Metric in Oracle Cloud Infrastructure License Manager service
---

# Data Source: oci_license_manager_license_metric
This data source provides details about a specific License Metric resource in Oracle Cloud Infrastructure License Manager service.

Retrieves the license metrics for a given compartment.

## Example Usage

```hcl
data "oci_license_manager_license_metric" "test_license_metric" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	is_compartment_id_in_subtree = var.license_metric_is_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration. 
* `is_compartment_id_in_subtree` - (Optional) Indicates if the given compartment is the root compartment.


## Attributes Reference

The following attributes are exported:

* `license_record_expiring_soon_count` - Total number of license records that will expire within 90 days in a particular compartment. 
* `total_byol_instance_count` - Total number of BYOL instances in a particular compartment. 
* `total_license_included_instance_count` - Total number of License Included (LI) instances in a particular compartment. 
* `total_product_license_count` - Total number of product licenses in a particular compartment. 

