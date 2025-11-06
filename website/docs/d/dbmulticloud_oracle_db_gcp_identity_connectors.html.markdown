---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_gcp_identity_connectors"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_gcp_identity_connectors"
description: |-
  Provides the list of Oracle Db Gcp Identity Connectors in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_gcp_identity_connectors
This data source provides the list of Oracle Db Gcp Identity Connectors in Oracle Cloud Infrastructure Dbmulticloud service.

Lists all Oracle DB GCP Identity Connectors based on the specified filters.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_gcp_identity_connectors" "test_oracle_db_gcp_identity_connectors" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oracle_db_gcp_identity_connector_display_name
	resource_id = oci_cloud_guard_resource.test_resource.id
	state = var.oracle_db_gcp_identity_connector_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [ID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Oracle DB GCP Identity Connector resources that match the specified display name.
* `resource_id` - (Optional) A filter to return Oracle DB Identity Connector resource that match the given resource [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `oracle_db_gcp_identity_connector_summary_collection` - The list of oracle_db_gcp_identity_connector_summary_collection.

### OracleDbGcpIdentityConnector Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Oracle DB GCP Identity Connector resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB GCP Identity Connector resource name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gcp_identity_connectivity_status` - The current Connectivity status of GCP Identity Connector resource.
* `gcp_location` - GCP Location.
* `gcp_nodes` - List of All VMs where GCP Identity Connector is configured for this VMCluster.
	* `host_id` - Host ID.
	* `host_name` - Host Name or Identity Connector name.
	* `status` - The current status of the GCP Identity Connector resource.
	* `time_last_checked` - time when the GCP Identity Connector's status was checked [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `gcp_resource_service_agent_id` - The ID of the GCP resource service agent.
* `gcp_workload_identity_pool_id` - The ID of the cloud GCP Workload Identity Pool.
* `gcp_workload_identity_provider_id` - The ID of the GCP Workload Identity Provider.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB GCP Identity Connector resource.
* `issuer_url` - OIDC token issuer Url.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `project_id` - Project id of the customer project.
* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the GCP VM Cluster resource.
* `state` - The current lifecycle state of the GCP Identity Connector resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB GCP Identity Connector resource was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB GCP Identity Connector resource was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

