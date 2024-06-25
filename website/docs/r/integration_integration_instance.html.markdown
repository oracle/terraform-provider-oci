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
	domain_id = oci_identity_domain.test_domain.id
	freeform_tags = {"bar-key"= "value"}
	idcs_at = var.integration_instance_idcs_at
	is_disaster_recovery_enabled = var.integration_instance_is_disaster_recovery_enabled
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
	shape = var.integration_instance_shape
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
* `domain_id` - (Optional) The OCID of the identity domain, that will be used to determine the corresponding Idcs Stripe and create an Idcs application within the stripe. This parameter is mutually exclusive with parameter: idcsAt, i.e only one of two parameters should be specified.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `idcs_at` - (Optional) (Updatable) IDCS Authentication token. This is required for all realms with IDCS. Its optional as its not required for non IDCS realms.
* `integration_instance_type` - (Required) (Updatable) Standard or Enterprise type, Oracle Integration Generation 2 uses ENTERPRISE and STANDARD, Oracle Integration 3 uses ENTERPRISEX and STANDARDX
* `domain_id` - (Optional) The OCID of the identity domain, that will be used to determine the  corresponding Idcs Stripe and create an Idcs application within the stripe.  This parameter is mutually exclusive with parameter: idcsAt, i.e only one of  two parameters should be specified.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `idcs_at` - (Optional) (Updatable) IDCS Authentication token. This is required for all realms with IDCS. Its optional as its not required for non IDCS realms.
* `integration_instance_type` - (Required) (Updatable) Standard or Enterprise type,  Oracle Integration Generation 2 uses ENTERPRISE and STANDARD,  Oracle Integration 3 uses ENTERPRISEX and STANDARDX
* `is_byol` - (Required) (Updatable) Bring your own license.
* `is_disaster_recovery_enabled` - (Optional) Is Disaster Recovery enabled or not.
* `is_file_server_enabled` - (Optional) (Updatable) The file server is enabled or not.
* `is_visual_builder_enabled` - (Optional) (Updatable) Visual Builder is enabled or not.
* `message_packs` - (Required) (Updatable) The number of configured message packs
* `network_endpoint_details` - (Optional) Base representation of a network endpoint.
	* `allowlisted_http_ips` - (Optional) Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26") An invalid IP or CIDR block will result in a 400 response.
	* `allowlisted_http_vcns` - (Optional) Virtual Cloud Networks allowed to access this network endpoint.
		* `allowlisted_ips` - (Optional) Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26") An invalid IP or CIDR block will result in a 400 response.
		* `id` - (Required) The Virtual Cloud Network OCID.
	* `is_integration_vcn_allowlisted` - (Optional) The Integration service's VCN is allow-listed to allow integrations to call back into other integrations
	* `network_endpoint_type` - (Required) The type of network endpoint.
* `shape` - (Optional) Shape
* `enable_process_automation_trigger` - (Optional) (Updatable) An optional property when incremented triggers Enable Process Automation. Could be set to any integer value.
* `extend_data_retention_trigger` - (Optional) (Updatable) An optional property when incremented triggers Extend Data Retention. Could be set to any integer value.
* `add_oracle_managed_custom_endpoint_trigger` - (Optional) (Updatable) An optional property when incremented triggers Add Oracle Managed Custom Endpoint. Could be set to any integer value.
* `enable_process_automation_trigger` - (Optional) (Updatable) An optional property when incremented triggers Enable Process Automation. Could be set to any integer value.
* `remove_oracle_managed_custom_endpoint_trigger` - (Optional) (Updatable) An optional property when incremented triggers Remove Oracle Managed Custom Endpoint. Could be set to any integer value.
* `failover_trigger` - (Optional) (Updatable) An optional property when incremented triggers Failover. Could be set to any integer value.
  * `network_endpoint_type` - (Required) The type of network endpoint.
* `failover_trigger` - (Optional) (Updatable) An optional property when incremented triggers Failover. Could be set to any integer value.
	* `network_endpoint_type` - (Required) The type of network endpoint.
