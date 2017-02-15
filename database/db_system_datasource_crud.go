// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
)

type DBSystemDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBSystems
}

func (s *DBSystemDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	limit := uint64(s.D.Get("limit").(int))

	opts := &baremetal.PageListOptions{}
	options.SetPageOptions(s.D, opts)

	s.Res = &baremetal.ListDBSystems{DBSystems: []baremetal.DBSystem{}}

	for {
		var list *baremetal.ListDBSystems
		if list, e = s.Client.ListDBSystems(compartmentID, limit, opts); e != nil {
			break
		}

		s.Res.DBSystems = append(s.Res.DBSystems, list.DBSystems...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, opts); !hasNextPage {
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
