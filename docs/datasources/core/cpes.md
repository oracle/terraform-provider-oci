# oci\_core\_cpe

[Cpe Reference][7f9e168e]

  [7f9e168e]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/ "CpeReference"

List Customer-Premises Equipment objects (CPEs). A virtual representation of the actual router on-premises at your site at your end of the IPSec VPN connection.

## Example Usage

```
data "oci_core_cpes" "s" {
  compartment_id = "compartmentid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The page to fetch.

## Attributes Reference

* `cpes` - The list of CPEs.

## CPEs reference
* `compartment_id` - The OCID of the compartment containing the CPE.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The CPE's Oracle ID (OCID).
* `ip_address` - The public IP address of the on-premises router.
* `time_created` - The date and time the image was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
