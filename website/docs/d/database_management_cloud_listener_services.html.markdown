---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_listener_services"
sidebar_current: "docs-oci-datasource-database_management-cloud_listener_services"
description: |-
  Provides the list of Cloud Listener Services in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_listener_services
This data source provides the list of Cloud Listener Services in Oracle Cloud Infrastructure Database Management service.

Lists the database services registered with the specified cloud listener
for the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_cloud_listener_services" "test_cloud_listener_services" {
	#Required
	cloud_listener_id = oci_database_management_cloud_listener.test_cloud_listener.id
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	opc_named_credential_id = var.cloud_listener_service_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_listener_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud listener.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `cloud_listener_service_collection` - The list of cloud_listener_service_collection.

### CloudListenerService Reference

The following attributes are exported:

* `items` - An array of cloud listener services.
	* `listener_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud listener.
	* `managed_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	* `name` - The name of the service.

