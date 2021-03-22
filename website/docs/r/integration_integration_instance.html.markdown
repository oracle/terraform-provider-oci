---
subcategory: "Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_integration_integration_instance"
sidebar_current: "docs-oci-resource-integration-integration_instance"
description: |-
  Provides the Integration Instance resource in Oracle Cloud Infrastructure Integration service
---

# oci_integration_integration_instance
This resource provides the Integration Instance resource in Oracle Cloud Infrastructure Integration service.

Creates a new Integration Instance.


## Example Usage

```hcl
resource "oci_integration_integration_instance" "test_integration_instance" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.integration_instance_display_name
	integration_instance_type = var.integration_instance_integration_instance_type
	is_byol = var.integration_instance_is_byol
	message_packs = var.integration_instance_message_packs

	#Optional
	alternate_custom_endpoints {
		#Required
		hostname = var.integration_instance_alternate_custom_endpoints_hostname

		#Optional
		certificate_secret_id = oci_vault_secret.test_secret.id
	}
	consumption_model = var.integration_instance_consumption_model
	custom_endpoint {
		#Required
		hostname = var.integration_instance_custom_endpoint_hostname

		#Optional
		certificate_secret_id = oci_vault_secret.test_secret.id
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	idcs_at = var.integration_instance_idcs_at
	is_file_server_enabled = var.integration_instance_is_file_server_enabled
	is_visual_builder_enabled = var.integration_instance_is_visual_builder_enabled
	network_endpoint_details {
		#Required
		network_endpoint_type = var.integration_instance_network_endpoint_details_network_endpoint_type

		#Optional
		allowlisted_http_ips = var.integration_instance_network_endpoint_details_allowlisted_http_ips
		allowlisted_http_vcns {
			#Required
			id = var.integration_instance_network_endpoint_details_allowlisted_http_vcns_id

			#Optional
			allowlisted_ips = var.integration_instance_network_endpoint_details_allowlisted_http_vcns_allowlisted_ips
		}
		is_integration_vcn_allowlisted = var.integration_instance_network_endpoint_details_is_integration_vcn_allowlisted
	}
	state = var.integration_instance_target_state
}
```

## Argument Reference

The following arguments are supported:

* `alternate_custom_endpoints` - (Optional) (Updatable) A list of alternate custom endpoints to be used for the integration instance URL (contact Oracle for alternateCustomEndpoints availability for a specific instance). 
	* `certificate_secret_id` - (Optional) (Updatable) Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. All certificates should be stored in a single base64 encoded secret Note the update will fail if this is not a valid certificate. 
	* `hostname` - (Required) (Updatable) A custom hostname to be used for the integration instance URL, in FQDN format.
* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `consumption_model` - (Optional) Optional parameter specifying which entitlement to use for billing purposes. Only required if the account possesses more than one entitlement.
* `custom_endpoint` - (Optional) (Updatable) Details for a custom endpoint for the integration instance (update).
	* `certificate_secret_id` - (Optional) (Updatable) Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. All certificates should be stored in a single base64 encoded secret Note the update will fail if this is not a valid certificate. 
	* `hostname` - (Required) (Updatable) A custom hostname to be used for the integration instance URL, in FQDN format.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Integration Instance Identifier.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `idcs_at` - (Optional) (Updatable) IDCS Authentication token. This is required for all realms with IDCS. Its optional as its not required for non IDCS realms.
* `integration_instance_type` - (Required) (Updatable) Standard or Enterprise type
* `is_byol` - (Required) (Updatable) Bring your own license.
* `is_file_server_enabled` - (Optional) (Updatable) The file server is enabled or not.
* `is_visual_builder_enabled` - (Optional) (Updatable) Visual Builder is enabled or not.
* `message_packs` - (Required) (Updatable) The number of configured message packs
* `network_endpoint_details` - (Optional) Base representation of a network endpoint. 
	* `allowlisted_http_ips` - (Optional) Source IP addresses or IP address ranges ingress rules. 
	* `allowlisted_http_vcns` - (Optional) Virtual Cloud Networks allowed to access this network endpoint. 
		* `allowlisted_ips` - (Optional) Source IP addresses or IP address ranges ingress rules. 
		* `id` - (Required) The Virtual Cloud Network OCID. 
	* `is_integration_vcn_allowlisted` - (Optional) The Integration service's VCN is allow-listed to allow integrations to call back into other integrations
	* `network_endpoint_type` - (Required) The type of network endpoint. 
* `state` - (Optional) (Updatable) The target state for the instance. Could be set to ACTIVE or INACTIVE


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alternate_custom_endpoints` - A list of alternate custom endpoints used for the integration instance URL. 
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. 
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified). 
	* `hostname` - A custom hostname to be used for the integration instance URL, in FQDN format.
* `compartment_id` - Compartment Identifier.
* `consumption_model` - The entitlement used for billing purposes.
* `custom_endpoint` - Details for a custom endpoint for the integration instance.
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. 
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified). 
	* `hostname` - A custom hostname to be used for the integration instance URL, in FQDN format.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Integration Instance Identifier, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `instance_url` - The Integration Instance URL.
* `integration_instance_type` - Standard or Enterprise type
* `is_byol` - Bring your own license.
* `is_file_server_enabled` - The file server is enabled or not.
* `is_visual_builder_enabled` - Visual Builder is enabled or not.
* `message_packs` - The number of configured message packs (if any)
* `network_endpoint_details` - Base representation of a network endpoint. 
	* `allowlisted_http_ips` - Source IP addresses or IP address ranges ingress rules. 
	* `allowlisted_http_vcns` - Virtual Cloud Networks allowed to access this network endpoint. 
		* `allowlisted_ips` - Source IP addresses or IP address ranges ingress rules. 
		* `id` - The Virtual Cloud Network OCID. 
	* `is_integration_vcn_allowlisted` - The Integration service's VCN is allow-listed to allow integrations to call back into other integrations
	* `network_endpoint_type` - The type of network endpoint. 
* `state` - The current state of the integration instance.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `time_created` - The time the the Integration Instance was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the IntegrationInstance was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Integration Instance
	* `update` - (Defaults to 1 hours), when updating the Integration Instance
	* `delete` - (Defaults to 1 hours), when destroying the Integration Instance


## Import

IntegrationInstances can be imported using the `id`, e.g.

```
$ terraform import oci_integration_integration_instance.test_integration_instance "id"
```

