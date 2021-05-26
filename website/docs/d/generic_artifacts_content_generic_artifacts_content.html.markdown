---
subcategory: "Generic Artifacts Content"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generic_artifacts_content_generic_artifacts_content"
sidebar_current: "docs-oci-datasource-generic_artifacts_content-generic_artifacts_content"
description: |-
  Provides details about a specific Generic Artifacts Content in Oracle Cloud Infrastructure Generic Artifacts Content service
---

# Data Source: oci_generic_artifacts_content_generic_artifacts_content
This data source provides details about a specific Generic Artifacts Content resource in Oracle Cloud Infrastructure Generic Artifacts Content service.

Gets the specified artifact's content.

## Example Usage

```hcl
data "oci_generic_artifacts_content_generic_artifacts_content" "test_generic_artifacts_content" {
	#Required
	artifact_id = oci_generic_artifacts_content_artifact.test_artifact.id
}
```

## Argument Reference

The following arguments are supported:

* `artifact_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.  Example: `ocid1.genericartifact.oc1..exampleuniqueID` 


## Attributes Reference

The following attributes are exported:


