
# oci_identity_regions

## Region DataSource

Gets a list of regions.

### List Operation
Lists all the regions offered by Oracle Cloud Infrastructure.

The following attributes are exported:

* `regions` - The list of regions.

### Example Usage

```hcl
data "oci_identity_regions" "test_regions" {
}
```
### Region Reference

The following attributes are exported:

* `key` - The key of the region.  Allowed values are: - `PHX` - `IAD` - `FRA` - `LHR` 
* `name` - The name of the region.  Allowed values are: - `us-phoenix-1` - `us-ashburn-1` - `eu-frankfurt-1` - `uk-london-1` 
