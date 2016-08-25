package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VolumeDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListVolumes
}

func (s *VolumeDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(s.D, "availability_domain", "limit", "page")

	if s.Res, e = s.Client.ListVolumes(compartmentID, opts...); e != nil {
		return
	}

	return
}

func (s *VolumeDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		volumes := []map[string]string{}
		for _, v := range s.Res.Volumes {
			vol := map[string]string{
				"availability_domain": v.AvailabilityDomain,
				"compartment_id":      v.CompartmentID,
				"display_name":        v.DisplayName,
				"id":                  v.ID,
				"size_in_mbs":         v.SizeInMBs,
				"state":               v.State,
				"time_created":        v.TimeCreated.String(),
			}
			volumes = append(volumes, vol)
		}
		s.D.Set("volumes", volumes)
	}
	return
}
