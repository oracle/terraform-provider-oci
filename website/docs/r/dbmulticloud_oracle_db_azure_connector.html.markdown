---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_connector"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_azure_connector"
description: |-
  Provides the Oracle Db Azure Connector resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_azure_connector
This resource provides the Oracle Db Azure Connector resource in Oracle Cloud Infrastructure Dbmulticloud service.

Creates Oracle DB Azure Connector Resource and configured Azure Identity in Oracle Cloud Infrastructure Database Resource.

  Patch Azure Arc Agent on VM Cluster with new version.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_azure_connector" "test_oracle_db_azure_connector" {
	#Required
	azure_identity_mechanism = var.oracle_db_azure_connector_azure_identity_mechanism
	azure_resource_group = var.oracle_db_azure_connector_azure_resource_group
	azure_subscription_id = oci_onesubscription_subscription.test_subscription.id
	azure_tenant_id = oci_dbmulticloud_azure_tenant.test_azure_tenant.id
	compartment_id = var.compartment_id
	db_cluster_resource_id = oci_cloud_guard_resource.test_resource.id
	display_name = var.oracle_db_azure_connector_display_name
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id

	#Optional
	access_token = var.oracle_db_azure_connector_access_token
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	new_version = var.oracle_db_azure_connector_new_version
	private_endpoint_dns_alias = var.oracle_db_azure_connector_private_endpoint_dns_alias
	private_endpoint_ip_address = var.oracle_db_azure_connector_private_endpoint_ip_address
	system_tags = var.oracle_db_azure_connector_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `access_token` - (Optional) (Updatable) Azure bearer access token. If bearer access token is provided then Service Principal details are not requires.
* `azure_identity_mechanism` - (Required) (Updatable) Azure Identity Mechanism.
* `azure_resource_group` - (Required) (Updatable) Azure Resource Group Name.
* `azure_subscription_id` - (Required) (Updatable) Azure Subscription ID.
* `azure_tenant_id` - (Required) (Updatable) Azure Tenant ID.
* `compartment_id` - (Required) (Updatable) The ID of the compartment that contains Oracle DB Azure Connector Resource.
* `db_cluster_resource_id` - (Required) (Updatable) The ID of the DB Cluster Resource where this Azure Arc Agent Identity to configure.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Oracle DB Azure Connector Resource name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `new_version` - (Optional) (Updatable) 
* `oracle_db_azure_connector_id` - (Required) 
* `private_endpoint_dns_alias` - (Optional) (Updatable) Private endpoint DNS Alias.
* `private_endpoint_ip_address` - (Optional) (Updatable) Private endpoint IP.
* `system_tags` - (Optional) (Updatable) System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_token` - Azure bearer access token. If bearer access token is provided then Service Principal detail is not required.
* `arc_agent_nodes` - List of All VMs where Arc Agent is Install under VMCluster.
	* `current_arc_agent_version` - Current Arc Agent Version installed on this node of VM Cluster.
	* `host_id` - Host ID.
	* `host_name` - Host Name or Azure Arc Agent Name.
	* `status` - The current status of the Azure Arc Agent Resource.
	* `time_last_checked` - time when the Azure Arc Agent's status was checked [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `azure_identity_mechanism` - Azure Identity Mechanism.
* `azure_resource_group` - Azure Resource Group Name.
* `azure_subscription_id` - Azure Subscription ID.
* `azure_tenant_id` - Azure Tenant ID.
* `compartment_id` - The ID of the compartment that contains Oracle DB Azure Connector resource.
* `db_cluster_resource_id` - The ID of the DB Cluster Resource where this Azure Arc Agent identity to configure.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB Azure Connector resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The ID of the Oracle DB Azure Connector resource.
* `last_modification` - Description of the latest modification of the Oracle DB Azure Connector Resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `private_endpoint_dns_alias` - Private endpoint DNS Alias.
* `private_endpoint_ip_address` - Private endpoint IP.
* `state` - The current lifecycle state of the Azure Arc Agent Resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Connector Resource was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Connector Resource was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Azure Connector
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Azure Connector
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Azure Connector


## Import

OracleDbAzureConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector "id"
```

