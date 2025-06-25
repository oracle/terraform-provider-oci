---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity_type"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_entity_type"
description: |-
  Provides details about a specific Log Analytics Entity Type in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_entity_type
This data source provides details about a specific Log Analytics Entity Type resource in Oracle Cloud Infrastructure Log Analytics service.

Retrieve the log analytics entity type with the given name.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_entity_type" "test_log_analytics_entity_type" {
	#Required
	entity_type_name = var.log_analytics_entity_type_name
	namespace = var.log_analytics_entity_type_namespace
}
```

## Argument Reference

The following arguments are supported:

* `entity_type_name` - (Required) Log analytics entity type name. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `category` - Log analytics entity type category. Category will be used for grouping and filtering. 
* `cloud_type` - Log analytics entity type group. That can be CLOUD (OCI) or NON_CLOUD otherwise. 
* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `internal_name` - Internal name for the log analytics entity type. 
* `name` - Log analytics entity type name. 
* `properties` - The parameters used in file patterns specified in log sources for this log analytics entity type. 
	* `description` - Description for the log analytics entity type property. 
	* `name` - Log analytics entity type property name. 
* `state` - The current lifecycle state of the log analytics entity type.
* `time_created` - Time the log analytics entity type was created. An RFC3339 formatted datetime string. 
* `time_updated` - Time the log analytics entity type was updated. An RFC3339 formatted datetime string. 

