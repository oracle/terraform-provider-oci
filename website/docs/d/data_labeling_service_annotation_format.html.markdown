---
subcategory: "Data Labeling Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_labeling_service_annotation_format"
sidebar_current: "docs-oci-datasource-data_labeling_service-annotation_format"
description: |-
  Provides details about a specific Annotation Format in Oracle Cloud Infrastructure Data Labeling Service service
---

# Data Source: oci_data_labeling_service_annotation_format
This data source provides details about a specific Annotation Format resource in Oracle Cloud Infrastructure Data Labeling Service service.

These are a static list in a given region.

## Example Usage

```hcl
data "oci_data_labeling_service_annotation_format" "test_annotation_format" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `items` - List of annotation formats.
	* `name` - A unique name for the target AnnotationFormat for the Dataset. 

