// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DBSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &crud.TwoHours,
			Delete: &crud.TwoHours,
			Update: &crud.TwoHours,
		},
		Create: createDBSystem,
		Read:   readDBSystem,
		Delete: deleteDBSystem,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_public_keys": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},
			"database_edition": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			"db_home": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_password": {
										Type:     schema.TypeString,
										Required: true,
									},
									"db_name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"db_version": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"disk_redundancy": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"listener_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func deleteDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}

type DBSystemResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DBSystem
}

func (s *DBSystemResourceCrud) ID() string {
	return s.Res.ID
}

func (s *DBSystemResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *DBSystemResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *DBSystemResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *DBSystemResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *DBSystemResourceCrud) State() string {
	return s.Res.State
}

func (s *DBSystemResourceCrud) Create() (e error) {
	availabilityDomain := s.D.Get("availability_domain").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	shape := s.D.Get("shape").(string)
	subnetID := s.D.Get("subnet_id").(string)
	sshPublicKeys := []string{}
	for _, key := range s.D.Get("ssh_public_keys").([]interface{}) {
		sshPublicKeys = append(sshPublicKeys, key.(string))
	}
	cpuCoreCount := uint64(s.D.Get("cpu_core_count").(int))

	opts := &baremetal.LaunchDBSystemOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}
	if databaseEdition, ok := s.D.GetOk("database_edition"); ok {
		opts.DatabaseEdition = baremetal.DatabaseEdition(databaseEdition.(string))
	}

	if rawDBHome, ok := s.D.GetOk("db_home"); ok {
		l := rawDBHome.([]interface{})
		if len(l) > 0 {
			dbHome := l[0].(map[string]interface{})
			db := dbHome["database"].([]interface{})[0].(map[string]interface{})
			displayName := dbHome["display_name"]
			adminPassword := db["admin_password"].(string)
			dbName := db["db_name"].(string)
			dbVersion := dbHome["db_version"].(string)

			dbHomeOpts := &baremetal.DisplayNameOptions{}
			if displayName != nil {
				dbHomeOpts.DisplayName = displayName.(string)
			}

			dbHomeDetails := baremetal.NewCreateDBHomeDetails(
				adminPassword, dbName, dbVersion, dbHomeOpts,
			)
			opts.DBHome = dbHomeDetails
		}
	}

	if diskRedundancy, ok := s.D.GetOk("disk_redundancy"); ok {
		opts.DiskRedundancy = baremetal.DiskRedundancy(diskRedundancy.(string))
	}
	if domain, ok := s.D.GetOk("domain"); ok {
		opts.Domain = domain.(string)
	}
	if hostname, ok := s.D.GetOk("hostname"); ok {
		opts.Hostname = hostname.(string)
	}

	s.Res, e = s.Client.LaunchDBSystem(
		availabilityDomain, compartmentID, shape, subnetID,
		sshPublicKeys, cpuCoreCount, opts,
	)

	return
}

func (s *DBSystemResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetDBSystem(s.D.Id())
	return
}

func (s *DBSystemResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("shape", s.Res.Shape)
	s.D.Set("subnet_id", s.Res.SubnetID)
	s.D.Set("ssh_public_keys", s.Res.SSHPublicKeys)
	s.D.Set("cpu_core_count", s.Res.CPUCoreCount)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("database_edition", s.Res.DatabaseEdition)

	db := map[string]interface{}{
		"admin_password": s.Res.DBHome.Database.AdminPassword,
		"db_name":        s.Res.DBHome.Database.DBName,
	}
	dbHome := map[string]interface{}{
		"database":     []interface{}{db},
		"db_version":   s.Res.DBHome.DBVersion,
		"display_name": s.Res.DBHome.DisplayName,
	}
	s.D.Set("db_home", []interface{}{dbHome})

	s.D.Set("disk_redundancy", s.Res.DiskRedundancy)
	s.D.Set("domain", s.Res.Domain)
	s.D.Set("hostname", s.Res.Hostname)
	s.D.Set("id", s.Res.ID)
	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)
	s.D.Set("listener_port", s.Res.ListenerPort)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *DBSystemResourceCrud) Delete() (e error) {
	return s.Client.TerminateDBSystem(s.D.Id(), nil)
}
