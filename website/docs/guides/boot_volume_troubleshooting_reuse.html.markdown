---
layout: "oci"
page_title: "Boot Volume reuse & troubleshooting"
sidebar_current: "docs-oci-guide-boot_volume_troubleshooting"
description: |-
  The Oracle Cloud Infrastructure provider. Boot Volume reuse and troubleshooting
---


## OCI Terraform Provider Boot volume reuse and troubleshooting
This guide details the following scenarios:<br/>
1. Preserving boot volume when performing instance scaling<br/>
2. Boot volume troubleshooting and repair

To read more about boot volumes, see [Boot Volumes](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumes.htm)


### Preserve boot volume with instance scaling

You may want to upscale / downscale instances (change instance shape) while using the same boot volume. When you terminate 
your instance, you can keep the associated boot volume and use it to launch a new instance using a different instance type
or shape.

To achieve this, you need to detach the boot volume from the running instance. This can be performed by either 
terminating the instance while preserving the boot volume or by stopping the instance and detaching the boot volume,

All terraform resources of type `oci_core_instance` have the parameter `preserve_boot_volume` set as true by default. 
This parameter ensures that upon termination of the instance, the attached boot volume is not terminated.

```
resource "oci_core_instance" "TFInstance" {
  ...
  state = "STOPPED"                  // set this state to stop the instance
  preserve_boot_volume = true
}

output "bootVolumeFromInstance" {
  value = ["${oci_core_instance.TFInstance.boot_volume_id}"]
}
```

Once the boot volume is detached, the OCID of the boot volume can be referred as the source of the new instance, as 
illustrated below

```
resource "oci_core_instance" "TFScaleInstance" {
  ...
  source_details {
    source_type = "bootVolume"
    
    // reference the original boot volume id here
    source_id   = "ocid1.bootvolume.oc1.phx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"   
  }
}
```


<br/>
### OCI Terraform Provider boot volume troubleshooting and repair

If you think a boot volume issue is causing a compute instance problem, you can stop the instance and detach the boot volume. 
Then you can attach it to another instance as a data volume to troubleshoot it. After resolving the issue, you can then 
reattach it to the original instance or use it to launch a new instance. 

Once the boot volume has been detached, the OCID of the boot volume can be referred as the block volume parameter for 
another instance.

```
resource "oci_core_volume_attachment" "TFBlockAttach" {
  ...
  attachment_type = "iscsi"
  compartment_id  = "ocid1.compartment.oc1..xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  
  // new instance
  instance_id     = "ocid1.instance.oc1.phx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"    
  
  // attach the boot volume as a block volume
  volume_id       = "ocid1.bootvolume.oc1.phx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"                                                                                    
}
```

Once you have resolved the issue, detach this volume from the second instance and attach it as a boot volume to the 
original instance.

```
resource "oci_core_instance" "TFScaleInstance" {
  ...

  source_details {
    source_type = "bootVolume"
    
    // attach back as boot volume
    // reference the volume id here
    source_id   = "ocid1.bootvolume.oc1.phx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"       
  }
}
```

