package identity

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type APIKeyResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.APIKey
}

func (s *APIKeyResourceCrud) ID() string {
	return s.Res.KeyID
}

func (s *APIKeyResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *APIKeyResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *APIKeyResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *APIKeyResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *APIKeyResourceCrud) State() string {
	return s.Res.State
}

func (s *APIKeyResourceCrud) Create() (e error) {
	userID := s.D.Get("user_id").(string)
	key := s.D.Get("key_value").(string)

	s.Res, e = s.Client.UploadAPIKey(userID, key)

	return
}

func (s *APIKeyResourceCrud) SetData() {
	s.D.Set("key_value", s.Res.KeyValue)
	s.D.Set("fingerprint", s.Res.Fingerprint)
	s.D.Set("user_id", s.Res.UserID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *APIKeyResourceCrud) Delete() (e error) {
	userID := s.D.Get("user_id").(string)
	fingerprint := s.D.Get("fingerprint").(string)
	return s.Client.DeleteAPIKey(userID, fingerprint)
}
