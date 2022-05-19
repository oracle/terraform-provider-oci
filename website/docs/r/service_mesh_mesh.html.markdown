---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_mesh"
sidebar_current: "docs-oci-resource-service_mesh-mesh"
description: |-
  Provides the Mesh resource in Oracle Cloud Infrastructure Service Mesh service
---

# oci_service_mesh_mesh
This resource provides the Mesh resource in Oracle Cloud Infrastructure Service Mesh service.

Creates a new Mesh.


## Example Usage

```hcl
resource "oci_service_mesh_mesh" "test_mesh" {
	#Required
	certificate_authorities {
		#Required
		id = var.mesh_certificate_authorities_id
	}
	compartment_id = var.compartment_id
	display_name = var.mesh_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.mesh_description
	freeform_tags = {"bar-key"= "value"}
	mtls {
		#Required
		minimum = var.mesh_mtls_minimum
	}
}
```

## Argument Reference

The following arguments are supported:

* `certificate_authorities` - (Required) The OCID of the certificate authority resource OCID to use for creating leaf certificates.
	* `id` - (Required) The OCID of the certificate authority resource.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `display_name` - (Required) (Updatable) A user-friendly name. The name does not have to be unique and can be changed after creation. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `mtls` - (Optional) (Updatable) Sets a minimum level of mTLS authentication for all virtual services within the mesh.
	* `minimum` - (Required) (Updatable) DISABLED: No minimum virtual services within this mesh can use any mTLS authentication mode. PERMISSIVE: Virtual services within this mesh can use either PERMISSIVE or STRICT modes. STRICT: All virtual services within this mesh must use STRICT mode. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `certificate_authorities` - A list of certificate authority resources to use for creating leaf certificates for mTLS authentication. Currently we only support one certificate authority, but this may expand in future releases. 
	* `id` - The OCID of the certificate authority resource.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `display_name` - A user-friendly name. The name does not have to be unique and can be changed after creation. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `mtls` - Sets a minimum level of mTLS authentication for all virtual services within the mesh.
	* `minimum` - DISABLED: No minimum virtual services within this mesh can use any mTLS authentication mode. PERMISSIVE: Virtual services within this mesh can use either PERMISSIVE or STRICT modes. STRICT: All virtual services within this mesh must use STRICT mode. 
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Mesh
	* `update` - (Defaults to 20 minutes), when updating the Mesh
	* `delete` - (Defaults to 20 minutes), when destroying the Mesh


## Import

Meshes can be imported using the `id`, e.g.

```
$ terraform import oci_service_mesh_mesh.test_mesh "id"
```

