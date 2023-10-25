---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_knowledge_base"
sidebar_current: "docs-oci-resource-adm-knowledge_base"
description: |-
  Provides the Knowledge Base resource in Oracle Cloud Infrastructure Adm service
---

# oci_adm_knowledge_base
This resource provides the Knowledge Base resource in Oracle Cloud Infrastructure Adm service.

Creates a new Knowledge Base.

## Example Usage

```hcl
resource "oci_adm_knowledge_base" "test_knowledge_base" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.knowledge_base_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) The name of the knowledge base.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The name of the knowledge base.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
* `state` - The current lifecycle state of the knowledge base.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The creation date and time of the knowledge base (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_updated` - The date and time the knowledge base was last updated (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Knowledge Base
	* `update` - (Defaults to 20 minutes), when updating the Knowledge Base
	* `delete` - (Defaults to 20 minutes), when destroying the Knowledge Base


## Import

KnowledgeBases can be imported using the `id`, e.g.

```
$ terraform import oci_adm_knowledge_base.test_knowledge_base "id"
```

