---
subcategory: "Generic Artifacts Content"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generic_artifacts_content_artifact_by_path"
sidebar_current: "docs-oci-datasource-generic_artifacts_content-artifact_by_path"
description: |-
  Provides details about a specific Artifact By Path in Oracle Cloud Infrastructure Generic Artifacts Content service
---

# Data Source: oci_generic_artifacts_content_artifact_by_path
This data source provides details about a specific Artifact By Path resource in Oracle Cloud Infrastructure Generic Artifacts Content service.

Get generic artifact content.

## Example Usage

```hcl
data "oci_generic_artifacts_content_artifact_by_path" "test_artifact_by_path" {
	#Required
	artifact_path = var.artifact_by_path_artifact_path
	repository_id = oci_artifacts_repository.test_repository.id
	version = var.artifact_by_path_version
}
```

## Argument Reference

The following arguments are supported:

* `artifact_path` - (Required) The generic artifact path.  Example: `foo/bar/baz` 
* `repository_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.  Example: `ocid1.repository.oc1..exampleuniqueID` 
* `version` - (Required) The generic artifact version.  Example: `1.1.2` 


## Attributes Reference

The following attributes are exported:


