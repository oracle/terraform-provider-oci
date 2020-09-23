---
layout: "oci"
page_title: "Best Practices"
sidebar_current: "docs-oci-guide-best_practices"
description: |-
  The Oracle Cloud Infrastructure provider. Best Practices
---

## Terraform Provider Best Practices

Following are recommended best practices for writing configurations for the Oracle Cloud Infrastructure Terraform provider.

###  Sensitive Data May Be Stored In Statefile

> **Warning**: The state contains all resource attributes that are specified as part of configuration files. If you manage any sensitive data with Terraform (like database or user passwords, instance or load balancer private keys, etc), treat the state itself as sensitive data. 
Please refer to [Sensitive Data in State](https://www.terraform.io/docs/state/sensitive-data.html) for more details. 


### Referencing Images

When launching Compute instances, your Terraform configuration should use the same image every time you execute a Terraform Apply command.

To ensure this, specify the image OCID directly, rather than locating it using the `oci_core_image` data source. 
This is because the `oci_core_image` data source calls into the ListImages API, whose return values can change over 
time as new images are periodically added and older ones deleted. For a list of Oracle-Provided images and their OCIDs, 
see [Oracle-Provided Images](https://docs.cloud.oracle.com/iaas/Content/Compute/References/images.htm). 
For more information, see the write up in this issue: [Results of oci_core_images will change over time for Oracle-provided images](https://github.com/oracle/terraform-provider-oci/issues/352).

We recommend the following pattern for specifying an image for a given region:

```hcl
variable "image_id" {
  type = "map"
  default = {
    // See https://docs.cloud.oracle.com/iaas/images/
    // Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}
```

A Compute instance can use this in the following way:

```hcl
resource "oci_core_instance" "TFInstance" {
  image = var.image_id[var.region]
  ...
}
```


### Availability Domains

With respect to Availability Domains, we caution against the common pattern of iterating over the results of the `oci_identity_availability_domains` data source, as shown here:

```hcl
// Get all availability domains for the region
data "oci_identity_availability_domains" "ads" {
  compartment_id = var.tenancy_ocid
}
  
// Then either use it to get a single AD name based on the index:
resource "oci_core_instance" "nat" {
  availability_domain = lookup(data.oci_identity_availability_domains.ads.availability_domains[var.nat_instance_ad],"name")
  ...
}
  
// Or iterate through all the ADs:
resource "oci_core_subnet" "nat" {
  count = length(data.oci_identity_availability_domains.ads.availability_domains)
  availability_domain = lookup(data.oci_identity_availability_domains.ad.availability_domains[count.index], "name")
  ...
}
```

The recommendation is to explicitly list the Availability Domain names for the regions in your configuration. To do so, use a variable that you have defined as follows:

```hcl
variable "ad_list" {
  type = "list"
}
```

You can then use the variable as shown here:

```hcl
// Index:
resource "oci_core_instance" "nat" {
  availability_domain = var.ad_list[var.nat_instance_ad_index]
  ...
}
  
// Or iterate through all the ADs:
resource "oci_core_subnet" "nat" {
  count = length(var.ad_list)
  availability_domain = var.ad_list[count.index]
  ...
}
```

You can then set the ad_list variable directly by using the availability domain names for your tenant and region, as shown here:

```hcl
variable "ad_list" {
  type = "list"
  default = ["kIdk:PHX-AD-1","kIdk:PHX-AD-2","kIdk:PHX-AD-3"]
}
```

The advantage of using this method is that it gives you control over your availability domain usage and prevents unexpected changes over time. 
However, this approach is problematic when configurations are shared between tenancies and regions, since availability domain names are tenancy and region-specific.

A convenient alternative is to instead set the ad_list value by using the oci_identity_availability_domains data source. 
You should do this in the configuration, then pass them into the resources or modules. This effectively centralizes the list of ADs, 
making it is easy to switch to an explicit list later, should that become necessary.

```hcl
data "oci_identity_availability_domains" "ad" {
  compartment_id = var.tenancy_ocid
}
 
data "template_file" "ad_names" {
  count = length(data.oci_identity_availability_domains.ad.availability_domains)
  template = lookup(data.oci_identity_availability_domains.ad.availability_domains[count.index], "name")
}
  
module "ssm_network" {
  ad_list = data.template_file.ad_names.*.rendered
  ...
}
```
