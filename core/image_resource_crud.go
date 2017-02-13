// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type ImageResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Image
}

func (s *ImageResourceCrud) ID() string {
	return s.Res.ID
}

func (s *ImageResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *ImageResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *ImageResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDisabled}
}

func (s *ImageResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *ImageResourceCrud) State() string {
	return s.Res.State
}

func (s *ImageResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	instanceID := s.D.Get("instance_id").(string)

	opts := &baremetal.CreateOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateImage(compartmentID, instanceID, opts)

	return
}

func (s *ImageResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetImage(s.D.Id())
	return
}

func (s *ImageResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateImage(s.D.Id(), opts)

	return
}

func (s *ImageResourceCrud) SetData() {
	s.D.Set("base_image_id", s.Res.BaseImageID)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("create_image_allowed", s.Res.CreateImageAllowed)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("operating_system", s.Res.OperatingSystem)
	s.D.Set("operating_system_version", s.Res.OperatingSystemVersion)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *ImageResourceCrud) Delete() (e error) {
	return s.Client.DeleteImage(s.D.Id(), nil)
}
