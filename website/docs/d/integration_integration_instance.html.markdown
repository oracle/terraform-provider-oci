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
	* `alias` - When creating the DNS CNAME record for the custom hostname, this value must be specified in the rdata. 
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. 
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified). 
	* `hostname` - A custom hostname to be used for the integration instance URL, in FQDN format.
* `attachments` - A list of associated attachments to other services 
	* `is_implicit` - 
		* If role == `PARENT`, the attached instance was created by this service instance 
		* If role == `CHILD`, this instance was created from attached instance on behalf of a user 
	* `target_id` - The OCID of the target instance (which could be any other Oracle Cloud Infrastructure PaaS/SaaS resource), to which this instance is attached.
	* `target_instance_url` - The dataplane instance URL of the attached instance
	* `target_role` - The role of the target attachment. 
		* `PARENT` - The target instance is the parent of this attachment. 
		* `CHILD` - The target instance is the child of this attachment.  
	* `target_service_type` - The type of the target instance, such as "FUSION".
* `compartment_id` - Compartment Identifier.
* `consumption_model` - The entitlement used for billing purposes.
* `custom_endpoint` - Details for a custom endpoint for the integration instance.
	* `alias` - When creating the DNS CNAME record for the custom hostname, this value must be specified in the rdata. 
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname. 
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified). 
	* `hostname` - A custom hostname to be used for the integration instance URL, in FQDN format.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Integration Instance Identifier, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `idcs_info` - Information for IDCS access
	* `idcs_app_display_name` - The IDCS application display name associated with the instance
	* `idcs_app_id` - The IDCS application ID associated with the instance
	* `idcs_app_location_url` - URL for the location of the IDCS Application (used by IDCS APIs)
	* `idcs_app_name` - The IDCS application name associated with the instance
	* `instance_primary_audience_url` - The URL used as the primary audience for integration flows in this instance type: string    
* `instance_url` - The Integration Instance URL.
* `integration_instance_type` - Standard or Enterprise type,  Oracle Integration Generation 2 uses ENTERPRISE and STANDARD,  Oracle Integration 3 uses ENTERPRISEX and STANDARDX 
* `is_byol` - Bring your own license.
* `is_file_server_enabled` - The file server is enabled or not.
* `is_visual_builder_enabled` - Visual Builder is enabled or not.
* `message_packs` - The number of configured message packs (if any)
* `network_endpoint_details` - Base representation of a network endpoint. 
	* `allowlisted_http_ips` - Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26") An invalid IP or CIDR block will result in a 400 response. 
	* `allowlisted_http_vcns` - Virtual Cloud Networks allowed to access this network endpoint. 
		* `allowlisted_ips` - Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26") An invalid IP or CIDR block will result in a 400 response. 
		* `id` - The Virtual Cloud Network OCID. 
	* `is_integration_vcn_allowlisted` - The Integration service's VCN is allow-listed to allow integrations to call back into other integrations
	* `network_endpoint_type` - The type of network endpoint.
* `private_endpoint_outbound_connection` - Base representation for Outbound Connection (Reverse Connection). 
	* `nsg_ids` - One or more Network security group Ids. This is an optional argument.
	* `outbound_connection_type` - The type of Outbound Connection. 
	* `subnet_id` - Customer Private Network VCN Subnet OCID. This is a required argument.
* `shape` - Shape
* `state` - The current state of the integration instance.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `time_created` - The time the the Integration Instance was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the IntegrationInstance was updated. An RFC3339 formatted datetime string.

