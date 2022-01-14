---
subcategory: "Visual Builder"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_visual_builder_vb_instance"
sidebar_current: "docs-oci-datasource-visual_builder-vb_instance"
description: |-
  Provides details about a specific Vb Instance in Oracle Cloud Infrastructure Visual Builder service
---

# Data Source: oci_visual_builder_vb_instance
This data source provides details about a specific Vb Instance resource in Oracle Cloud Infrastructure Visual Builder service.

Gets a VbInstance by identifier

## Example Usage

```hcl
data "oci_visual_builder_vb_instance" "test_vb_instance" {
	#Required
	vb_instance_id = oci_visual_builder_vb_instance.test_vb_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `vb_instance_id` - (Required) Unique Vb Instance identifier.


## Attributes Reference

The following attributes are exported:

* `alternate_custom_endpoints` - A list of alternate custom endpoints used for the vb instance URL. 
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. 
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified). 
	* `hostname` - A custom hostname to be used for the vb instance URL, in FQDN format.
* `compartment_id` - Compartment Identifier.
* `consumption_model` - The entitlement used for billing purposes.
* `custom_endpoint` - Details for a custom endpoint for the vb instance.
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. 
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified). 
	* `hostname` - A custom hostname to be used for the vb instance URL, in FQDN format.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Vb Instance Identifier, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `instance_url` - The Vb Instance URL.
* `is_visual_builder_enabled` - Visual Builder is enabled or not.
* `node_count` - The number of Nodes
* `state` - The current state of the vb instance.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the VbInstance was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the VbInstance was updated. An RFC3339 formatted datetime string.

