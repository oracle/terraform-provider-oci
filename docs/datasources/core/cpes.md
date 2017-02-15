# baremetal\_core\_cpe

List customer premise equipment objects (CPEs).

## Example Usage

```
data "baremetal_core_cpes" "s" {
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
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The CPE's Oracle ID (OCID).
* `ip_address` - The public IP address of the on-premise router.
* `time_created` - The date and time the image was created.
