# oci_core_volume_backup_policy_assignment

## VolumeBackupPolicyAssignment Resource

### VolumeBackupPolicyAssignment Reference

The following attributes are exported:

* `asset_id` - The OCID of the asset (e.g. a volume) to which the policy has been assigned.
* `id` - The OCID of the volume backup policy assignment.
* `policy_id` - The OCID of the volume backup policy that has been assigned to an asset.
* `time_created` - The date and time the volume backup policy assignment was created. Format defined by RFC3339. 



### Create Operation
Assigns a policy to the specified asset, such as a volume. Note that a given asset can
only have one policy assigned to it; if this method is called for an asset that previously
has a different policy assigned, the prior assignment will be silently deleted.


The following arguments are supported:

* `asset_id` - (Required) The OCID of the asset (e.g. a volume) to which to assign the policy.
* `policy_id` - (Required) The OCID of the volume backup policy to assign to an asset.


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_volume_backup_policy_assignment" "test_volume_backup_policy_assignment" {
	#Required
	asset_id = "${oci_core_asset.test_asset.id}"
	policy_id = "${oci_core_policy.test_policy.id}"
}
```

# oci_core_volume_backup_policy_assignments

## VolumeBackupPolicyAssignment DataSource

Gets a list of volume_backup_policy_assignments.

### List Operation
Gets the volume backup policy assignment for the specified asset. Note that the
assetId query parameter is required, and that the returned list will contain at most
one item (since any given asset can only have one policy assigned to it).

The following arguments are supported:

* `asset_id` - (Required) The OCID of an asset (e.g. a volume).


The following attributes are exported:

* `volume_backup_policy_assignments` - The list of volume_backup_policy_assignments.

### Example Usage

```hcl
data "oci_core_volume_backup_policy_assignments" "test_volume_backup_policy_assignments" {
	#Required
	asset_id = "${oci_core_asset.test_asset.id}"
}
```