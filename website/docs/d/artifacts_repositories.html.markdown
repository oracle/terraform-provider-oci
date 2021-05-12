---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_repositories"
sidebar_current: "docs-oci-datasource-artifacts-repositories"
description: |-
  Provides the list of Repositories in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_repositories
This data source provides the list of Repositories in Oracle Cloud Infrastructure Artifacts service.

Lists repositories in the specified compartment.

## Example Usage

```hcl
data "oci_artifacts_repositories" "test_repositories" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.repository_display_name
	id = var.repository_id
	is_immutable = var.repository_is_immutable
	state = var.repository_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `id` - (Optional) A filter to return the resources for the specified OCID. 
* `is_immutable` - (Optional) A filter to return resources that match the isImmutable value. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state name exactly. 


## Attributes Reference

The following attributes are exported:

* `repository_collection` - The list of repository_collection.

### Repository Reference

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

