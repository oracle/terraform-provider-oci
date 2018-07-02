## Tagging OCI Resources
When you have many resources (for example, instances, VCNs, load balancers, and block volumes) across multiple compartments in your tenancy, it can become difficult to track resources used for specific purposes, or to aggregate them, report on them, or take bulk actions on them. Tagging allows you to define keys and values and associate them with resources. You can then use the tags to help you organize and list resources based on your business needs. See [Tagging Overview](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#overview) to familiarize yourself with concept of tagging and features available.

## Managing Tags and Tag Namespaces
* See [tag_namespaces.md](https://github.com/oracle/terraform-provider-oci/blob/master/docs/identity/tag_namespaces.md) for guidance on managing lifecycle of tag namespaces.
* See [tags.md](https://github.com/oracle/terraform-provider-oci/blob/master/docs/identity/tag_namespaces.md) for guidance on managing lifecycle of tags.

## How To Manage Tags on OCI Resources
* **Freeform tags:** Freeform tags are simple key value map
* **Defined tags:** Defined tags provide a key/value map and are organized by combining the tag namespaces with tag keys using dot notation. For example, a tag namespace called `HumanResources` could have a key named `CostCenter`. You then associate the namespace and key `HumanResource.CostCenter` and then assign the desired tag, as shown in the following example.

### Examples:
#### Example 1:
```hcl
//Hand curated way
resource "oci_core_instance" "t" {
    .
    .
    .
    
    freeform_tags =  {
           Environment = "Prod"
           Department = "Ops"
   }
    defined_tags = {
        HumanResources.CostCenter = "42"
        Operations.Project = "Beta"
        HumanResources.Environment = "Production"
    }
}
```
#### Example 2:
```hcl
//Using Locals(available in terraform 0.10.3 or later) & interpolation

locals {
  //Put all common tags here
  common_tags = "${map(
                        "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}","value",
                        "HumanResources.Environment", "Production",
                        "Operations.Project", "Beta"
                        )}"
}

resource "oci_core_instance" t {
    .
    .
    .
  
    freeform_tags = "${map("key${count.index}", "value${count.index}",
        "domain", "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
        )}"
  
    defined_tags = "${merge(
                        local.common_tags,
                        map(
                            "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag2.name}", "awesome-tag-example"
                        )
                    )}"
  
}
```
### Taggable OCI Resources

* **Core**
    * [Console Histories](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/console_histories.md)
    * [CPEs](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/cpes.md)  
    * [DHCP Options](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/dhcp_options.md)
    * [DRGs](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/drgs.md)
    * [Images](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/images.md)
    * [Instance Console Connections](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/instance_console_connections.md)
    * [Instances](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/instances.md)
    * [Internet Gateways](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/internet_gateways.md)
    * [IPSec Connections](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/ip_sec_connections.md)
    * [Local Peering Gateways](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/local_peering_gateways.md)
    * [Private IPs](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/private_ips.md)
    * [Public IPs](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/public_ips.md)
    * [Route Tables](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/route_tables.md)
    * [Security Lists](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/security_lists.md)
    * [Subnets](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/subnets.md)
    * [VCNs](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/vcns.md)
    * [VNICs](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/vnics.md)
    * [Volume Backups](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/volume_backups.md)
    * [Volumes](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/volumes.md)
    * [VolumeGroups](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/volume_groups.md)
    * [VolumeGroupBackups](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/volume_group_backups.md)
* **Database**
    * [DBSystems](https://github.com/oracle/terraform-provider-oci/tree/master/docs/database/db_systems.md)
* **Identity**
    * [Compartments](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/compartments.md)
    * [Groups](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/groups.md)
    * [Policies](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/policies.md)
    * [Tag Namespaces](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/tag_namespaces.md)
    * [Tags](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/tags.md)
    * [Tenancies](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/tenancies.md)
    * [Users](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/users.md)
* **Object Storage**
    * [Buckets](https://github.com/oracle/terraform-provider-oci/tree/master/docs/object_storage/buckets.md)