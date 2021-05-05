---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_tables"
sidebar_current: "docs-oci-datasource-core-drg_route_tables"
description: |-
  Provides the list of Drg Route Tables in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_route_tables
This data source provides the list of Drg Route Tables in Oracle Cloud Infrastructure Core service.

Lists the DRG route tables for the specified DRG.

Use the `ListDrgRouteRules` operation to retrieve the route rules in a table.


## Example Usage

```hcl
data "oci_core_drg_route_tables" "test_drg_route_tables" {
	#Required
	drg_id = oci_core_drg.test_drg.id

	#Optional
	display_name = var.drg_route_table_display_name
	import_drg_route_distribution_id = oci_core_drg_route_distribution.test_drg_route_distribution.id
	state = var.drg_route_table_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `import_drg_route_distribution_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution. 
* `state` - (Optional) A filter that only returns matches for the specified lifecycle state. The value is case insensitive. 


## Attributes Reference

The following attributes are exported:

* `drg_route_tables` - The list of drg_route_tables.

### DrgRouteTable Reference

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

