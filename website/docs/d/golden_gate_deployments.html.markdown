---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployments"
sidebar_current: "docs-oci-datasource-golden_gate-deployments"
description: |-
  Provides the list of Deployments in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployments
This data source provides the list of Deployments in Oracle Cloud Infrastructure Golden Gate service.

Lists the Deployments in a compartment.


## Example Usage

```hcl
data "oci_golden_gate_deployments" "test_deployments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	assignable_connection_id = oci_golden_gate_connection.test_connection.id
	assigned_connection_id = oci_golden_gate_connection.test_connection.id
	display_name = var.deployment_display_name
	fqdn = var.deployment_fqdn
	lifecycle_sub_state = var.deployment_lifecycle_sub_state
	state = var.deployment_state
	supported_connection_type = var.deployment_supported_connection_type
}
```

## Argument Reference

The following arguments are supported:

* `assignable_connection_id` - (Optional) Return the deployments to which the specified connectionId may be assigned. 
* `assigned_connection_id` - (Optional) The OCID of the connection which for the deployment must be assigned. 
* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `fqdn` - (Optional) A filter to return only the resources that match the 'fqdn' given. 
* `lifecycle_sub_state` - (Optional) A filter to return only the resources that match the 'lifecycleSubState' given. 
* `state` - (Optional) A filter to return only the resources that match the 'lifecycleState' given. 
* `supported_connection_type` - (Optional) The connection type which the deployment must support. 


## Attributes Reference

The following attributes are exported:

* `deployment_collection` - The list of deployment_collection.

### Deployment Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `cpu_core_count` - The Minimum number of OCPUs to be made available for this Deployment. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_backup_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced.
* `deployment_diagnostic_data` - Information regarding the deployment diagnostic collection 
	* `bucket` - Name of the bucket where the object is to be uploaded in the object storage
	* `diagnostic_state` - The state of the deployment diagnostic collection. 
	* `namespace` - Name of namespace that serves as a container for all of your buckets
	* `object` - Name of the diagnostic collected and uploaded to object storage
	* `time_diagnostic_end` - The time until which the diagnostic collection should collect the logs. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	* `time_diagnostic_start` - The time from which the diagnostic collection should collect the logs. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
* `deployment_type` - The type of deployment, which can be any one of the Allowed values.  NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.  Its use is discouraged in favor of 'DATABASE_ORACLE'.
* `deployment_url` - The URL of a resource. 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `fqdn` - A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `ingress_ips` - List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.  Customers may optionally set up ingress security rules to restrict traffic from these IP addresses. 
	* `ingress_ip` - A Private Endpoint IPv4 or IPv6 Address created in the customer's subnet. 
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Deployment's CPU core count. 
* `is_healthy` - True if all of the aggregate resources are working correctly. 
* `is_latest_version` - Indicates if the resource is the the latest available version. 
* `is_public` - True if this object is publicly available. 
* `is_storage_utilization_limit_exceeded` - Indicator will be true if the amount of storage being utilized exceeds the allowable storage utilization limit.  Exceeding the limit may be an indication of a misconfiguration of the deployment's GoldenGate service. 
* `license_model` - The Oracle license model that applies to a Deployment. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `lifecycle_sub_state` - Possible GGS lifecycle sub-states. 
* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the loadbalancer in the customer's subnet. The loadbalancer of the public deployment created in the customer subnet. 
* `load_balancer_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a public subnet in the customer tenancy. Can be provided only for public deployments. If provided, the loadbalancer will be created in this subnet instead of the service tenancy. For backward compatiblity this is an optional property for now, but it will become mandatory (for public deployments only) after October 1, 2024. 
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `maintenance_configuration` - Attributes for configuring automatic deployment maintenance. 
	* `bundle_release_upgrade_period_in_days` - Defines auto upgrade period for bundle releases. Manually configured period cannot be longer than service defined period for bundle releases. This period must be shorter or equal to major release upgrade period. Not passing this field during create will equate to using the service default. 
	* `interim_release_upgrade_period_in_days` - Defines auto upgrade period for interim releases. This period must be shorter or equal to bundle release upgrade period. 
	* `is_interim_release_auto_upgrade_enabled` - By default auto upgrade for interim releases are not enabled. If auto-upgrade is enabled for interim release,  you have to specify interimReleaseUpgradePeriodInDays too. 
	* `major_release_upgrade_period_in_days` - Defines auto upgrade period for major releases. Manually configured period cannot be longer than service defined period for major releases. Not passing this field during create will equate to using the service default. 
	* `security_patch_upgrade_period_in_days` - Defines auto upgrade period for releases with security fix. Manually configured period cannot be longer than service defined period for security releases. Not passing this field during create will equate to using the service default. 
* `maintenance_window` - Defines the maintenance window, when automatic actions can be performed. 
	* `day` - Days of the week. 
	* `start_hour` - Start hour for maintenance period. Hour is in UTC. 
* `next_maintenance_action_type` - Type of the next maintenance. 
* `next_maintenance_description` - Description of the next maintenance. 
* `nsg_ids` - An array of Network Security Group OCIDs used to define network access for either Deployments or Connections. 
* `ogg_data` - Deployment Data for an OggDeployment 
	* `admin_username` - The GoldenGate deployment console username. 
	* `certificate` - The base64 encoded content of the PEM file containing the SSL certificate. 
	* `credential_store` - The type of credential store for OGG. 
	* `deployment_name` - The name given to the GoldenGate service deployment. The name must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter. 
	* `identity_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Identity Domain when IAM credential store is used. 
	* `ogg_version` - Version of OGG 
	* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the deployment password is stored. 
* `private_ip_address` - The private IP address in the customer's VCN representing the access point for the associated endpoint service in the GoldenGate service VCN. 
* `public_ip_address` - The public IP address representing the access point for the Deployment. 
* `state` - Possible lifecycle states. 
* `storage_utilization_in_bytes` - The amount of storage being utilized (in bytes) 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet of the deployment's private endpoint. The subnet must be a private subnet. For backward compatibility, public subnets are allowed until May 31 2025, after which the private subnet will be enforced. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_of_next_maintenance` - The time of next maintenance schedule. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_ogg_version_supported_until` - The time until OGG version is supported. After this date has passed OGG version will not be available anymore. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_upgrade_required` - Note: Deprecated: Use timeOfNextMaintenance instead, or related upgrade records  to check, when deployment will be forced to upgrade to a newer version. Old description: The date the existing version in use will no longer be considered as usable and an upgrade will be required.  This date is typically 6 months after the version was released for use by GGS.  The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

