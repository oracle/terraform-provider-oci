---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_repository"
sidebar_current: "docs-oci-datasource-artifacts-repository"
description: |-
  Provides details about a specific Repository in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_repository
This data source provides details about a specific Repository resource in Oracle Cloud Infrastructure Artifacts service.

Gets the specified repository's information.

## Example Usage

```hcl
data "oci_artifacts_repository" "test_repository" {
	#Required
	repository_id = oci_artifacts_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `repository_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.  Example: `ocid1.artifactrepository.oc1..exampleuniqueID` 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the repository's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The repository description.
* `display_name` - The repository name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.  Example: `ocid1.artifactrepository.oc1..exampleuniqueID` 
* `is_immutable` - Whether the repository is immutable. The artifacts of an immutable repository cannot be overwritten.
* `repository_type` - The repository's supported artifact type.
* `state` - The current state of the repository.
* `time_created` - An RFC 3339 timestamp indicating when the repository was created.

