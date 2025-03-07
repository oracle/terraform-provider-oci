---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_defined_metadata_artifact_content"
sidebar_current: "docs-oci-datasource-datascience-model_defined_metadata_artifact_content"
description: |-
  Provides details about a specific Model Defined Metadata Artifact Content in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_defined_metadata_artifact_content
This data source provides details about a specific Model Defined Metadata Artifact Content resource in Oracle Cloud Infrastructure Data Science service.

Downloads model defined metadata artifact content for specified model metadata key.

## Example Usage

```hcl
data "oci_datascience_model_defined_metadata_artifact_content" "test_model_defined_metadata_artifact_content" {
	#Required
	metadatum_key_name = oci_kms_key.test_key.name
	model_id = oci_datascience_model.test_model.id

	#Optional
	range = var.model_defined_metadata_artifact_content_range
}
```

## Argument Reference

The following arguments are supported:

* `metadatum_key_name` - (Required) The name of the model metadatum in the metadata.
* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `range` - (Optional) Optional byte range to fetch, as described in [RFC 7233](https://tools.ietf.org/html/rfc7232#section-2.1), section 2.1. Note that only a single range of bytes is supported. 


## Attributes Reference

The following attributes are exported:


