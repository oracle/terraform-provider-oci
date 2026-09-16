---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_retrieve_vmware_binaries"
sidebar_current: "docs-oci-datasource-ocvp-retrieve_vmware_binaries"
description: |-
Retrieves the available VMware binaries for an SDDC in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_retrieve_vmware_binaries
This data source retrieves the available VMware binaries for an SDDC in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

## Example Usage

```hcl
data "oci_ocvp_retrieve_vmware_binaries" "test_retrieve_vmware_binaries" {
  sddc_id = var.sddc_id
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) One or more name/value filters to apply to the returned items.
* `sddc_id` - (Required) The [OCID](/Content/General/Concepts/identifiers.htm) of the SDDC.

## Attributes Reference

The following attributes are exported:

* `items` - The list of available VMware binaries.
  * `checksum` - Base64-encoded SHA256 hash of the VMware binary object data.
  * `description` - Description of the VMware binary.
  * `file_name` - The VMware binary file name.
  * `size_in_bytes` - Size of the VMware binary file in bytes.
