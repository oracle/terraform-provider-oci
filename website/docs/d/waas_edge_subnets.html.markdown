---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_edge_subnets"
sidebar_current: "docs-oci-datasource-waas-edge_subnets"
description: |-
  Provides the list of Edge Subnets in Oracle Cloud Infrastructure Waas service
---

# Data Source: oci_waas_edge_subnets
This data source provides the list of Edge Subnets in Oracle Cloud Infrastructure Waas service.

Return the list of the tenant's edge node subnets. Use these CIDR blocks to restrict incoming traffic to your origin. These subnets are owned by Oracle Cloud Infrastructure and forward traffic to customer origins. They are not associated with specific regions or compartments.

## Example Usage

```hcl
data "oci_waas_edge_subnets" "test_edge_subnets" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `edge_subnets` - The list of edge_subnets.

### EdgeSubnet Reference

The following attributes are exported:

* `cidr` - An edge node subnet. This can include /24 or /8 addresses.
* `region` - The name of the region containing the indicated subnet.
* `time_modified` - The date and time the last change was made to the indicated edge node subnet, expressed in RFC 3339 timestamp format.

