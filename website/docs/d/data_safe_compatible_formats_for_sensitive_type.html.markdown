---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_compatible_formats_for_sensitive_type"
sidebar_current: "docs-oci-datasource-data_safe-compatible_formats_for_sensitive_type"
description: |-
  Provides details about a specific Compatible Formats For Sensitive Type in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_compatible_formats_for_sensitive_type
This data source provides details about a specific Compatible Formats For Sensitive Type resource in Oracle Cloud Infrastructure Data Safe service.

Gets a list of library masking formats compatible with the existing sensitive types.
For each sensitive type, it returns the assigned default masking format as well as
the other library masking formats that have the sensitiveTypeIds attribute containing
the OCID of the sensitive type.


## Example Usage

```hcl
data "oci_data_safe_compatible_formats_for_sensitive_type" "test_compatible_formats_for_sensitive_type" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.compatible_formats_for_sensitive_type_access_level
	compartment_id_in_subtree = var.compatible_formats_for_sensitive_type_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 


## Attributes Reference

The following attributes are exported:

* `formats_for_sensitive_type` - An array of library masking formats compatible with the existing sensitive types.
	* `masking_formats` - An array of the library masking formats compatible with the sensitive type.
		* `description` - The description of the masking format.
		* `id` - The OCID of the masking format.
		* `name` - The name of the masking format.
	* `sensitive_type_id` - The OCID of the sensitive type.

