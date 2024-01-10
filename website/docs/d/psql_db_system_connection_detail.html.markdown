---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_db_system_connection_detail"
sidebar_current: "docs-oci-datasource-psql-db_system_connection_detail"
description: |-
  Provides details about a specific Db System Connection Detail in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_db_system_connection_detail
This data source provides details about a specific Db System Connection Detail resource in Oracle Cloud Infrastructure Psql service.

Gets the database system connection details.

## Example Usage

```hcl
data "oci_psql_db_system_connection_detail" "test_db_system_connection_detail" {
	#Required
	db_system_id = oci_psql_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) A unique identifier for the database system.


## Attributes Reference

The following attributes are exported:

* `ca_certificate` - The CA certificate to be used by the PosgreSQL client to connect to the database. The CA certificate is used to authenticate the server identity.  It is issued by PostgreSQL Service Private CA. 
* `instance_endpoints` - The list of database instance node endpoints in the database system.
	* `db_instance_id` - Unique identifier of the database instance node.
	* `endpoint` - Information about the database instance node endpoint.
		* `fqdn` - The FQDN of the endpoint.
		* `ip_address` - The IP address of the endpoint.
		* `port` - The port address of the endpoint.
* `primary_db_endpoint` - Information about the database instance node endpoint.
	* `fqdn` - The FQDN of the endpoint.
	* `ip_address` - The IP address of the endpoint.
	* `port` - The port address of the endpoint.

