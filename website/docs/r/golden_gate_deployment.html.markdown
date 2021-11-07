---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment"
sidebar_current: "docs-oci-resource-golden_gate-deployment"
description: |-
  Provides the Deployment resource in Oracle Cloud Infrastructure Golden Gate service
---

# oci_golden_gate_deployment
This resource provides the Deployment resource in Oracle Cloud Infrastructure Golden Gate service.

Creates a new Deployment.


## Example Usage

```hcl
resource "oci_golden_gate_deployment" "test_deployment" {
	#Required
	compartment_id = var.compartment_id
	cpu_core_count = var.deployment_cpu_core_count
	deployment_type = var.deployment_deployment_type
	display_name = var.deployment_display_name
	is_auto_scaling_enabled = var.deployment_is_auto_scaling_enabled
	license_model = var.deployment_license_model
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deployment_backup_id = oci_golden_gate_deployment_backup.test_deployment_backup.id
	description = var.deployment_description
	fqdn = var.deployment_fqdn
	freeform_tags = {"bar-key"= "value"}
	is_public = var.deployment_is_public
	nsg_ids = var.deployment_nsg_ids
	ogg_data {
		#Required
		admin_password = var.deployment_ogg_data_admin_password
		admin_username = var.deployment_ogg_data_admin_username
		deployment_name = oci_golden_gate_deployment.test_deployment.name

		#Optional
		certificate = var.deployment_ogg_data_certificate
		key = var.deployment_ogg_data_key
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `cpu_core_count` - (Required) (Updatable) The Minimum number of OCPUs to be made available for this Deployment. 
* `defined_tags` - (Optional) (Updatable) Tags defined for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_backup_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced. 
* `deployment_type` - (Required) The type of deployment, the value determines the exact 'type' of service executed in the Deployment. NOTE: Use of the value OGG is maintained for backward compatibility purposes.  Its use is discouraged  in favor of the equivalent DATABASE_ORACLE value. 
* `description` - (Optional) (Updatable) Metadata about this specific object. 
* `display_name` - (Required) (Updatable) An object's Display Name. 
* `fqdn` - (Optional) (Updatable) A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_auto_scaling_enabled` - (Required) (Updatable) Indicates if auto scaling is enabled for the Deployment's CPU core count. 
* `is_public` - (Optional) (Updatable) True if this object is publicly available. 
* `license_model` - (Required) (Updatable) The Oracle license model that applies to a Deployment. 
* `nsg_ids` - (Optional) (Updatable) An array of [Network Security Group](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm) OCIDs used to define network access for a deployment. 
* `ogg_data` - (Optional) (Updatable) Deployment Data for creating an OggDeployment 
	* `admin_password` - (Required) (Updatable) The password associated with the GoldenGate deployment console username. The password must be 8 to 30 characters long and must contain at least 1 uppercase, 1 lowercase, 1 numeric, and 1 special character. Special characters such as ‘$’, ‘^’, or ‘?’ are not allowed. 
	* `admin_username` - (Required) (Updatable) The GoldenGate deployment console username. 
	* `certificate` - (Optional) (Updatable) A PEM-encoded SSL certificate. 
	* `deployment_name` - (Required) The name given to the GoldenGate service deployment. The name must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter. 
	* `key` - (Optional) (Updatable) A PEM-encoded private key. 
* `subnet_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `cpu_core_count` - The Minimum number of OCPUs to be made available for this Deployment. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_backup_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced. 
* `deployment_type` - The type of deployment, the value determines the exact 'type' of service executed in the Deployment. NOTE: Use of the value OGG is maintained for backward compatibility purposes.  Its use is discouraged  in favor of the equivalent DATABASE_ORACLE value. 
* `deployment_url` - The URL of a resource. 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `fqdn` - A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Deployment's CPU core count. 
* `is_healthy` - True if all of the aggregate resources are working correctly. 
* `is_latest_version` - Indicates if the resource is the the latest available version. 
* `is_public` - True if this object is publicly available. 
* `license_model` - The Oracle license model that applies to a Deployment. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `lifecycle_sub_state` - Possible GGS lifecycle sub-states. 
* `nsg_ids` - An array of [Network Security Group](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm) OCIDs used to define network access for a deployment. 
* `ogg_data` - Deployment Data for an OggDeployment 
	* `admin_username` - The GoldenGate deployment console username. 
	* `certificate` - A PEM-encoded SSL certificate. 
	* `deployment_name` - The name given to the GoldenGate service deployment. The name must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter. 
	* `ogg_version` - Version of OGG 
* `private_ip_address` - The private IP address in the customer's VCN representing the access point for the associated endpoint service in the GoldenGate service VCN. 
* `public_ip_address` - The public IP address representing the access point for the Deployment. 
* `state` - Possible lifecycle states. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_upgrade_required` - The date the existing version in use will no longer be considered as usable and an upgrade will be required.  This date is typically 6 months after the version was released for use by GGS.  The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Deployment
	* `update` - (Defaults to 20 minutes), when updating the Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Deployment


## Import

Deployments can be imported using the `id`, e.g.

```
$ terraform import oci_golden_gate_deployment.test_deployment "id"
```

