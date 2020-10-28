
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

List Namespaces.

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

* `compartment_id` - Tenancy ID
* `is_onboarded` - if tenancy is onboarded to logging analytics
* `namespace` - namespace name

