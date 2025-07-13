---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_group_artifact_content"
sidebar_current: "docs-oci-datasource-datascience-model_group_artifact_content"
description: |-
  Provides details about a specific Model Group Artifact Content in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_group_artifact_content
This data source provides details about a specific Model Group Artifact Content resource in Oracle Cloud Infrastructure Data Science service.

Downloads the model artifact for the specified model group.

## Example Usage

```hcl
data "oci_datascience_model_group_artifact_content" "test_model_group_artifact_content" {
	#Required
	model_group_id = oci_datascience_model_group.test_model_group.id

	#Optional
	range = var.model_group_artifact_content_range
}
```

## Argument Reference

The following arguments are supported:

* `model_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroup.
* `range` - (Optional) Optional byte range to fetch, as described in [RFC 7233](https://tools.ietf.org/html/rfc7232#section-2.1), section 2.1. Note that only a single range of bytes is supported. 


## Attributes Reference

The following attributes are exported:


