---
subcategory: "Em Warehouse"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_em_warehouse_em_warehouse_resource_usage"
sidebar_current: "docs-oci-datasource-em_warehouse-em_warehouse_resource_usage"
description: |-
  Provides details about a specific Em Warehouse Resource Usage in Oracle Cloud Infrastructure Em Warehouse service
---

# Data Source: oci_em_warehouse_em_warehouse_resource_usage
This data source provides details about a specific Em Warehouse Resource Usage resource in Oracle Cloud Infrastructure Em Warehouse service.

Gets a EmWarehouseResourceUsage by identifier

## Example Usage

```hcl
data "oci_em_warehouse_em_warehouse_resource_usage" "test_em_warehouse_resource_usage" {
	#Required
	em_warehouse_id = oci_em_warehouse_em_warehouse.test_em_warehouse.id
}
```

## Argument Reference

The following arguments are supported:

* `em_warehouse_id` - (Required) unique EmWarehouse identifier


## Attributes Reference

The following attributes are exported:

* `em_instance_count` - EmInstanceCount
* `em_instances` - List of emInstances
	* `em_discoverer_url` - emdDiscoverer url
	* `em_host` - emHost name
	* `em_id` - operations Insights Warehouse Identifier
	* `targets_count` - EmInstance Target count
* `id` - Unique identifier that is immutable on creation
* `operations_insights_warehouse_id` - operations Insights Warehouse Identifier
* `schema_name` - schema name
* `targets_count` - EmInstance Target count

