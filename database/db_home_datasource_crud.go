package database

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DBHomeDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.DBHome
}

func (s *DBHomeDatasourceCrud) Get() (e error) {
	id := s.D.Get("db_home_id").(string)
	s.Res, e = s.Client.GetDBHome(id)
	return
}

func (s *DBHomeDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(s.Res.ID)
		s.D.Set("compartment_id", s.Res.CompartmentID)
		s.D.Set("db_system_id", s.Res.DBSystemID)
		s.D.Set("db_version", s.Res.DBVersion)
		s.D.Set("display_name", s.Res.DisplayName)
		s.D.Set("id", s.Res.ID)
		s.D.Set("state", s.Res.State)
		s.D.Set("time_created", s.Res.TimeCreated)
	}
	return
}
