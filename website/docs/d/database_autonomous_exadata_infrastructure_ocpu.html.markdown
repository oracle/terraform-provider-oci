---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_exadata_infrastructure_ocpu"
sidebar_current: "docs-oci-datasource-database-autonomous_exadata_infrastructure_ocpu"
description: |-
  Provides details about a specific Autonomous Exadata Infrastructure Ocpu in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_exadata_infrastructure_ocpu
This data source provides details about a specific Autonomous Exadata Infrastructure Ocpu resource in Oracle Cloud Infrastructure Database service.

Gets details of the available and consumed OCPUs for the specified Autonomous Exadata Infrastructure resource.


## Example Usage

```hcl
data "oci_database_autonomous_exadata_infrastructure_ocpu" "test_autonomous_exadata_infrastructure_ocpu" {
	#Required
	autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_exadata_infrastructure_id` - (Required) The Autonomous Exadata Infrastructure  [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `by_workload_type` - 
	* `adw` - The total number of OCPU cores in use for Autonomous Data Warehouse databases in the infrastructure instance.
	* `atp` - The total number of OCPU cores in use for Autonomous Transaction Processing databases in the infrastructure instance.
* `consumed_cpu` - The total number of consumed OCPUs in the Autonomous Exadata Infrastructure instance.
* `total_cpu` - The total number of OCPUs in the Autonomous Exadata Infrastructure instance.

