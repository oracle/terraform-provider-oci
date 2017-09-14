#### About this tool

This tool will target a terraform plan directory and transform all
**baremetal** names found in *.tf* and *.tfstate* files to the new
**oci** provider name. It creates a backup of the target directory
_\<plan-directory\>.backup_ as a sibling folder.

It also detects and modifies provider blocks not specifying a region
to explicitly use "us-phoenix-1" since the default region value has 
been removed from the provider.

#### Provider Setup and Config Verfification
You can use the new and old providers side by side. This is helpful for 
knowing you are migrating plans from a known good state (no pending 
changes when running plan), and will be able to verify successful 
migration (no pending changes running plan with the new provider).

If you configure your plugins with a .terraformrc file, add an entry for the new oci provider, example:
```
providers {
	oci = "/Users/moi/providers/terraform-provider-oci"
	baremetal = "/Users/moi/providers/terraform-provider-baremetal"
}
```

Otherwise, drop the new provider in your terraform plugin directory.

#### Using the tool

To migrate a plan, follow these steps:  

From the plan directory, run `terraform plan`, make sure there are no
pending changes in your plan.

Execute the **oci-tool** binary, passing the path to your plan
directory, example:  
`oci-tool migrate -dir=<plan-path>`

After migrating a plan file, run `terraform plan` again and verify
there are no new pending modifications.

For Terraform v.10+, you will need to initialize terraform for the
directory using  
`terraform init`

If the migration was not successful, manually restore the plan files
from the .backup directory or run:  
`oci-tool backup -dir=<plan-path> -restore`

After you have verified the migration was successful, delete the
backup folder or run:  
`oci-tool backup -dir=<plan-path> -purge`
