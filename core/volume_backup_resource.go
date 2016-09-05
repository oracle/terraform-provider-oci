package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func VolumeBackupResource() *schema.Resource {
	return &schema.Resource{
		Create: createVolumeBackup,
		Read:   readVolumeBackup,
		Update: updateVolumeBackup,
		Delete: deleteVolumeBackup,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_mbs": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_request_received": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_size_in_mbs": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeBackupResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeBackupResourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}

func updateVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeBackupResourceCrud{D: d, Client: client}
	return crud.UpdateResource(d, sync)
}

func deleteVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeBackupResourceCrud{D: d, Client: client}
	return crud.DeleteResource(sync)
}
