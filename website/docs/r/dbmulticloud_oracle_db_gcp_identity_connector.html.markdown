---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_gcp_identity_connector"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_gcp_identity_connector"
description: |-
  Provides the Oracle Db Gcp Identity Connector resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_gcp_identity_connector
This resource provides the Oracle Db Gcp Identity Connector resource in Oracle Cloud Infrastructure Dbmulticloud service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/latest/OracleDbGcpIdentityConnector

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/dbmulticloud

Creates Oracle DB GCP Identity Connector resource.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_gcp_identity_connector" "test_oracle_db_gcp_identity_connector" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.oracle_db_gcp_identity_connector_display_name
	gcp_location = var.oracle_db_gcp_identity_connector_gcp_location
	gcp_resource_service_agent_id = oci_cloud_bridge_agent.test_agent.id
	gcp_workload_identity_pool_id = oci_dataflow_pool.test_pool.id
	gcp_workload_identity_provider_id = oci_identity_identity_provider.test_identity_provider.id
	issuer_url = var.oracle_db_gcp_identity_connector_issuer_url
	project_id = oci_ai_document_project.test_project.id
	resource_id = oci_cloud_guard_resource.test_resource.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Oracle DB GCP Identity Connector resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Oracle DB Google GCP Identity Connector resource name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gcp_location` - (Required) (Updatable) GCP Location.
* `gcp_resource_service_agent_id` - (Required) (Updatable) The ID of the GCP resource service agent.
* `gcp_workload_identity_pool_id` - (Required) (Updatable) The ID of the cloud GCP Workload Identity Pool.
* `gcp_workload_identity_provider_id` - (Required) (Updatable) The ID of the GCP Workload Identity Provider.
* `issuer_url` - (Required) (Updatable) OIDC token issuer Url
* `project_id` - (Required) (Updatable) Project id of the customer project.
* `resource_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the GCP VM Cluster resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Gcp Identity Connector
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Gcp Identity Connector
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Gcp Identity Connector


## Import

OracleDbGcpIdentityConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector "id"
```

