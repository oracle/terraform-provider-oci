# oci_audit_configuration

## Configuration Resource

### Configuration Reference

The following attributes are exported:

* `retention_period_days` - The retention period days



### Create Operation
Create the configuration

The following arguments are supported:
* `compartment_id` - (Required) ID of the root compartment (tenancy)
* `retention_period_days` - (Optional) The retention period days


### Update Operation
Update the configuration

The following arguments support updates:
* `retention_period_days` - (Optional) The retention period days


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_audit_configuration" "test_configuration" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  retention_period_days = "${var.configuration_retention_period_days}"
}
```


## Configuration Singular DataSource


### Get Operation
Get the configuration

The following arguments are supported:

* `compartment_id` - (Required) ID of the root compartment (tenancy)


### Example Usage

```hcl
data "oci_audit_configuration" "test_configuration" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```
