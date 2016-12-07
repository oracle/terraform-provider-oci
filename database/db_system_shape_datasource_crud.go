package database

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DBSystemShapeDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListDBSystemShapes
}

func (s *DBSystemShapeDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	limit := uint64(s.D.Get("limit").(int))

	opts := &baremetal.PageListOptions{}
	if val, ok := resource.GetOk("page"); ok {
		opts.Page = val.(string)
	}

	s.Res = &baremetal.ListDBSystemShapes{}

	for {
		var list *baremetal.ListDBSystemShapes
		if list, e = s.Client.ListDBSystemShapes(compartmentID, limit, opts); e != nil {
			break
		}

		s.Res.DBSystemShapes = append(s.Res.DBSystemShapes, list.DBSystemShapes...)

		if hasNextPage := setNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBSystemShapeDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]string{}
		for _, v := range s.Res.DBSystemShapes {
			res := map[string]string{
				"available_core_count": v.AvailableCoreCount,
				"name":                 v.Name,
				"shape":                v.Shape,
			}
			resources = append(resources, res)
		}
		s.D.Set("db_system_shapes", resources)
	}
	return
}
