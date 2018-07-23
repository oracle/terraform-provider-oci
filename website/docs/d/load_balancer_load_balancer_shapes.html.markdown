---
layout: "oci"
page_title: "OCI: oci_load_balancer_shapes"
sidebar_current: "docs-oci-datasource-load_balancer-shapes"
description: |-
Provides a list of LoadBalancerShapes
---
# Data Source: oci_load_balancer_shapes
The LoadBalancerShapes data source allows access to the list of OCI load_balancer_shapes

Lists the valid load balancer shapes.

## Example Usage

```hcl
data "oci_load_balancer_shapes" "test_load_balancer_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer shapes to list.


## Attributes Reference

The following attributes are exported:

* `shapes` - The list of shapes.

### LoadBalancerShape Reference

The following attributes are exported:

* `name` - The name of the shape.  Example: `100Mbps` 

