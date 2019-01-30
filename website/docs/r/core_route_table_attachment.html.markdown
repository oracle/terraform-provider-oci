---
layout: "oci"
page_title: "OCI: oci_core_route_table_attachment"
sidebar_current: "docs-oci-resource-core-route-table-attachment"
description: |-
  Provides the ability to associate a route table for a subnet in Oracle Cloud Infrastructure Core service
---

# oci_core_route_table_attachment
This resource provides the ability to associate a route table for a subnet in Oracle Cloud Infrastructure Core service.

Attaches the specified `route table` to the specified `subnet`.

## Example Usage

```hcl
resource "oci_core_route_table_attachment" "test_route_table_attachment" {
  #Required	
  subnet_id = "${oci_core_subnet.test_subnet.id}"
  route_table_id ="${oci_core_route_table.test_route_table.id}"
}
```

## Argument Reference

The following arguments are supported:

* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet.
* `route_table_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table.


## Attributes Reference

The following attributes are exported:

* `subnet_id` - (Required) The OCID of the subnet.
* `route_table_id` - (Required) The OCID of the route table.

## Import

Route Table Attachment can be imported using the `id`, e.g.

```
$ terraform import oci_core_route_table_attachment.test_route_table_attachment "{subnetId}/{routeTableId}" 
```
