---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_virtual_circuit_public_prefixes"
sidebar_current: "docs-oci-datasource-core-virtual_circuit_public_prefixes"
description: |-
  Provides the list of Virtual Circuit Public Prefixes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_virtual_circuit_public_prefixes
This data source provides the list of Virtual Circuit Public Prefixes in Oracle Cloud Infrastructure Core service.

Lists the public IP prefixes and their details for the specified
public virtual circuit.


## Example Usage

```hcl
data "oci_core_virtual_circuit_public_prefixes" "test_virtual_circuit_public_prefixes" {
	#Required
	virtual_circuit_id = oci_core_virtual_circuit.test_virtual_circuit.id

	#Optional
	verification_state = var.virtual_circuit_public_prefix_verification_state
}
```

## Argument Reference

The following arguments are supported:

* `verification_state` - (Optional) A filter to only return resources that match the given verification state.

	The state value is case-insensitive. 
* `virtual_circuit_id` - (Required) The OCID of the virtual circuit.


## Attributes Reference

The following attributes are exported:

* `virtual_circuit_public_prefixes` - The list of virtual_circuit_public_prefixes.

### VirtualCircuitPublicPrefix Reference

The following attributes are exported:

* `cidr_block` - Publix IP prefix (CIDR) that the customer specified.
* `verification_state` - Oracle must verify that the customer owns the public IP prefix before traffic for that prefix can flow across the virtual circuit. Verification can take a few business days. `IN_PROGRESS` means Oracle is verifying the prefix. `COMPLETED` means verification succeeded. `FAILED` means verification failed and traffic for this prefix will not flow across the connection. 

