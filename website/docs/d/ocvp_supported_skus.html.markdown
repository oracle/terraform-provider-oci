---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_supported_skus"
sidebar_current: "docs-oci-datasource-ocvp-supported_skus"
description: |-
  Provides the list of Supported Skus in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_supported_skus
This data source provides the list of Supported Skus in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.  
!> **WARNING:** This data source is deprecated and will be removed, please use "oci_ocvp_supported_commitments" instead.

Lists supported SKUs.


## Example Usage

```hcl
data "oci_ocvp_supported_skus" "test_supported_skus" {
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

* `items` - The list of the supported SKUs.

### Items Reference

The following attributes are exported:

* `name` - name of SKU

