# oci_core_instance_credentials

## InstanceCredential Singular DataSource

### InstanceCredential Reference

The following attributes are exported:

* `password` - The password for the username.
* `username` - The username.



### Get Operation
Gets the generated credentials for the instance. Only works for Windows instances. The returned credentials
are only valid for the initial login.


The following arguments are supported:

* `instance_id` - (Required) The OCID of the instance.


### Example Usage

```hcl
data "oci_core_instance_credentials" "test_instance_credential" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
}
```
