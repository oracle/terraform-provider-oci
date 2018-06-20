# oci_identity_tag

## Tag Resource

### Tag Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag definition.
* `is_retired` - Whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name of the tag. The name must be unique across all tags in the tag namespace and can't be changed. 
* `tag_namespace_id` - The OCID of the namespace that contains the tag definition.
* `time_created` - Date and time the tag was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new tag in the specified tag namespace.

You must specify either the OCID or the name of the tag namespace that will contain this tag definition.

You must also specify a *name* for the tag, which must be unique across all tags in the tag namespace
and cannot be changed. The name can contain any ASCII character except the space (_) or period (.) characters.
Names are case insensitive. That means, for example, "myTag" and "mytag" are not allowed in the same namespace.
If you specify a name that's already in use in the tag namespace, a 409 error is returned.

You must also specify a *description* for the tag.
It does not have to be unique, and you can change it with
[UpdateTag](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/tagging/20170901/Tag/UpdateTag).


The following arguments are supported:

* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) The description you assign to the tag during creation.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the tag during creation. The name must be unique within the tag namespace and cannot be changed. 
* `is_retired` - Whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `tag_namespace_id` - (Required) The OCID of the tag namespace. 


### Update Operation
Updates the the specified tag definition. You can update `description`, and `isRetired`.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_retired` - Whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_identity_tag" "test_tag" {
	#Required
	description = "${var.tag_description}"
	name = "${var.tag_name}"
	tag_namespace_id = "${oci_identity_tag_namespace.test_tag_namespace.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_retired = false
}
```

# oci_identity_tags

## Tag DataSource

Gets a list of tags.

### List Operation
Lists the tag definitions in the specified tag namespace.

The following arguments are supported:

* `tag_namespace_id` - (Required) The OCID of the tag namespace. 


The following attributes are exported:

* `tags` - The list of tags.

### Example Usage

```hcl
data "oci_identity_tags" "test_tags" {
	#Required
	tag_namespace_id = "${oci_identity_tag_namespace.test_tag_namespace.id}"
}
```