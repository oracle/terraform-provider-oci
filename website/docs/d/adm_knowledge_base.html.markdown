---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_knowledge_base"
sidebar_current: "docs-oci-datasource-adm-knowledge_base"
description: |-
  Provides details about a specific Knowledge Base in Oracle Cloud Infrastructure ADM service
---

# Data Source: oci_adm_knowledge_base
This data source provides details about a specific Knowledge Base resource in Oracle Cloud Infrastructure ADM service.

Returns the details of the specified Knowledge Base.

## Example Usage

```hcl
data "oci_adm_knowledge_base" "test_knowledge_base" {
	#Required
	knowledge_base_id = oci_adm_knowledge_base.test_knowledge_base.id
}
```

## Argument Reference

The following arguments are supported:

* `knowledge_base_id` - (Required) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of a Knowledge Base, as a URL path parameter.


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

