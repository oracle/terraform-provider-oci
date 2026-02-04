---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespaces"
sidebar_current: "docs-oci-datasource-log_analytics-namespaces"
description: |-
  Provides the list of Namespaces in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespaces
This data source provides the list of Namespaces in Oracle Cloud Infrastructure Log Analytics service.

Given a tenancy OCID, this API returns the namespace of the tenancy if it is valid and subscribed to the region.  The
result also indicates if the tenancy is onboarded with Log Analytics.


## Example Usage

```hcl
data "oci_log_analytics_namespaces" "test_namespaces" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	is_compartment_delete = var.namespace_is_compartment_delete
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `is_compartment_delete` - (Optional) if true, the request is from compartment delete service. 


## Attributes Reference

The following attributes are exported:

* `namespace_collection` - The list of namespace_collection.

### Namespace Reference

The following attributes are exported:

* `compartment_id` - This is the tenancy ID
* `is_archiving_enabled` - This indicates if old data can be archived for a tenancy
* `is_onboarded` - This indicates if the tenancy is onboarded to Log Analytics
* `namespace` - This is the namespace name of a tenancy
* `state` - The current state of the compartment.
* `is_logSet_enabled` - This indicates if the tenancy is logSet enable
* `is_data_ever_ingested` - This indicates if the tenancy is data ever ingested
