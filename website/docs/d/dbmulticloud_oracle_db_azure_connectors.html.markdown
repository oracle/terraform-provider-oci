---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_connectors"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_connectors"
description: |-
  Provides the list of Oracle Db Azure Connectors in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_connectors
This data source provides the list of Oracle Db Azure Connectors in Oracle Cloud Infrastructure Dbmulticloud service.

Lists the all Oracle DB Azure Connector Resource based on filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_connectors" "test_oracle_db_azure_connectors" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	db_cluster_resource_id = oci_cloud_guard_resource.test_resource.id
	display_name = var.oracle_db_azure_connector_display_name
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
	state = var.oracle_db_azure_connector_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_cluster_resource_id` - (Optional) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Resource.
* `display_name` - (Optional) A filter to return Oracle DB Azure Connector Resource that match the given display name.
* `oracle_db_azure_connector_id` - (Optional) A filter to return Oracle DB Azure Blob Mount Resources.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_azure_connector_summary_collection` - The list of oracle_db_azure_connector_summary_collection.

### OracleDbAzureConnector Reference

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

