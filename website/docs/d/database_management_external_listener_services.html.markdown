---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_listener_services"
sidebar_current: "docs-oci-datasource-database_management-external_listener_services"
description: |-
  Provides the list of External Listener Services in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_listener_services
This data source provides the list of External Listener Services in Oracle Cloud Infrastructure Database Management service.

Lists the database services registered with the specified external listener
for the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_external_listener_services" "test_external_listener_services" {
	#Required
	external_listener_id = oci_database_management_external_listener.test_external_listener.id
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	opc_named_credential_id = var.external_listener_service_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `external_listener_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external listener.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `external_listener_service_collection` - The list of external_listener_service_collection.

### ExternalListenerService Reference

The following attributes are exported:

* `items` - An array of external listener services.
	* `listener_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external listener.
	* `managed_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	* `name` - The name of the service.

