---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_exadata_infrastructure_shapes"
sidebar_current: "docs-oci-datasource-database-autonomous_exadata_infrastructure_shapes"
description: |-
  Provides the list of Autonomous Exadata Infrastructure Shapes in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_exadata_infrastructure_shapes
This data source provides the list of Autonomous Exadata Infrastructure Shapes in Oracle Cloud Infrastructure Database service.

**Deprecated.** 


## Example Usage

```hcl
data "oci_database_autonomous_exadata_infrastructure_shapes" "test_autonomous_exadata_infrastructure_shapes" {
	#Required
	availability_domain = var.autonomous_exadata_infrastructure_shape_availability_domain
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the Availability Domain.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_exadata_infrastructure_shapes` - The list of autonomous_exadata_infrastructure_shapes.

### AutonomousExadataInfrastructureShape Reference

The following attributes are exported:

* `available_core_count` - The maximum number of CPU cores that can be enabled on the Autonomous Exadata Infrastructure.
* `core_count_increment` - The increment in which core count can be increased or decreased.
* `maximum_node_count` - The maximum number of nodes available for the shape.
* `minimum_core_count` - The minimum number of CPU cores that can be enabled on the Autonomous Exadata Infrastructure.
* `minimum_node_count` - The minimum number of nodes available for the shape.
* `name` - The name of the shape used for the Autonomous Exadata Infrastructure.

