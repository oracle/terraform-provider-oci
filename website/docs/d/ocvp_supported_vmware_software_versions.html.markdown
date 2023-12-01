---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_supported_vmware_software_versions"
sidebar_current: "docs-oci-datasource-ocvp-supported_vmware_software_versions"
description: |-
  Provides the list of Supported Vmware Software Versions in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_supported_vmware_software_versions
This data source provides the list of Supported Vmware Software Versions in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Lists the versions of bundled VMware software supported by the Oracle Cloud
VMware Solution.


## Example Usage

```hcl
data "oci_ocvp_supported_vmware_software_versions" "test_supported_vmware_software_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	host_shape_name = oci_core_shape.test_shape.name
	version = var.supported_vmware_software_version_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `host_shape_name` - (Optional) A filter to return only resources that match or support the given ESXi host shape.
* `version` - (Optional) A filter to return only resources that match the given VMware software version exactly.


## Attributes Reference

The following attributes are exported:

* `items` - The list of items.

### SupportedVmwareSoftwareVersion Reference

The following attributes are exported:

* `items` - A list of the supported versions of bundled VMware software.
	* `description` - A description of the software in the bundle.
	* `esxi_software_versions` - A list of supported ESXi software versions. 
		* `description` - A description of the software in the bundle.
		* `supported_host_shape_names` - A list of ESXi host shapes supported by the version of bundled software.
		* `version` - A short, unique string that identifies the version of bundled software. 
	* `version` - A short, unique string that identifies the version of bundled software. 

