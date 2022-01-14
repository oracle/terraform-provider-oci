---
subcategory: "Visual Builder"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_visual_builder_vb_instance"
sidebar_current: "docs-oci-resource-visual_builder-vb_instance"
description: |-
  Provides the Vb Instance resource in Oracle Cloud Infrastructure Visual Builder service
---

# oci_visual_builder_vb_instance
This resource provides the Vb Instance resource in Oracle Cloud Infrastructure Visual Builder service.

Creates a new Vb Instance.


## Example Usage

```hcl
resource "oci_visual_builder_vb_instance" "test_vb_instance" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.vb_instance_display_name
	node_count = var.vb_instance_node_count

	#Optional
	alternate_custom_endpoints {
		#Required
		hostname = var.vb_instance_alternate_custom_endpoints_hostname

		#Optional
		certificate_secret_id = oci_vault_secret.test_secret.id
	}
	consumption_model = var.vb_instance_consumption_model
	custom_endpoint {
		#Required
		hostname = var.vb_instance_custom_endpoint_hostname

		#Optional
		certificate_secret_id = oci_vault_secret.test_secret.id
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	idcs_open_id = oci_visual_builder_idcs_open.test_idcs_open.id
	is_visual_builder_enabled = var.vb_instance_is_visual_builder_enabled
}
```

## Argument Reference

The following arguments are supported:

* `alternate_custom_endpoints` - (Optional) (Updatable) A list of alternate custom endpoints to be used for the vb instance URL (contact Oracle for alternateCustomEndpoints availability for a specific instance). 
	* `certificate_secret_id` - (Optional) (Updatable) Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. All certificates should be stored in a single base64 encoded secret Note the update will fail if this is not a valid certificate. 
	* `hostname` - (Required) (Updatable) A custom hostname to be used for the vb instance URL, in FQDN format.
* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `consumption_model` - (Optional) Optional parameter specifying which entitlement to use for billing purposes. Only required if the account possesses more than one entitlement.
* `custom_endpoint` - (Optional) (Updatable) Details for a custom endpoint for the vb instance (update).
	* `certificate_secret_id` - (Optional) (Updatable) Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. All certificates should be stored in a single base64 encoded secret Note the update will fail if this is not a valid certificate. 
	* `hostname` - (Required) (Updatable) A custom hostname to be used for the vb instance URL, in FQDN format.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Vb Instance Identifier.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `idcs_open_id` - (Optional) (Updatable) Encrypted IDCS Open ID token. This is required for pre-UCPIS cloud accounts, but not UCPIS, hence not a required parameter
* `is_visual_builder_enabled` - (Optional) (Updatable) Visual Builder is enabled or not.
* `node_count` - (Required) (Updatable) The number of Nodes


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vb Instance
	* `update` - (Defaults to 20 minutes), when updating the Vb Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Vb Instance


## Import

VbInstances can be imported using the `id`, e.g.

```
$ terraform import oci_visual_builder_vb_instance.test_vb_instance "id"
```

