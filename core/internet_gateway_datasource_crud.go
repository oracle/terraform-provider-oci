package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type InternetGatewayDatasourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.ListInternetGateways
}

func (s *InternetGatewayDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.ListOptions{}
	setListOptions(s.D, opts)

	s.Resource = &baremetal.ListInternetGateways{
		Gateways: []baremetal.InternetGateway{},
	}

	for {
		var list *baremetal.ListInternetGateways
		if list, e = s.Client.ListInternetGateways(compartmentID, vcnID, opts); e != nil {
			break
		}

		s.Resource.Gateways = append(s.Resource.Gateways, list.Gateways...)

		if hasNextPage := setNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s InternetGatewayDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Resource.Gateways {

			resource := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"enabled":        v.IsEnabled,
				"state":          v.State,
				"time_modified":  v.ModifiedTime.String(),
				"time_created":   v.TimeCreated.String(),
			}

			resources = append(resources, resource)
		}

		s.D.Set("gateways", resources)

	}

	return
}
