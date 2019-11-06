---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_services"
sidebar_current: "docs-oci-datasource-core-services"
description: |-
  Provides the list of Services in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_services
This data source provides the list of Services in Oracle Cloud Infrastructure Core service.

Lists the available [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Service/) objects that you can enable for a
service gateway in this region.


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

* `cidr_block` - A string that represents the regional public IP address ranges for the Oracle service or services covered by this `Service` object. Also known as the `Service` object's *service CIDR label*.

	When you set up a route rule to route traffic to the service gateway, use this value as the rule's destination. See [Route Table](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/RouteTable/). Also, when you set up a security list rule to cover traffic with the service gateway, use the `cidrBlock` value as the rule's destination (for an egress rule) or the source (for an ingress rule). See [Security List](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/SecurityList/).

	Example: `oci-phx-objectstorage` 
* `description` - Description of the Oracle service or services covered by this `Service` object.  Example: `OCI PHX Object Storage` 
* `id` - The `Service` object's [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - Name of the `Service` object. This name can change and is not guaranteed to be unique.  Example: `OCI PHX Object Storage` 

