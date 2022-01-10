---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_standard_tag_namespace_templates"
sidebar_current: "docs-oci-datasource-identity-tag_standard_tag_namespace_templates"
description: |-
  Provides the list of Tag Standard Tag Namespace Templates in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tag_standard_tag_namespace_templates
This data source provides the list of Tag Standard Tag Namespace Templates in Oracle Cloud Infrastructure Identity service.

Lists available standard tag namespaces that users can create.


## Example Usage

```hcl
data "oci_identity_tag_standard_tag_namespace_templates" "test_tag_standard_tag_namespace_templates" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `standard_tag_namespace_templates` - The list of standard_tag_namespace_templates.

### TagStandardTagNamespaceTemplate Reference

The following attributes are exported:

* `description` - The default description of the tag namespace that users can use to create the tag namespace
* `standard_tag_namespace_name` - The reserved name of this standard tag namespace
* `status` - The status of the standard tag namespace

