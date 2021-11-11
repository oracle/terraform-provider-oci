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

The allowed parameters for IKE IPSec


## Example Usage

```hcl
data "oci_core_ipsec_algorithm" "test_ipsec_algorithm" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `allowed_phase_one_parameters` - Phase One Parameters
	* `authentication_algorithms` - Phase One Authentication Algorithms
	* `dh_groups` - DH Groups
	* `encryption_algorithms` - Phase One Encryption Algorithms
* `allowed_phase_two_parameters` - Phase Two Parameters
	* `authentication_algorithms` - Phase Two Authentication Algorithms
	* `encryption_algorithms` - Phase Two Encryption Algorithms
	* `pfs_dh_groups` - PFS DH Groups
* `default_phase_one_parameters` - Phase One Parameters
	* `default_authentication_algorithms` - Default Phase One Authentication Algorithms
	* `default_dh_groups` - Default DH Groups
	* `default_encryption_algorithms` - Default Phase One Encryption Algorithms
* `default_phase_two_parameters` - Phase Two Parameters
	* `default_authentication_algorithms` - Default Phase Two Authentication Algorithms
	* `default_encryption_algorithms` - Default Phase Two Encryption Algorithms
	* `default_pfs_dh_group` - Default PFS DH Group

