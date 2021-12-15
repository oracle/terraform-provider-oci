---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouse_resource_usage_summary"
sidebar_current: "docs-oci-datasource-opsi-operations_insights_warehouse_resource_usage_summary"
description: |-
  Provides details about a specific Operations Insights Warehouse Resource Usage Summary in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_operations_insights_warehouse_resource_usage_summary
This data source provides details about a specific Operations Insights Warehouse Resource Usage Summary resource in Oracle Cloud Infrastructure Opsi service.

Gets the details of resources used by an Operations Insights Warehouse.
There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.


## Example Usage

```hcl
data "oci_opsi_operations_insights_warehouse_resource_usage_summary" "test_operations_insights_warehouse_resource_usage_summary" {
	#Required
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
}
```

## Argument Reference

The following arguments are supported:

* `operations_insights_warehouse_id` - (Required) Unique Operations Insights Warehouse identifier


## Attributes Reference

The following attributes are exported:

* `cpu_used` - Number of OCPUs used by OPSI Warehouse ADW. Can be fractional. 
* `id` - OPSI Warehouse OCID
* `state` - Possible lifecycle states
* `storage_used_in_gbs` - Storage by OPSI Warehouse ADW in GB. 

