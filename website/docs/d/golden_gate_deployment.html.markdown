---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment"
sidebar_current: "docs-oci-datasource-golden_gate-deployment"
description: |-
  Provides details about a specific Deployment in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment
This data source provides details about a specific Deployment resource in Oracle Cloud Infrastructure Golden Gate service.

Retrieves a deployment.


## Example Usage

```hcl
data "oci_golden_gate_deployment" "test_deployment" {
	#Required
	deployment_id = oci_golden_gate_deployment.test_deployment.id
}
```

## Argument Reference

The following arguments are supported:

* `deployment_id` - (Required) A unique Deployment identifier. 


## Attributes Reference

The following attributes are exported:

* `backup_schedule` - Defines the schedule of the deployment backup. 
	* `bucket` - Name of the bucket where the object is to be uploaded in the object storage
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
	* `frequency_backup_scheduled` - The frequency of the deployment backup schedule. Frequency can be DAILY, WEEKLY or MONTHLY. 
	* `is_metadata_only` - Parameter to allow users to create backup without trails
	* `namespace` - Name of namespace that serves as a container for all of your buckets
	* `time_backup_scheduled` - The start timestamp for the deployment backup schedule. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-10-25T18:19:29.600Z`.
* `availability_domain` - The availability domain of a placement.
* `category` - The deployment category defines the broad separation of the deployment type into three categories. Currently the separation is 'DATA_REPLICATION', 'STREAM_ANALYTICS' and 'DATA_TRANSFORMS'. 
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
* `deployment_role` - The type of the deployment role. 
* `deployment_type` - The type of deployment, which can be any one of the Allowed values.  NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.  Its use is discouraged in favor of 'DATABASE_ORACLE'.
* `deployment_type` - The type of deployment, which can be any one of the Allowed values.  NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.  Its use is discouraged in favor of 'DATABASE_ORACLE'.
* `deployment_url` - The URL of a resource. 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `environment_type` - Specifies whether the deployment is used in a production or development/testing environment. 
* `fault_domain` - The fault domain of a placement.
* `fqdn` - A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `ingress_ips` - List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.  Customers may optionally set up ingress security rules to restrict traffic from these IP addresses. 
	* `ingress_ip` - A Private Endpoint IPv4 or IPv6 Address created in the customer's subnet. 
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Deployment's CPU core count. 
* `is_healthy` - True if all of the aggregate resources are working correctly. 
* `is_latest_version` - Indicates if the resource is the the latest available version. 
* `is_public` - True if this object is publicly available. 
* `is_storage_utilization_limit_exceeded` - Deprecated: This field is not updated and will be removed in future versions. If storage utilization exceeds the limit, the respective warning message will appear in deployment messages, which can be accessed through /messages?deploymentId=. Indicator will be true if the amount of storage being utilized exceeds the allowable storage utilization limit.  Exceeding the limit may be an indication of a misconfiguration of the deployment's GoldenGate service. 
* `license_model` - The Oracle license model that applies to a Deployment. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `lifecycle_sub_state` - Possible GGS lifecycle sub-states. 
* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the loadbalancer in the customer's subnet. The loadbalancer of the public deployment created in the customer subnet.
* `load_balancer_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a public subnet in the customer tenancy. Can be provided only for public deployments. If provided, the loadbalancer will be created in this subnet instead of the service tenancy. For backward compatibility, this is an optional property. It will become mandatory for public deployments after October 1, 2024.
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
	* `group_to_roles_mapping` - Defines the IDP Groups to GoldenGate roles mapping. This field is used only for IAM deployment and does not have any impact on non-IAM deployments. For IAM deployment, when user does not specify this mapping, then it has null value and default mapping is used. User belonging to each group can only perform the actions according to the role the respective group is mapped to. 
		* `administrator_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role administratorGroup. It grants full access to the user, including the ability to alter general, non-security related operational parameters and profiles of the server. 
		* `operator_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role operatorGroup. It allows users to perform only operational actions, like starting and stopping resources. Operators cannot alter the operational parameters or profiles of the MA server. 
		* `security_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role securityGroup. It grants administration of security related objects and invoke security related service requests. This role has full privileges. 
		* `user_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role userGroup. It allows information-only service requests, which do not alter or affect the operation of either the MA. Examples of query and read-only information include performance metric information and resource status and monitoring information 
	* `identity_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Identity Domain when IAM credential store is used. 
	* `ogg_version` - Version of OGG 
	* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the deployment password is stored. 
* `placements` - An array of local peers of deployment 
	* `availability_domain` - The availability domain of a placement.
	* `fault_domain` - The fault domain of a placement.
* `private_ip_address` - The private IP address in the customer's VCN representing the access point for the associated endpoint service in the GoldenGate service VCN. 
* `public_ip_address` - The public IP address representing the access point for the Deployment. 
* `source_deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `state` - Possible lifecycle states. 
* `storage_utilization_in_bytes` - The amount of storage being utilized (in bytes) 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet of the deployment's private endpoint. The subnet must be a private subnet. For backward compatibility, public subnets are allowed until May 31 2025, after which the private subnet will be enforced. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_last_backup_scheduled` - The timestamp of last deployment backup scheduled. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-10-25T18:19:29.600Z`. 
* `time_next_backup_scheduled` - The timestamp of next deployment backup scheduled. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-10-26T20:19:29.600Z`. 
* `time_of_next_maintenance` - The time of next maintenance schedule. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_ogg_version_supported_until` - The time until OGG version is supported. After this date has passed OGG version will not be available anymore. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_role_changed` - The time of the last role change. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_upgrade_required` - Note: Deprecated: Use timeOfNextMaintenance instead, or related upgrade records  to check, when deployment will be forced to upgrade to a newer version. Old description: The date the existing version in use will no longer be considered as usable and an upgrade will be required.  This date is typically 6 months after the version was released for use by GGS.  The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

