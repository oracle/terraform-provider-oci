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
result also indicates if the tenancy is onboarded with Logging Analytics.


## Example Usage

```hcl
data "oci_log_analytics_namespaces" "test_namespaces" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `namespace_collection` - The list of namespace_collection.

### Namespace Reference

The following attributes are exported:

* `compartment_id` - The is the tenancy ID
* `is_onboarded` - This indicates if the tenancy is onboarded to Logging Analytics
* `namespace` - This is the namespace name of a tenancy

