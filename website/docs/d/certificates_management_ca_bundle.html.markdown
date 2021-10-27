---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_ca_bundle"
sidebar_current: "docs-oci-datasource-certificates_management-ca_bundle"
description: |-
  Provides details about a specific Ca Bundle in Oracle Cloud Infrastructure Certificates Management service
---

# Data Source: oci_certificates_management_ca_bundle
This data source provides details about a specific Ca Bundle resource in Oracle Cloud Infrastructure Certificates Management service.

Gets details about the specified CA bundle.

## Example Usage

```hcl
data "oci_certificates_management_ca_bundle" "test_ca_bundle" {
	#Required
	ca_bundle_id = oci_certificates_management_ca_bundle.test_ca_bundle.id
}
```

## Argument Reference

The following arguments are supported:

* `ca_bundle_id` - (Required) The OCID of the CA bundle.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment for the CA bundle.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A brief description of the CA bundle.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the CA bundle.
* `lifecycle_details` - Additional information about the current lifecycle state of the CA bundle.
* `name` - A user-friendly name for the CA bundle. Names are unique within a compartment. Avoid entering confidential information. Valid characters include uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
* `state` - The current lifecycle state of the CA bundle.
* `time_created` - A property indicating when the CA bundle was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 

