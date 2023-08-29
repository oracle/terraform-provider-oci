---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_peers"
sidebar_current: "docs-oci-datasource-database-autonomous_database_peers"
description: |-
  Provides the list of Autonomous Database Peers in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_peers
This data source provides the list of Autonomous Database Peers in Oracle Cloud Infrastructure Database service.

Lists the Autonomous Database peers for the specified Autonomous Database.


## Example Usage

```hcl
data "oci_database_autonomous_database_peers" "test_autonomous_database_peers" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_database_peer_collection` - The list of autonomous_database_peer_collection.

### AutonomousDatabasePeer Reference

The following attributes are exported:

* `items` - This array holds details about Autonomous Database Peers for Oracle an Autonomous Database.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	* `region` - The name of the region where this peer Autonomous Database clone exists.

