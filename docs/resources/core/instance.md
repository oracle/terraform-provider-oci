
# baremetal\_core\_instances

Gets a list of instances.

## Example Usage

```
resource "baremetal_core_instance" "t" {
    availability_domain = "availability_domain"
    compartment_id = "compartment_id"
    display_name = "display_name"
    image = "imageid"
    shape = "shapeid"
    subnet_id = "subnetid"
    metadata {
        ssh_authorized_keys = "mypublickey"
    }
    extended_metadata {
        some_string = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
    }
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `shape` - (Required) The shape of an instance.
* `subnet_id` - (Required) The OCID of the subnet.
* `availability_domain` - (Optional) The name of the Availability Domain.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `image_id` - (Required) The OCID of the image used to boot the instance.
* `metadata` - (Optional) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.
* `extended_metadata` - (Optional) Like metadata but allows nested metadata if you pass a valid JSON string as a value

## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

## Instance Reference
* `availability_domain` - The Availability Domain the instance is running in.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The OCID of the instance.
* `image_id` - The image used to boot the instance. You can enumerate all available images by calling ListImages.
* `state` - The current state of the instance: [PROVISIONING, RUNNING, STARTING, STOPPING, STOPPED, CREATING_IMAGE, TERMINATING, TERMINATED]
* `metadata` - Custom metadata that you provide.
* `extended_metadata` - Custom nested metadata that you provide. If you pass in a valid JSON string as a value then it will be converted to a JSON object; otherwise we will take the string value.
* `region` - The region that contains the Availability Domain the instance is running in.
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance.
* `time_created` - The date and time the instance was created.

* `public_ip` - The public ip of instance vnic (if enabled).
* `private_ip` - The private ip of instance vnic (if enabled).
