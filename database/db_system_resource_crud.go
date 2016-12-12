package database

import (
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
