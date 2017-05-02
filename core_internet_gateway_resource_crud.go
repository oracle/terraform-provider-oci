// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type InternetGatewayResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.InternetGateway
}

func (s *InternetGatewayResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *InternetGatewayResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceProvisioning,
	}
}

func (s *InternetGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceAvailable,
	}
}

func (s *InternetGatewayResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceTerminating,
	}
}

func (s *InternetGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		baremetal.ResourceTerminated,
	}
}

func (s *InternetGatewayResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)
	isEnabled := s.D.Get("enabled").(bool)

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.D.Get("display_name").(string)

	s.Resource, e = s.Client.CreateInternetGateway(compartmentID, vcnID, isEnabled, opts)
	return
}

func (s *InternetGatewayResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetInternetGateway(s.D.Id())
	return
}

func (s *InternetGatewayResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateGatewayOptions{}
	if isEnabled, ok := s.D.GetOk("enabled"); ok {
		opts.IsEnabled = isEnabled.(bool)
	}

	s.Resource, e = s.Client.UpdateInternetGateway(s.D.Id(), opts)
	return
}

func (s *InternetGatewayResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("enabled", s.Resource.IsEnabled)
	s.D.Set("time_modified", s.Resource.ModifiedTime.String())
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
}

func (s *InternetGatewayResourceCrud) Delete() (e error) {
	return s.Client.DeleteInternetGateway(s.D.Id(), nil)
}
