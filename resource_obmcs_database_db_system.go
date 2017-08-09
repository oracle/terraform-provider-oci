// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DBSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &crud.ZeroTime,
			Delete: &crud.TwoHours,
			Update: &crud.TwoHours,
		},
		Create: createDBSystem,
		Read:   readDBSystem,
		Delete: deleteDBSystem,
		Schema: map[string]*schema.Schema{
			//Required
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
			"database_edition": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"db_home": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_password": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
										ForceNew:  true,
									},
									"db_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"db_workload": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"character_set": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ncharacter_set": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"pdb_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"db_version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					},
				},
			},
			"hostname": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			//Optional
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
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
			"backup_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cluster_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"data_storage_percentage": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},

			//Computed
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
			"scan_dns_record_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_ip_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vip_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func createDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateDBSystemResource(d, sync)
}

func readDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func deleteDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
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
	cpuCoreCount := uint64(s.D.Get("cpu_core_count").(int))
	databaseEdition := baremetal.DatabaseEdition(s.D.Get("database_edition").(string))
	hostname := s.D.Get("hostname").(string)
	shape := s.D.Get("shape").(string)
	sshPublicKeys := []string{}
	for _, key := range s.D.Get("ssh_public_keys").([]interface{}) {
		sshPublicKeys = append(sshPublicKeys, key.(string))
	}
	subnetID := s.D.Get("subnet_id").(string)
	rawDBHome := s.D.Get("db_home")
	var dbHomeDetails baremetal.CreateDBHomeDetails
	l := rawDBHome.([]interface{})
	if len(l) > 0 {
		dbHome := l[0].(map[string]interface{})
		db := dbHome["database"].([]interface{})[0].(map[string]interface{})
		dbVersion := dbHome["db_version"].(string)
		displayName := dbHome["display_name"]
		adminPassword := db["admin_password"].(string)
		dbName := db["db_name"].(string)
		dbWorkload := db["db_workload"]
		characterSet := db["character_set"]
		ncharacterSet := db["ncharacter_set"]
		pdbName := db["pdb_name"]

		dbHomeOpts := &baremetal.CreateDBHomeOptions{}
		if displayName != nil {
			dbHomeOpts.DisplayName = displayName.(string)
		}
		dbOpts := &baremetal.CreateDatabaseOptions{}
		if dbWorkload != nil {
			dbOpts.DBWorkload = dbWorkload.(string)
		}
		if characterSet != nil {
			dbOpts.CharacterSet = characterSet.(string)
		}
		if ncharacterSet != nil {
			dbOpts.NCharacterSet = ncharacterSet.(string)
		}
		if pdbName != nil {
			dbOpts.PDBName = pdbName.(string)
		}

		dbHomeDetails = baremetal.NewCreateDBHomeDetails(
			baremetal.NewCreateDatabaseDetails(adminPassword, dbName, dbOpts),
			dbVersion,
			dbHomeOpts,
		)
	}

	opts := &baremetal.LaunchDBSystemOptions{}
	if backupSubnetId, ok := s.D.GetOk("backup_subnet_id"); ok {
		opts.BackupSubnetId = backupSubnetId.(string)
	}
	if clusterName, ok := s.D.GetOk("cluster_name"); ok {
		opts.ClusterName = clusterName.(string)
	}
	if dataStoragePercentage, ok := s.D.GetOk("data_storage_percentage"); ok {
		opts.DataStoragePercentage = dataStoragePercentage.(int)
	}
	if diskRedundancy, ok := s.D.GetOk("disk_redundancy"); ok {
		opts.DiskRedundancy = baremetal.DiskRedundancy(diskRedundancy.(string))
	}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}
	if domain, ok := s.D.GetOk("domain"); ok {
		opts.Domain = domain.(string)
	}

	s.Res, e = s.Client.LaunchDBSystem(
		availabilityDomain, compartmentID, cpuCoreCount, databaseEdition, dbHomeDetails,
		hostname, shape, sshPublicKeys, subnetID,
		opts,
	)

	return
}

func (s *DBSystemResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetDBSystem(s.D.Id())
	return
}

func (s *DBSystemResourceCrud) SetData() {
	//Required
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("cpu_core_count", s.Res.CPUCoreCount)
	s.D.Set("database_edition", s.Res.DatabaseEdition)
	s.D.Set("db_home", s.Res.DBHome)
	//leave hostname commented out. Refreshing hostname causes problems because API adds suffix in some cases (like Exadata).
	//s.D.Set("hostname", s.Res.Hostname)
	s.D.Set("shape", s.Res.Shape)
	s.D.Set("ssh_public_keys", s.Res.SSHPublicKeys)
	s.D.Set("subnet_id", s.Res.SubnetID)

	//Optional
	s.D.Set("backup_subnet_id", s.Res.BackupSubnetID)
	s.D.Set("cluster_name", s.Res.ClusterName)
	s.D.Set("data_storage_percentage", s.Res.DataStoragePercentage)
	s.D.Set("disk_redundancy", s.Res.DiskRedundancy)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("domain", s.Res.Domain)

	//Computed
	s.D.Set("id", s.Res.ID)
	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)
	s.D.Set("listener_port", s.Res.ListenerPort)
	s.D.Set("scan_dns_record_id", s.Res.ScanDnsRecordId)
	s.D.Set("scan_ip_ids", s.Res.ScanIpIds)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("version", s.Res.Version)
	s.D.Set("vip_ids", s.Res.VipIds)
}

func (s *DBSystemResourceCrud) Delete() (e error) {
	return s.Client.TerminateDBSystem(s.D.Id(), nil)
}
