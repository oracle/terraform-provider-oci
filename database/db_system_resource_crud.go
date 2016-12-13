package database

import (
	"fmt"
	"reflect"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DBSystemResourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.DBSystem
}

func (s *DBSystemResourceCrud) ID() string {
	return s.Resource.ID
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

	if dbHome, ok := s.D.GetOk("db_home"); ok {
		fmt.Println(dbHome)
		fmt.Println(reflect.TypeOf(dbHome))
		// opts.DisplayName = dbHome.()
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

	s.Resource, e = s.Client.LaunchDBSystem(
		availabilityDomain, compartmentID, shape, subnetID,
		sshPublicKeys, cpuCoreCount, opts,
	)

	return
}

func (s *DBSystemResourceCrud) Get() (e error) {
	return
}

func (s *DBSystemResourceCrud) SetData() {
}

func (s *DBSystemResourceCrud) Delete() (e error) {
	return
}
