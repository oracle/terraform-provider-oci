---
layout: "oci"
page_title: "OCI: oci_load_balancer_policies"
sidebar_current: "docs-oci-datasource-load_balancer-policies"
description: |-
Provides a list of LoadBalancerPolicies
---
# Data Source: oci_load_balancer_policies
The LoadBalancerPolicies data source allows access to the list of OCI load_balancer_policies

Lists the available load balancer policies.

## Example Usage

```hcl
data "oci_load_balancer_policies" "test_load_balancer_policies" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer policies to list.


## Attributes Reference

The following attributes are exported:

* `policies` - The list of policies.

### LoadBalancerPolicy Reference

The following attributes are exported:

* `name` - The name of a load balancing policy.  Example: 'LEAST_CONNECTIONS' 

