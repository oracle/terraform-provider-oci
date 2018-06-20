# oci_identity_tag_namespace

## TagNamespace Resource

### TagNamespace Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the tag namespace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag namespace.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag namespace.
* `is_retired` - Whether the tag namespace is retired. For more information, see [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name of the tag namespace. It must be unique across all tag namespaces in the tenancy and cannot be changed. 
* `time_created` - Date and time the tagNamespace was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new tag namespace in the specified compartment.

You must specify the compartment ID in the request object (remember that the tenancy is simply the root
compartment).

You must also specify a *name* for the namespace, which must be unique across all namespaces in your tenancy
and cannot be changed. The name can contain any ASCII character except the space (_) or period (.).
Names are case insensitive. That means, for example, "myNamespace" and "mynamespace" are not allowed
in the same tenancy. Once you created a namespace, you cannot change the name.
If you specify a name that's already in use in the tenancy, a 409 error is returned.

You must also specify a *description* for the namespace.
It does not have to be unique, and you can change it with
[UpdateTagNamespace](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/tagging/20170101/TagNamespace/UpdateTagNamespace).

Tag namespaces cannot be deleted, but they can be retired.
See [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring) for more information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the tag namespace.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) The description you assign to the tag namespace during creation.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the tag namespace during creation. It must be unique across all tag namespaces in the tenancy and cannot be changed.
* `is_retired` - (Optional) Whether the tag namespace is retired. For more information, see [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
 


### Update Operation
Updates the the specified tag namespace. You can't update the namespace name.

Updating `isRetired` to 'true' retires the namespace and all the tag definitions in the namespace. Reactivating a
namespace (changing `isRetired` from 'true' to 'false') does not reactivate tag definitions.
To reactivate the tag definitions, you must reactivate each one individually *after* you reactivate the namespace,
using [UpdateTag](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Tag/UpdateTag). For more information about retiring tag namespaces, see
[Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring).

You can't add a namespace with the same name as a retired namespace in the same tenancy.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag namespace during creation.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_retired` - Whether the tag namespace is retired. For more information, see [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_identity_tag_namespace" "test_tag_namespace" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.tag_namespace_description}"
	name = "${var.tag_namespace_name}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_retired = false
}
```

# oci_identity_tag_namespaces

## TagNamespace DataSource

Gets a list of tag_namespaces.

### List Operation
Lists the tag namespaces in the specified compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `include_subcompartments` - (Optional) An optional boolean parameter indicating whether to retrieve all tag namespaces in subcompartments. If this parameter is not specified, only the tag namespaces defined in the specified compartment are retrieved. 


The following attributes are exported:

* `tag_namespaces` - The list of tag_namespaces.

### Example Usage

```hcl
data "oci_identity_tag_namespaces" "test_tag_namespaces" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	include_subcompartments = "${var.tag_namespace_include_subcompartments}"
}
```