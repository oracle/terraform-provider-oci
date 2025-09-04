---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_connector"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_connector"
description: |-
  Provides details about a specific Oracle Db Azure Connector in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_connector
This data source provides details about a specific Oracle Db Azure Connector resource in Oracle Cloud Infrastructure Dbmulticloud service.

Retrieves the Oracle DB Azure Identity Connector for a specified resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_connector" "test_oracle_db_azure_connector" {
	#Required
	oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_azure_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Connector resource.


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

