// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

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
