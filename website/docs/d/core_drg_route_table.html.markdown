---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_table"
sidebar_current: "docs-oci-datasource-core-drg_route_table"
description: |-
  Provides details about a specific Drg Route Table in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_route_table
This data source provides details about a specific Drg Route Table resource in Oracle Cloud Infrastructure Core service.

Gets the specified DRG route table's information.

## Example Usage

```hcl
data "oci_core_drg_route_table" "test_drg_route_table" {
	#Required
	drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id
}
```

## Argument Reference

The following arguments are supported:

* `drg_route_table_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DRG is in. The DRG route table is always in the same compartment as the DRG. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG that contains this route table. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table. 
* `import_drg_route_distribution_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements from referenced attachments are inserted into the DRG route table. 
* `is_ecmp_enabled` - If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises network, enable ECMP on the DRG route table to which these attachments import routes. 
* `state` - The DRG route table's current state.
* `time_created` - The date and time the DRG route table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

