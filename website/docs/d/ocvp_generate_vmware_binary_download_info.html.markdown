---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_generate_vmware_binary_download_info"
sidebar_current: "docs-oci-datasource-ocvp-generate_vmware_binary_download_info"
description: |-
Generates VMware binary download information for an SDDC in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_generate_vmware_binary_download_info
This data source generates VMware binary download information for an SDDC in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

## Example Usage

```hcl
data "oci_ocvp_generate_vmware_binary_download_info" "test_generate_vmware_binary_download_info" {
  sddc_id                  = var.sddc_id
  vmware_binary_file_name  = var.vmware_binary_file_name
}
```

## Argument Reference

The following arguments are supported:

* `sddc_id` - (Required) The [OCID](/Content/General/Concepts/identifiers.htm) of the SDDC.
* `vmware_binary_file_name` - (Required) The VMware binary file name to generate download information for.

## Attributes Reference

The following attributes are exported:

* `file_name` - The VMware binary file name.
* `time_expires` - The expiration date and time for the download URL, in RFC3339 format.
* `url` - The URL to download the VMware binary.
