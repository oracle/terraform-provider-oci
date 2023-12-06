---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_algorithm"
sidebar_current: "docs-oci-datasource-core-ipsec_algorithm"
description: |-
  Provides details about a specific Ipsec Algorithm in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_algorithm
This data source provides details about a specific Ipsec Algorithm resource in Oracle Cloud Infrastructure Core service.

The parameters allowed for IKE IPSec tunnels.


## Example Usage

```hcl
data "oci_core_ipsec_algorithm" "test_ipsec_algorithm" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `allowed_phase_one_parameters` - Allowed phase one parameters.
	* `authentication_algorithms` - Allowed phase one authentication algorithms.
	* `dh_groups` - Allowed phase one Diffie-Hellman groups.
	* `encryption_algorithms` - Allowed phase one encryption algorithms.
* `allowed_phase_two_parameters` - Allowed phase two parameters.
	* `authentication_algorithms` - Allowed phase two authentication algorithms.
	* `encryption_algorithms` - Allowed phase two encryption algorithms.
	* `pfs_dh_groups` - Allowed perfect forward secrecy Diffie-Hellman groups.
* `default_phase_one_parameters` - Default phase one parameters.
	* `default_authentication_algorithms` - Default phase one authentication algorithms.
	* `default_dh_groups` - Default phase one Diffie-Hellman groups.
	* `default_encryption_algorithms` - Default phase one encryption algorithms.
* `default_phase_two_parameters` - Default phase two parameters.
	* `default_authentication_algorithms` - Default phase two authentication algorithms.
	* `default_encryption_algorithms` - Default phase two encryption algorithms.
	* `default_pfs_dh_group` - Default perfect forward secrecy Diffie-Hellman groups.

