---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_manifest"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_manifest"
description: |-
  Provides details about a specific Software Source Manifest in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_manifest
This data source provides details about a specific Software Source Manifest resource in Oracle Cloud Infrastructure Os Management Hub service.

Returns an archive containing the list of packages in the software source.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_manifest" "test_software_source_manifest" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


## Attributes Reference

The following attributes are exported:

* `content` - Provides the manifest content used to update the package list of the software source.
* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
