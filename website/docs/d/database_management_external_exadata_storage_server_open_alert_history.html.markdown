---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_storage_server_open_alert_history"
sidebar_current: "docs-oci-datasource-database_management-external_exadata_storage_server_open_alert_history"
description: |-
  Provides details about a specific External Exadata Storage Server Open Alert History in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_exadata_storage_server_open_alert_history
This data source provides details about a specific External Exadata Storage Server Open Alert History resource in Oracle Cloud Infrastructure Database Management service.

Gets the open alerts from the specified Exadata storage server.


## Example Usage

```hcl
data "oci_database_management_external_exadata_storage_server_open_alert_history" "test_external_exadata_storage_server_open_alert_history" {
	#Required
	external_exadata_storage_server_id = oci_database_management_external_exadata_storage_server.test_external_exadata_storage_server.id
}
```

## Argument Reference

The following arguments are supported:

* `external_exadata_storage_server_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server.


## Attributes Reference

The following attributes are exported:

* `alerts` - A list of open alerts.
	* `message` - The alert message.
	* `severity` - The severity of the alert.
	* `time_start_at` - The start time of the alert.
	* `type` - The type of alert.

