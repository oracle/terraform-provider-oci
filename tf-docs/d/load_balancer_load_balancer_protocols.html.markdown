---
layout: "oci"
page_title: "OCI: oci_load_balancer_protocols"
sidebar_current: "docs-oci-datasource-load_balancer-protocols"
description: |-
Provides a list of LoadBalancerProtocols
---
# Data Source: oci_load_balancer_protocols
The LoadBalancerProtocols data source allows access to the list of OCI load_balancer_protocols

Lists all supported traffic protocols.

## Example Usage

```hcl
data "oci_load_balancer_protocols" "test_load_balancer_protocols" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer protocols to list.


## Attributes Reference

The following attributes are exported:

* `protocols` - The list of protocols.

### LoadBalancerProtocol Reference

The following attributes are exported:

* `name` - The name of a protocol.  Example: 'HTTP' 

