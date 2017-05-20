// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DBSystemDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBSystems,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DBSystemResource(),
			},
		},
	}
}

func readDBSystems(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type DBSystemDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBSystems
}

func (s *DBSystemDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetPageOptions(s.D, &opts.PageListOptions)
	options.SetLimitOptions(s.D, &opts.LimitListOptions)

	s.Res = &baremetal.ListDBSystems{DBSystems: []baremetal.DBSystem{}}

	for {
		var list *baremetal.ListDBSystems
		if list, e = s.Client.ListDBSystems(compartmentID, opts); e != nil {
			break
		}

		s.Res.DBSystems = append(s.Res.DBSystems, list.DBSystems...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBSystemDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, r := range s.Res.DBSystems {
			db := map[string]interface{}{
				"admin_password": r.DBHome.Database.AdminPassword,
				"db_name":        r.DBHome.Database.DBName,
			}
			dbHome := map[string]interface{}{
				"database":     []interface{}{db},
				"db_version":   r.DBHome.DBVersion,
				"display_name": r.DBHome.DisplayName,
			}
			res := map[string]interface{}{
				"availability_domain": r.AvailabilityDomain,
				"compartment_id":      r.CompartmentID,
				"shape":               r.Shape,
				"subnet_id":           r.SubnetID,
				"ssh_public_keys":     r.SSHPublicKeys,
				"cpu_core_count":      int(r.CPUCoreCount),
				"db_home":             []interface{}{dbHome},
				"display_name":        r.DisplayName,
				"database_edition":    r.DatabaseEdition,
				"disk_redundancy":     r.DiskRedundancy,
				"domain":              r.Domain,
				"hostname":            r.Hostname,
				"id":                  r.ID,
				"lifecycle_details":   r.LifecycleDetails,
				"listener_port":       int(r.ListenerPort),
				"state":               r.State,
				"time_created":        r.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("db_systems", resources)
	}
	return
}
