package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func VolumeBackupDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumeBackups,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_backups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VolumeBackupResource(),
			},
		},
	}
}

func readVolumeBackups(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeBackupDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
