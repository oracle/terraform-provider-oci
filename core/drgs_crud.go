package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DrgsSync struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.DrgList
}

func (s *DrgsSync) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(
		s.D,
		"limit",
		"page",
	)

	if s.Res, e = s.Client.ListDrgs(compartmentID, opts...); e != nil {
		return
	}

	return
}

func (s *DrgsSync) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]string{}
		for _, v := range s.Res.Drgs {
			res := map[string]string{
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("drgs", resources)
	}
	return
}
