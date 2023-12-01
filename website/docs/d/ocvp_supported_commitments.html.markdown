---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_supported_commitments"
sidebar_current: "docs-oci-datasource-ocvp-supported_commitments"
description: |-
  Provides the list of Supported Commitments in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_supported_commitments
This data source provides the list of Supported Commitments in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Lists supported Commitments.


## Example Usage

```hcl
data "oci_ocvp_supported_commitments" "test_supported_commitments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	host_shape_name = oci_core_shape.test_shape.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `host_shape_name` - (Optional) A filter to return only resources that match or support the given ESXi host shape.


## Attributes Reference

The following attributes are exported:

* `supported_commitment_summary_collection` - The list of supported_commitment_summary_collection.

### SupportedCommitment Reference

The following attributes are exported:

* `items` - A list of the supported Commitments.
	* `name` - name of Commitment

