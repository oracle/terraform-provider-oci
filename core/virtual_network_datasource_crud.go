package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VirtualNetworkDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.VirtualNetworkList
}

func (s *VirtualNetworkDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(s.D, "limit", "page")

	if s.Res, e = s.Client.ListVirtualNetworks(compartmentID, opts...); e != nil {
		return
	}

	return
}

func (s *VirtualNetworkDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]string{}
		for _, v := range s.Res.VirtualNetworks {
			res := map[string]string{
				"cidr_block":               v.CidrBlock,
				"compartment_id":           v.CompartmentID,
				"default_routing_table_id": v.DefaultRoutingTableID,
				"default_security_list_id": v.DefaultSecurityListID,
				"display_name":             v.DisplayName,
				"id":                       v.ID,
				"state":                    v.State,
				"time_created":             v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("virtual_networks", resources)
	}
	return
}
