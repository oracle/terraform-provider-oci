package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type RouteTableDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListRouteTables
}

func (s *RouteTableDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)
	opts := &baremetal.ListOptions{}
	setListOptions(s.D, opts)

	s.Res = &baremetal.ListRouteTables{RouteTables: []baremetal.RouteTable{}}

	for {
		var list *baremetal.ListRouteTables
		if list, e = s.Client.ListRouteTables(compartmentID, vcnID, opts); e != nil {
			break
		}

		s.Res.RouteTables = append(s.Res.RouteTables, list.RouteTables...)

		if hasNextPage := setNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}

	return
}

func (s *RouteTableDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())

		resources := []map[string]interface{}{}
		for _, v := range s.Res.RouteTables {

			rules := []map[string]interface{}{}
			for _, val := range v.RouteRules {
				rule := map[string]interface{}{
					"cidr_block":          val.CidrBlock,
					"display_name":        val.DisplayName,
					"network_entity_id":   val.NetworkEntityID,
					"network_entity_type": val.NetworkEntityType,
					"time_created":        val.TimeCreated.String(),
				}
				rules = append(rules, rule)
			}

			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"route_rules":    rules,
				"time_modified":  v.TimeModified.String(),
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("route_tables", resources)
	}
	return
}
