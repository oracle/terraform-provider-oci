// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
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
	client := m.(*OracleClients)
	sync := &DBSystemDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
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
				"character_set":  r.DBHome.Database.CharacterSet,
				"ncharacter_set": r.DBHome.Database.NCharacterSet,
				"pdb_name":       r.DBHome.Database.PDBName,
				"db_workload":    r.DBHome.Database.DBWorkload,
			}
			dbHome := map[string]interface{}{
				"database":     []interface{}{db},
				"db_version":   r.DBHome.DBVersion,
				"display_name": r.DBHome.DisplayName,
			}
			res := map[string]interface{}{
				"availability_domain":     r.AvailabilityDomain,
				"backup_subnet_id":        r.BackupSubnetID,
				"cluster_name":            r.ClusterName,
				"compartment_id":          r.CompartmentID,
				"cpu_core_count":          int(r.CPUCoreCount),
				"data_storage_percentage": r.DataStoragePercentage,
				"data_storage_size_in_gb": int(r.DataStorageSizeInGBs),
				"database_edition":        r.DatabaseEdition,
				"db_home":                 []interface{}{dbHome},
				"display_name":            r.DisplayName,
				"disk_redundancy":         r.DiskRedundancy,
				"domain":                  r.Domain,
				"hostname":                r.Hostname,
				"id":                      r.ID,
				"license_model":           r.LicenseModel,
				"lifecycle_details":       r.LifecycleDetails,
				"listener_port":           int(r.ListenerPort),
				"node_count":              int(r.NodeCount),
				"reco_storage_size_in_gb": int(r.RecoStorageSizeInGB),
				"scan_dns_record_id":      r.ScanDnsRecordId,
				"scan_ip_ids":             r.ScanIpIds,
				"shape":                   r.Shape,
				"ssh_public_keys":         r.SSHPublicKeys,
				"state":                   r.State,
				"subnet_id":               r.SubnetID,
				"time_created":            r.TimeCreated.String(),
				"version":                 r.Version,
				"vip_ids":                 r.VipIds,
			}
			resources = append(resources, res)
		}
		s.D.Set("db_systems", resources)
	}
	return
}
