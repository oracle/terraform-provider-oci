This solution provides a method to deploy a completely realized VCN including subnets, routing tables, and basic security lists.  The template also deploys an Internet Gateway for public subnets.  

In addition, the template allows you to build an arbitrary number of subnets, spread across ADs as desired, each with unique subnet masks (within the bounds and address sizes of the VCN, of course), and specify which of the subnets are public and which are private.  

Of course, if you do not want to do that much work, you can simply specify which subnets are public and which are private, and have the template evenly balance the subnet across the ADs.

Or, you can specify the public/private-ness of the subnet and pick what ADs the belong in.

Other features of the template:

- Automatic generation of display names for each resource created
- Creation of specific routing tables for public and private subnets
- Creation of a single global security list, plus a per-subnet security list
- Auto population of common basic entries into the security lists

The naming standard is one that I use within my tenancy, but you can alter them to meet your needs.  All standards are embedded in the templates.tf file in the root of the tempalte.  Read: https://www.terraform.io/docs/providers/template/d/file.html on how to use templates to create your own naming standards.

Please consult the Changelog for the latest changes to this process!

There are several prerequisites:

1. Python 2.7.x must be installed and the path to the Python binary identified.

2. You must identify an address space you want to use in the form of a CIDR specification.  For instance you can use "10.0.0.0/24" if you like.

3. The number of desired subnets of this VCN needs to be identified.

4. For each subnet, you must identify whether they are to be public, or private - meaning that you need to specify if you want the subnet to access the Internet directly, or to only have access to the private address space of the VCN.

5. The following must be specified in your shell environment (prefixed with TF_VAR_ of course):
    - tenancy_ocid
    - user_ocid
    - fingerprint
    - private_key_path
    - private_key_password (if required)
    - ssh_public_key (the actual public key, not the file)
    - region

NOTE: A template env-vars file is provided as part of this example.  Simply complete the items inside the template and source the result into your shell by using:

. ./env-vars    

Optionally:

1. Identify how you want to spread the subnets across the available ADs.  You can pick specific ADs for some or all of the subnets (see below).

2. Identify how much of the address space you want to give to each subnet by specifying the bit portion of the CIDR block (the number after the '/' in a CIDR specification) you want to give to each subnet.  Make sure you understand the limiations of splitting the subnets.

How to use this template:

1. Set your environment variables
2. Edit the ./modules/calc/datasources.tf file.  Change the first entry inside the "program" line to point to your Python binary.  For example, if your python binary is at /usr/bin/python, change the entry from:

program = [ "/usr/local/bin/python2.7", "${path.module}/tfnetnum.py"]

to:

program = [ "/usr/bin/python", "${path.module}/tfnetnum.py"]

*BE SURE NOT TO MAKE ANY OTHER CHANGES*

3. Edit the variables.tf file.  In this file there are a few items to customize to your liking:
	In the 'vcn_state' map variable - 
	- cidr_block: Set the IP address space and netmask in this portion of the variable. 
	              Use the standard format of 'IP Address/Netmaks'
	- compartment: Set the *Display Name* of the compartment in which to create the VCN
	
	In the 'subnet' *map variable* -
	- public: Keeping the [] brackets, for each subnet you are going to create, identify
	          whether the subnet should allow Public IPs (set to 'true'), or prevent
	          Public IPs (set to 'false').  Make sure that a) there are commas between the
	          entries, and b) there are the same number of entries as subnets you want to
	          to create. *The template uses this field to calculate the number of subnets
	          to create!*
	- ad: Again, keeping the [] brackets, for each subnet you are going to create,
	      identify which AD by *number only (1, 2, or 3)* you want the subnet to be
	      created in. If you want the template to pick the subnets for you, enter a '0'
	      as the AD number.  *Just as with the public variable, you must specify an entry
	      for each subnet*
	      
	To specify the netmask on a per subnet basis, complete the following two additional 
	tasks -
	- subnet_masks: For each subnet to be created, specify the desired subnet mask for
	                that subnet.  The masks will correspond to the listing of the ADs you
	                created in the 'ad' variable.  See below for an example.
	- In the *module "subnet"* section: Remove the comment indicator ('#') from in front 
	  of the 'source = "./modules/calc" line, and put it in front of the 
	  'source = "./modules/auto"' line.  
4. Save the variables.tf file.
5. Run 'terraform init'
6. Run 'terraform plan'.  Do this.  Really.  Review the output and make sure it is creating the subnets in the order you expect with the sizes you expect.  This also acts as a check against trying to create custom subnets that do not fit within the VCN (it throws an error on the last subnet telling you it will not fit).
7. Run 'terraform apply'.  Don't blink - it is a fast process.
8. Perform any additional configuration on your newly created VCN infrastructure.
9. Deploy lots of instances and services.  Build cool applications. Have fun.

