// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DBVersionDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBVersions,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supports_pdb": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func readDBVersions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	reader := &DBVersionDatasourceCrud{}
	reader.D = d
	reader.Client = client
	return crud.ReadResource(reader)
}

type DBVersionDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBVersions
}

func (s *DBVersionDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetPageOptions(s.D, &opts.PageListOptions)
	options.SetLimitOptions(s.D, &opts.LimitListOptions)

	s.Res = &baremetal.ListDBVersions{}

	for {
		var list *baremetal.ListDBVersions
		if list, e = s.Client.ListDBVersions(compartmentID, opts); e != nil {
			break
		}

		s.Res.DBVersions = append(s.Res.DBVersions, list.DBVersions...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBVersionDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.DBVersions {
			res := map[string]interface{}{
				"version":      v.Version,
				"supports_pdb": v.SupportsPDB,
			}
			resources = append(resources, res)
		}
		s.D.Set("db_versions", resources)
	}
	return
}
