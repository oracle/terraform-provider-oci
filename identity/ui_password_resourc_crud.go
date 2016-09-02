package identity

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type UIPasswordResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.UIPassword
}

func (s *UIPasswordResourceCrud) ID() string {
	return s.D.Id()
}

func (s *UIPasswordResourceCrud) State() string {
	return s.Res.State
}

func (s *UIPasswordResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *UIPasswordResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *UIPasswordResourceCrud) Create() (e error) {
	userID := s.D.Get("user_id").(string)
	s.Res, e = s.Client.CreateOrResetUIPassword(userID)
	return
}

func (s *UIPasswordResourceCrud) SetData() {
	s.D.Set("inactive_status", s.Res.InactiveStatus)
	s.D.Set("state", s.Res.State)
	s.D.Set("password", s.Res.Password)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("user_id", s.Res.UserID)
}
