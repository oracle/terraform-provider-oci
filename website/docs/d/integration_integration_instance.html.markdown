---
subcategory: "Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_integration_integration_instance"
sidebar_current: "docs-oci-datasource-integration-integration_instance"
description: |-
  Provides details about a specific Integration Instance in Oracle Cloud Infrastructure Integration service
---

# Data Source: oci_integration_integration_instance
This data source provides details about a specific Integration Instance resource in Oracle Cloud Infrastructure Integration service.

Gets a IntegrationInstance by identifier

## Example Usage

```hcl
data "oci_integration_integration_instance" "test_integration_instance" {
	#Required
	integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `integration_instance_id` - (Required) Unique Integration Instance identifier.


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
* `state` - The current state of the integration instance.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `time_created` - The time the the Integration Instance was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the IntegrationInstance was updated. An RFC3339 formatted datetime string.

