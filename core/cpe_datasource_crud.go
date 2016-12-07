package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type CPEDatasourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.ListCpes
}

func (s *CPEDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListOptions{}
	setListOptions(s.D, opts)

	s.Resource = &baremetal.ListCpes{Cpes: []baremetal.Cpe{}}

	for {
		var list *baremetal.ListCpes
		if list, e = s.Client.ListCpes(compartmentID, opts); e != nil {
			break
		}

		s.Resource.Cpes = append(s.Resource.Cpes, list.Cpes...)

		if hasNextPage := setNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *CPEDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(time.Now().UTC().String())

		cpes := []map[string]interface{}{}

		for _, v := range s.Resource.Cpes {
			cpe := map[string]interface{}{
				"id":             v.ID,
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"ip_address":     v.IPAddress,
				"time_created":   v.TimeCreated.String(),
			}

			cpes = append(cpes, cpe)
		}

		s.D.Set("cpes", cpes)

	}
}
