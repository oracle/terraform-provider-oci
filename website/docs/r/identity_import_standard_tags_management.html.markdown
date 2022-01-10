---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_import_standard_tags_management"
sidebar_current: "docs-oci-resource-identity-import_standard_tags_management"
description: |-
  Provides the Import Standard Tags Management resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_import_standard_tags_management
This resource provides the Import Standard Tags Management resource in Oracle Cloud Infrastructure Identity service.

OCI will release Tag Namespaces that our customers can import.
These Tag Namespaces will provide Tags for our customers and Partners to provide consistency and enable data reporting.


## Example Usage

```hcl
resource "oci_identity_import_standard_tags_management" "test_import_standard_tags_management" {
	#Required
	compartment_id = var.compartment_id
	standard_tag_namespace_name = oci_identity_tag_namespace.test_tag_namespace.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment where the bulk create request is submitted and where the tag namespaces will be created. 
* `standard_tag_namespace_name` - (Required) The name of standard tag namespace that will be imported in bulk 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Import Standard Tags Management
	* `update` - (Defaults to 20 minutes), when updating the Import Standard Tags Management
	* `delete` - (Defaults to 20 minutes), when destroying the Import Standard Tags Management


## Import

ImportStandardTagsManagement can be imported using the `id`, e.g.

```
$ terraform import oci_identity_import_standard_tags_management.test_import_standard_tags_management "id"
```

