---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_db_management_private_endpoint"
sidebar_current: "docs-oci-datasource-database_management-db_management_private_endpoint"
description: |-
  Provides details about a specific Db Management Private Endpoint in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_db_management_private_endpoint
This data source provides details about a specific Db Management Private Endpoint resource in Oracle Cloud Infrastructure Database Management service.

Gets the details of a specific Database Management private endpoint.

## Example Usage

```hcl
data "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
	#Required
	db_management_private_endpoint_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `db_management_private_endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the Database Management private endpoint.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint.
* `is_cluster` - Specifies whether the Database Management private endpoint can be used for Oracle Databases in a cluster.
* `is_dns_resolution_enabled` - Specifies whether the Database Management private endpoint has DNS proxy server enabled to resolve private host name.
* `name` - The display name of the Database Management private endpoint.
* `nsg_ids` - The OCIDs of the Network Security Groups to which the Database Management private endpoint belongs. 
* `private_ip` - The IP addresses assigned to the Database Management private endpoint. 
* `state` - The current lifecycle state of the Database Management private endpoint.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Database Managament private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.

