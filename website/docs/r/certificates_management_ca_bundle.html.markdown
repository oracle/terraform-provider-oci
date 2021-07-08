---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_ca_bundle"
sidebar_current: "docs-oci-resource-certificates_management-ca_bundle"
description: |-
  Provides the Ca Bundle resource in Oracle Cloud Infrastructure Certificates Management service
---

# oci_certificates_management_ca_bundle
This resource provides the Ca Bundle resource in Oracle Cloud Infrastructure Certificates Management service.

Creates a new CA bundle according to the details of the request.

## Example Usage

```hcl
resource "oci_certificates_management_ca_bundle" "test_ca_bundle" {
	#Required
	ca_bundle_pem = var.ca_bundle_ca_bundle_pem
	compartment_id = var.compartment_id
	name = var.ca_bundle_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.ca_bundle_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `ca_bundle_pem` - (Required) (Updatable) Certificates (in PEM format) to include in the CA bundle.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment for the CA bundle.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A brief description of the CA bundle.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) A user-friendly name for the CA bundle. Names are unique within a compartment. Avoid entering confidential information. Valid characters include uppercase or lowercase letters, numbers, hyphens, underscores, and periods.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ca Bundle
	* `update` - (Defaults to 20 minutes), when updating the Ca Bundle
	* `delete` - (Defaults to 20 minutes), when destroying the Ca Bundle


## Import

CaBundles can be imported using the `id`, e.g.

```
$ terraform import oci_certificates_management_ca_bundle.test_ca_bundle "id"
```

