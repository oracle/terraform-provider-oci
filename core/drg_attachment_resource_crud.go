package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DrgAttachmentSync struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.DrgAttachment
}

func (s *DrgAttachmentSync) ID() string {
	return s.Res.ID
}

func (s *DrgAttachmentSync) CreatedPending() []string {
	return []string{baremetal.ResourceAttaching}
}

func (s *DrgAttachmentSync) CreatedTarget() []string {
	return []string{baremetal.ResourceAttached}
}

func (s *DrgAttachmentSync) DeletedPending() []string {
	return []string{baremetal.ResourceDetaching}
}

func (s *DrgAttachmentSync) DeletedTarget() []string {
	return []string{baremetal.ResourceDetached}
}

func (s *DrgAttachmentSync) State() string {
	return s.Res.State
}

func (s *DrgAttachmentSync) Create() (e error) {
	opts := baremetal.Options{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}
	drgID := s.D.Get("drg_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	s.Res, e = s.Client.CreateDrgAttachment(compartmentID, drgID, vcnID, opts)

	return
}

func (s *DrgAttachmentSync) Get() (e error) {
	s.Res, e = s.Client.GetDrgAttachment(s.D.Id())
	return
}

func (s *DrgAttachmentSync) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("drg_id", s.Res.DrgID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("vcn_id", s.Res.VcnID)
}

func (s *DrgAttachmentSync) Delete() (e error) {
	return s.Client.DeleteDrgAttachment(s.D.Id())
}
