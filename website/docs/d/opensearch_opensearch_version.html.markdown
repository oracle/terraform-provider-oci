---
subcategory: "Opensearch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opensearch_opensearch_version"
sidebar_current: "docs-oci-datasource-opensearch-opensearch_version"
description: |-
Provides details about a specific Opensearch Version in Oracle Cloud Infrastructure Opensearch service
---

# Data Source: oci_opensearch_opensearch_version
This data source provides details about a specific Opensearch Version resource in Oracle Cloud Infrastructure Opensearch service.

Lists the supported Opensearch versions


## Example Usage

```hcl
data "oci_opensearch_opensearch_version" "test_opensearch_version" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `items` - A list of OpenSearch versions.
    * `version` - The version of OpenSearch.
