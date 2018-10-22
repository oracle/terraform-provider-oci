---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_services"
sidebar_current: "docs-oci-datasource-core-services"
description: |-
  Provides the list of Services in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_services
This data source provides the list of Services in Oracle Cloud Infrastructure Core service.

Lists the available services that you can access through a service gateway in this region.


## Example Usage

```hcl
data "oci_core_services" "test_services" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `services` - The list of services.

### Service Reference

The following attributes are exported:

* `cidr_block` - A string that represents the public endpoints for the service. When you set up a route rule to route traffic to the service gateway, use this value as the destination CIDR block for the rule. See [Route Table](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/RouteTable/). 
* `description` - Description of the service. 
* `id` - The service's [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - Name of the service.

