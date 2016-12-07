package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type IPSecConnectionsDatasourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.ListIPSecConnections
}

func (s *IPSecConnectionsDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListIPSecConnsOptions{}
	setListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("cpe_id"); ok {
		opts.CpeID = val.(string)
	}
	if val, ok := s.D.GetOk("drg_id"); ok {
		opts.DrgID = val.(string)
	}

	s.Resource = &baremetal.ListIPSecConnections{
		Connections: []baremetal.IPSecConnection{},
	}

	for {
		var list *baremetal.ListIPSecConnections
		if list, e = s.Client.ListIPSecConnections(compartmentID, opts); e != nil {
			break
		}

		s.Resource.Connections = append(s.Resource.Connections, list.Connections...)

		if hasNextPage := setNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s IPSecConnectionsDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Resource.Connections {

			resource := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"drg_id":         v.DrgID,
				"cpe_id":         v.CpeID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"static_routes":  v.StaticRoutes,
				"time_created":   v.TimeCreated.String(),
			}

			resources = append(resources, resource)
		}

		s.D.Set("connections", resources)

	}

	return
}
