KVM Virtualization
===============================
This folder contains sample terraform templates to allow you to run a KVM hypervisor on Oracle Cloud Infrastructure, based on the [Installing and Configuring KVM on Bare Metal Instances with Multi-VNIC](https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/installing_kvm_multi_vnics.pdf) white-paper and [A Simple Guide to Nested KVM Virtualization on Oracle Cloud Infrastructure](https://blogs.oracle.com/cloud-infrastructure/nested-kvm-virtualization-on-oracle-iaas) blog post.


Oracle Cloud Infrastructure does not offer a prepackaged image with KVM preconfigured.
However, we want to provide you with the ability to start using KVM within your environments and move virtual machines into the cloud.

Offering KVM in a cloud environment has two benefits. First, it allows the extension of existing on-premises KVM environments into Oracle Cloud Infrastructure. Second, it provides the ability to install legacy operating systems and prepackaged virtual machines to use within your environment.

The ability to bring the KVM hypervisor is unique to Oracle and is enabled by using several Oracle Cloud Infrastructure features:
* Block Volume service
* Bare Metal Compute & VM Instances
* Multiple virtual network interface cards (multi-VNIC) capabilities within the Network Service


On Oracle Cloud Infrastructure (OCI) you can use a Bare Metal Compute shape for your KVM hypervisor or a Virtual Machine Compute Shape on Nested Virtualization. On this folder you can find sample code for both scenarios. 

* [KVM on Bare Metal](./kvm-baremetal/README.md)
* [KVM Nested](./kvm-nested/README.md)
