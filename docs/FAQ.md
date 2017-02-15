    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# FAQ for the Oracle Bare Metal Cloud Services Terraform provider

#### Q: What is the Oracle Bare Metal Cloud?  
https://cloud.oracle.com/en_US/bare-metal/architecture

#### Q: What is Terraform?
Terraform is an orchestration engine and language that enables you to safely and predictably create, change, and improve production infrastructure. It is an open source tool that codifies APIs into declarative configuration files that can be shared amongst team members, treated as code, edited, reviewed, and versioned.

#### Q: What is a Terraform provider?
Terraform is agnostic to the underlying cloud platform, it implements this with the provider model. Providers are pieces of software that act as an interface between the Cloud provider and the Terraform engine. There is a list of additional providers here - https://www.terraform.io/docs/providers/.

#### Q: Can I use Terraform to manage my infrastructure in multiple clouds?
Yes, Terraform supports configurations that can span multiple clouds and can allow you to manage infrastructure and resources in those clouds.

#### Q: What can I do with Terraform I can't do with the APIs?
Terraform allows you to define infrastructure configurations and then have those configurations implemented/created by Terraform automatically. In this respect, you could compare Terraform to similar solutions like OpenStack Heat, AWS CloudFormation, and others. http://blog.scottlowe.org/2015/11/25/intro-to-terraform/

#### Q: Can I use Terraform to manage both Oracle Public Cloud and Oracle Bare Metal Cloud S?
**Yes.** The Oracle Public Cloud Terraform provider is located here - https://github.com/oracle/terraform-provider-compute.
See - http://darylscorner.com/2016/11/using-terraform-across-multiple-cloud-providers/ for examples of Terraform multi-cloud configurations.

#### Q: What happens if I change infrastructure I'm managing with Terraform outside of Terraform?
The change you made outside of Terraform will be overwritten the next time you apply the configuration unless you add the `ignore_changes` parameter to the resource in the configuration file.

#### Q: How do I get help?
You can file an issue against the project  
https://github.com/oracle/terraform-provider-baremetal/issues

or meet us in the OBMCS forums  
https://community.oracle.com/community/cloud_computing/bare-metal
