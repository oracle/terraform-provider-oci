---
subcategory: "Visual Builder"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_visual_builder_vb_instances"
sidebar_current: "docs-oci-datasource-visual_builder-vb_instances"
description: |-
  Provides the list of Vb Instances in Oracle Cloud Infrastructure Visual Builder service
---

# Data Source: oci_visual_builder_vb_instances
This data source provides the list of Vb Instances in Oracle Cloud Infrastructure Visual Builder service.

Returns a list of Vb Instances.


## Example Usage

```hcl
data "oci_visual_builder_vb_instances" "test_vb_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.vb_instance_display_name
	state = var.vb_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) Life cycle state to query on.


## Attributes Reference

The following attributes are exported:

* `vb_instance_summary_collection` - The list of vb_instance_summary_collection.

### VbInstance Reference

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

