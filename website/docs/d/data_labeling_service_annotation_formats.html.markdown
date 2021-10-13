---
subcategory: "Data Labeling Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_labeling_service_annotation_formats"
sidebar_current: "docs-oci-datasource-data_labeling_service-annotation_formats"
description: |-
  Provides the list of Annotation Formats in Oracle Cloud Infrastructure Data Labeling Service service
---

# Data Source: oci_data_labeling_service_annotation_formats
This data source provides the list of Annotation Formats in Oracle Cloud Infrastructure Data Labeling Service service.

These are a static list in a given region.

## Example Usage

```hcl
data "oci_data_labeling_service_annotation_formats" "test_annotation_formats" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `annotation_format_collection` - The list of annotation_format_collection.

### AnnotationFormat Reference

The following attributes are exported:

* `items` - List of annotation formats.
	* `name` - A unique name for the target AnnotationFormat for the Dataset. 

