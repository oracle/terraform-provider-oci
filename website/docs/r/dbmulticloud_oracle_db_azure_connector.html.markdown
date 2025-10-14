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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/latest/OracleDbAzureConnector

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/dbmulticloud

Creates Oracle DB Azure Connector resource and configured Azure Identity in Oracle Database resource.

  Patch Azure Arc Agent on Oracle Cloud VM Cluster with new version.


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

* `access_token` - (Optional) (Updatable) Azure bearer access token.
* `azure_identity_mechanism` - (Required) (Updatable) Azure Identity mechanism.
* `azure_resource_group` - (Required) (Updatable) Azure Resource group name.
* `azure_subscription_id` - (Required) (Updatable) Azure Subscription ID.
* `azure_tenant_id` - (Required) (Updatable) Azure Tenant ID.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Connector resource.
* `db_cluster_resource_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Cloud VM Cluster resource where this Azure Arc Agent Identity to configure.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Oracle DB Azure Connector resource name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `new_version` - (Optional) (Updatable) 
* `oracle_db_azure_connector_id` - (Required) 
* `private_endpoint_dns_alias` - (Optional) (Updatable) Private endpoint's DNS alias.
* `private_endpoint_ip_address` - (Optional) (Updatable) Private endpoint IP.
* `system_tags` - (Optional) (Updatable) System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_token` - Azure bearer access token.
* `arc_agent_nodes` - List of all VMs where Arc Agent is installed under Cloud VM Cluster.
	* `current_arc_agent_version` - Current Arc Agent Version installed on this node of Oracle Cloud VM Cluster.
	* `host_id` - Host ID.
	* `host_name` - Host name or Azure Arc Agent name.
	* `status` - The current status of the Azure Arc Agent resource.
	* `time_last_checked` - Time when the Azure Arc Agent's status was checked [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `azure_identity_connectivity_status` - The current Connectivity status of Azure Identity Connector resource.
* `azure_identity_mechanism` - Azure Identity mechanism.
* `azure_resource_group` - Azure Resource group name.
* `azure_subscription_id` - Azure Subscription ID.
* `azure_tenant_id` - Azure Tenant ID.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Connector resource.
* `db_cluster_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Cloud VM Cluster resource where this Azure Arc Agent identity to configure.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB Azure Connector resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Connector resource.
* `last_modification` - Description of the latest modification of the Oracle DB Azure Connector resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `private_endpoint_dns_alias` - Private endpoint's DNS alias.
* `private_endpoint_ip_address` - Private endpoint IP.
* `state` - The current lifecycle state of the Azure Arc Agent resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Connector resource was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Connector resource was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

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