* `shape` - (Optional) Shape
* `state` - (Optional) (Updatable) The target state for the instance. Could be set to ACTIVE or INACTIVE


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alternate_custom_endpoints` - A list of alternate custom endpoints used for the integration instance URL.
	* `alias` - When creating the DNS CNAME record for the custom hostname, this value must be specified in the rdata.
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname.
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified).
	* `dns_type` - Type of DNS.
	* `dns_zone_name` - DNS Zone name
	* `hostname` - A custom hostname to be used for the integration instance URL, in FQDN format.
* `managed_type` - Indicates if custom endpoint is managed by oracle or customer.
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
	* `dns_type` - Type of DNS.
	* `dns_zone_name` - DNS Zone name
	* `hostname` - A custom hostname to be used for the integration instance URL, in FQDN format.
	* `managed_type` - Indicates if custom endpoint is managed by oracle or customer.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}`
	* `alias` - When creating the DNS CNAME record for the custom hostname, this value must be specified in the rdata.
	* `certificate_secret_id` - Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname.
	* `certificate_secret_version` - The secret version used for the certificate-secret-id (if certificate-secret-id is specified).
	* `hostname` - A custom hostname to be used for the integration instance URL, in FQDN format.
* `data_retention_period` - Data retention period set for given integration instance
* `disaster_recovery_details` - Disaster recovery details for the integration instance created in the region.
    * `cross_region_integration_instance_details` - Details of integration instance created in cross region for disaster recovery.
        * `id` - Cross region integration instance identifier
        * `region` - Cross region where integration instance is created
        * `role` - Role of the integration instance in the region
        * `time_role_changed` - Time when cross region integration instance role was changed
    * `regional_instance_url` - Region specific instance url for the integration instance in the region
    * `role` - Role of the integration instance in the region
* `display_name` - Integration Instance Identifier, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `idcs_info` - Information for IDCS access
	* `idcs_app_display_name` - The IDCS application display name associated with the instance
	* `idcs_app_id` - The IDCS application ID associated with the instance
	* `idcs_app_location_url` - URL for the location of the IDCS Application (used by IDCS APIs)
	* `idcs_app_name` - The IDCS application name associated with the instance
	* `instance_primary_audience_url` - The URL used as the primary audience for integration flows in this instance type: string * `instance_design_time_url` - The Integration Instance Design Time URL
* `instance_url` - The Integration Instance URL.
* `integration_instance_type` - Standard or Enterprise type, Oracle Integration Generation 2 uses ENTERPRISE and STANDARD, Oracle Integration 3 uses ENTERPRISEX and STANDARDX
* `instance_url` - The Integration Instance URL.
	* `instance_primary_audience_url` - The URL used as the primary audience for integration flows in this instance type: string
* `instance_url` - The Integration Instance URL.
* `integration_instance_type` - Standard or Enterprise type, Oracle Integration Generation 2 uses ENTERPRISE and STANDARD, Oracle Integration 3 uses ENTERPRISEX and STANDARDX
* `is_byol` - Bring your own license.
* `is_disaster_recovery_enabled` - Is Disaster Recovery enabled for the integrationInstance
* `is_file_server_enabled` - The file server is enabled or not.
* `is_visual_builder_enabled` - Visual Builder is enabled or not.
* `lifecycle_details` - Additional details of lifecycleState or substates
* `message_packs` - The number of configured message packs (if any)
* `network_endpoint_details` - Base representation of a network endpoint.
	* `allowlisted_http_ips` - Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26") An invalid IP or CIDR block will result in a 400 response.
	* `allowlisted_http_vcns` - Virtual Cloud Networks allowed to access this network endpoint.
		* `allowlisted_ips` - Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26") An invalid IP or CIDR block will result in a 400 response.
		* `id` - The Virtual Cloud Network OCID.
	* `is_integration_vcn_allowlisted` - The Integration service's VCN is allow-listed to allow integrations to call back into other integrations
	* `network_endpoint_type` - The type of network endpoint.
* `shape` - Shape
* `private_endpoint_outbound_connection` - Base representation for Outbound Connection (Reverse Connection).
	* `nsg_ids` - One or more Network security group Ids. This is an optional argument.
	* `outbound_connection_type` - The type of Outbound Connection.
	* `subnet_id` - Customer Private Network VCN Subnet OCID. This is a required argument.
* `shape` - Shape
* `state` - The current state of the integration instance.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the the Integration Instance was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the IntegrationInstance was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Integration Instance
	* `update` - (Defaults to 1 hours), when updating the Integration Instance
	* `delete` - (Defaults to 1 hours), when destroying the Integration Instance


## Import

IntegrationInstances can be imported using the `id`, e.g.

```
$ terraform import oci_integration_integration_instance.test_integration_instance "id"
```
