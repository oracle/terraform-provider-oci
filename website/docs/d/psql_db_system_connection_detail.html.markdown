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

Gets the DbSystem connection details.

## Example Usage

```hcl
data "oci_psql_db_system_connection_detail" "test_db_system_connection_detail" {
	#Required
	db_system_id = oci_psql_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) unique DbSystem identifier


## Attributes Reference

The following attributes are exported:

* `ca_certificate` - The CA certificate to be used by the Posgresql client to connect to the database. The CA certificate is used to authenticate the server identity.  It is issued by PostgreSQL Service Private CA. 
* `instance_endpoints` - The list of DbInstance endpoints in the DbSystem.
	* `db_instance_id` - Unique identifier of the DbInstance.
	* `endpoint` - The node endpoint information.
		* `fqdn` - The FQDN of the endpoint
		* `ip_address` - The IP Address of the endpoint
		* `port` - The port Address of the endpoint
* `primary_db_endpoint` - The node endpoint information.
	* `fqdn` - The FQDN of the endpoint
	* `ip_address` - The IP Address of the endpoint
	* `port` - The port Address of the endpoint

