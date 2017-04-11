// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type DrgAttachmentResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DrgAttachment
}

func (s *DrgAttachmentResourceCrud) ID() string {
	return s.Res.ID
}

func (s *DrgAttachmentResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceAttaching}
}

func (s *DrgAttachmentResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAttached}
}

func (s *DrgAttachmentResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDetaching}
}

func (s *DrgAttachmentResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDetached}
}

func (s *DrgAttachmentResourceCrud) State() string {
	return s.Res.State
}

func (s *DrgAttachmentResourceCrud) Create() (e error) {
	drgID := s.D.Get("drg_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.CreateOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateDrgAttachment(drgID, vcnID, opts)

	return
}

func (s *DrgAttachmentResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetDrgAttachment(s.D.Id())
	return
}

func (s *DrgAttachmentResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateDrgAttachment(compartmentID, opts)
	return
}

func (s *DrgAttachmentResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("drg_id", s.Res.DrgID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("vcn_id", s.Res.VcnID)
}

func (s *DrgAttachmentResourceCrud) Delete() (e error) {
	return s.Client.DeleteDrgAttachment(s.D.Id(), nil)
}
