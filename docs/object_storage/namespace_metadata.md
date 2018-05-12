# oci_object_storage_namespace_metadata

## NamespaceMetadata Resource

### NamespaceMetadata Reference

The following attributes are exported:

* `default_s3compartment_id` - The default compartment ID for an S3 client.
* `default_swift_compartment_id` - The default compartment ID for a Swift client.
* `namespace` - The namespace to which the metadata belongs.



### Create Operation


The following arguments are supported:

* `default_s3compartment_id` - (Optional) The default compartment ID for an S3 client.
* `default_swift_compartment_id` - (Optional) The default compartment ID for a Swift client.
* `namespace` - (Required) The namespace to which the metadata belongs.

### Update Operation
Change the default Swift/S3 compartmentId of user's namespace into the user-defined compartmentId. Upon doing
this, all subsequent bucket creations will use the new default compartment, but no previously created
buckets will be modified. A user must have the NAMESPACE_UPDATE permission to make changes to the default
compartments for S3 and Swift.


The following arguments support updates:
* `default_s3compartment_id` - The default compartment ID for an S3 client.
* `default_swift_compartment_id` - The default compartment ID for a Swift client.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
  namespace = "${var.bucket_namespace}"
  default_s3compartment_id = "${var.default_s3compartment_id}"
  default_swift_compartment_id = "${var.default_swift_compartment_id}"
}
```

# oci_object_storage_namespace_metadata

## NamespaceMetadata DataSource

Gets a single namespace_metadata

### Get Operation
Namespaces are unique. Namespaces are either the tenancy name or a random string automatically generated during
account creation. You cannot edit a namespace.

The following arguments are supported:
* `namespace` - (Required) The namespace to which the metadata belongs.


The following attributes are exported:
* `default_s3compartment_id` - The default compartment ID for an S3 client.
* `default_swift_compartment_id` - The default compartment ID for a Swift client.
* `namespace` - The namespace to which the metadata belongs.


### Example Usage

```hcl
data "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
  namespace = "${var.bucket_namespace}"
}
```