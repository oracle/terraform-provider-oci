---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database_dataguard_association_operation"
sidebar_current: "docs-oci-resource-database-autonomous_container_database_dataguard_association_operation"
description: |-
  Provides the Autonomous Container Database Dataguard Association Operation resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_container_database_dataguard_association_operation
This resource provides the Autonomous Container Database Dataguard Association Operation resource in Oracle Cloud Infrastructure Database service.

Perform a new Autonomous Container Database Dataguard Association Operation on an Autonomous Container Database that has Dataguard enabled


## Example Usage

```hcl
resource "oci_database_autonomous_container_database_dataguard_association_operation" "switchover" {
  operation = "switchover" # "failover" or "reinstate"
  autonomous_container_database_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["autonomous_container_database_id"]
  autonomous_container_database_dataguard_association_id = data.oci_database_autonomous_container_database_dataguard_associations.dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]
}
```


## Argument Reference

The following arguments are supported:

* `autonomous_container_database_dataguard_association_id` - (Required) The Autonomous Container Database Dataguard Association [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). This attribute is a forcenew attribute.
* `autonomous_container_database_id` - (Required) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). This attribute is a forcenew attribute.
* `operation` - (Required) There are three type of supported operations `switchover`, `failover`, `reinstate`. `switchover` can only be used for primary database while `failover` and `reinstate` can only be used for standby database. This attribute is a forcenew attribute.


## Import

AutonomousContainerDatabaseDataguardAssociationOperation does not support import.
