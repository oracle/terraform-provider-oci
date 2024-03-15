---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_query_quick_picks"
sidebar_current: "docs-oci-datasource-apm_traces-query_quick_picks"
description: |-
  Provides the list of Query Quick Picks in Oracle Cloud Infrastructure Apm Traces service
---

# Data Source: oci_apm_traces_query_quick_picks
This data source provides the list of Query Quick Picks in Oracle Cloud Infrastructure Apm Traces service.

Returns a list of predefined Quick Pick queries intended to assist the user
to choose a query to run.  There is no sorting applied on the results.


## Example Usage

```hcl
data "oci_apm_traces_query_quick_picks" "test_query_quick_picks" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID for the intended request. 


## Attributes Reference

The following attributes are exported:

* `quick_picks` - The list of quick_picks.

### QueryQuickPick Reference

The following attributes are exported:

* `quick_pick_name` - Quick Pick name for the query. 
* `quick_pick_query` - Query for the Quick Pick. 