Examples:

With great power comes complexity.  So here are some examples on how this can be used to give you a starting off point.  I am going to use the following for each example:

VCN CIDR range: 10.100.0.0/24
Compartment name: "my_compartment"

For all the examples, you would set the cidr_block value to the VCN CIDR range value, and the compartment value to the Compartment name value.

Specific configurations -

- I want to build a VCN with 6 subnets spread evenly across the ADs. The first subnet created should be public, and all the rest should be private IPs only -

If this is the configuration you want, set the following:

public = [ true, false, false, false, false, false ]
ad = [ 0, 0, 0, 0, 0, 0 ]

Leave the module "subnet" line to 'source = "./modules/auto"'

- I want to build a VCN that has 2 subnets in AD1, 3 in AD2, and 3 in AD3.  I want the VCN address space to be spread across them evenly.  The first subnet in AD2 should be public. -

If this is the configuration you want, set the following:

public = [ false, false, true, false, false, false, false ]
ad = [ 1, 1, 2, 2, 2, 3, 3, 3 ]

This also could have been entered as:
public = [ true, false, false, false, false, false, false ]
ad = [ 2, 1, 1, 2, 2, 3, 3, 3 ]

The order does not matter as long as the configuration between the public entry and the AD is consistent. 

Leave the module "subnet" line to 'source = "./modules/auto"'

- I want to build a VCN with 9 subnets, 3 subnets per AD.  There should be a /30 in each subnet with public IPs.  The remainder of the subnets should be evenly spread across the subnets and should be as large as possible. -

This example requires some thought.  I will not go into how to calculate the sizes, but the best possible way to do this is the following:

public = [ false, false, false, false, false, false, true, true, true ]
ad = [ 1, 2, 3, 1, 2, 3, 1, 2, 3 ]
subnet_masks = [ 27, 27, 27, 27, 27, 27, 30, 30, 30 ]

In the modules section, comment out the 'source = "./modules/auto"' line and uncommment the 'source = "./modules/calc"' line.

This configuration will build out 2 private subnets per AD with 32 addresses each, and 1 public subnet per AD with 2 possible addresses.

NOTE: You could have written this configuration this way...

public = [ true, false, false, true, false, false, true, false, false ]
ad = [ 1, 1, 1, 2, 2, 2, 3, 3, 3 ]
subnet_masks = [ 30, 27, 27, 30, 27, 27, 30, 27, 27 ]

...and it would have looked the same.  However, this would fail with an out of address space error.  Why?  The way Terraform calculates the address space offsets is sequentially.  So the first subnet takes the first part of the address space, the second the second, and so on.  By the time you get to the last address space, in our particular case, the /27 will not fit because of the boundary required.  There are two ways to fix this: a) reduce the /27 to a /29 or the last two to /28s, or b) increase the VCN space to a /23.

Limitations and defaults - 

I have built in some assumptions in the template as defaults.  They may work for your environment, and they may not, but the can be altered.  

1. Security list rules:  By default, the global security list has an inbound rule set for ICMP types 3 and 4, and each subnet has an inbound rule set for port 22 (SSH) for all connections.  This is representative of the OCI default rules for building instances.  If there are other rules you want to add by default, they can be inserted into the resource.tf file for either the "global_seclist" resource, or the "subnet_seclist" resource. It is STRONGLY RECOMMENDED that the defaults are not removed, but just added to.

2. Route tables: The "default" route table is not used for subnets.  There are two new route tables created - "External" for subnets with public IPs and "Internal" for subnets that are private only.  However, since the creation of a route table requires a route, both route tables are creates with a default route point to the IG that is also created by this process.  The IG is non-functional for the "Internal" route table as it applies to private subnets only, so it is ok to deploy in this way.  If you are going to create additional resources using this template that result in additional route targets, or want to change the display name of the route tables, change the "private_rt" and "public_rt" resources accordingly.

3. No DHCP options are set:  We use the default DHCP options in this tempalte.  If you want to deploy your own DHCP options set, add it as part of this template and change the entry in the "subnets" resource to point to the DHCP options OCID created.

4. DNS name:  I build the DNS names for the VCN and subnets based on my VCN naming standards.  You can change this in the "vcn" and "subnets" resource as desired.

5. The counts must match in the variables.tf:  I use the public part of the "subnet" variable to determine the number of subnets to create.  The number of ADs in the "ad" variable *MUST* match the number of entries in public.  If it is greater, the tail will be ignored.  The same is true for the subnet_masks, if used.  The number of masks also *MUST* match the number of entries in public.  

6. The entries in public, ad, and, if used, subnet_masks, all correlate to each other.  I did this instead of creating a list of maps.  Please let me know if that would be easier to understand.

I hope this is useful in the process of building the VCN infrastructure.

Enjoy.
