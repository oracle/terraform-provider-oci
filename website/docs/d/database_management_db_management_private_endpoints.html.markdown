---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_db_management_private_endpoints"
sidebar_current: "docs-oci-datasource-database_management-db_management_private_endpoints"
description: |-
  Provides the list of Db Management Private Endpoints in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_db_management_private_endpoints
This data source provides the list of Db Management Private Endpoints in Oracle Cloud Infrastructure Database Management service.

Gets a list of Database Management private endpoints.


## Example Usage

```hcl
data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.db_management_private_endpoint_name
	state = var.db_management_private_endpoint_state
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `state` - (Optional) The lifecycle state of a resource.
* `vcn_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `db_management_private_endpoint_collection` - The list of db_management_private_endpoint_collection.

### DbManagementPrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `description` - The description of the private endpoint.
* `id` - The OCID of the Database Management private endpoint.
* `name` - The display name of the private endpoint.
* `nsg_ids` - The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_ip` - The private IP addresses assigned to the private endpoint. 
* `state` - The current state of the private endpoint.
* `subnet_id` - The OCID of the subnet.
* `time_created` - The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN.

