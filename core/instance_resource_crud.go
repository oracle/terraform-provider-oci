// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

type InstanceResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.Instance
}

func (s *InstanceResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *InstanceResourceCrud) CustomTimeout() time.Duration {
	return 15 * time.Minute
}

func (s *InstanceResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceProvisioning,
		baremetal.ResourceStarting,
	}
}

func (s *InstanceResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceRunning}
}

func (s *InstanceResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *InstanceResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func resourceMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func (s *InstanceResourceCrud) Create() (e error) {
	availabilityDomain := s.D.Get("availability_domain").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	image := s.D.Get("image").(string)
	shape := s.D.Get("shape").(string)
	subnet := s.D.Get("subnet_id").(string)

	opts := &baremetal.LaunchInstanceOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}

	s.Resource, e = s.Client.LaunchInstance(
		availabilityDomain,
		compartmentID,
		image,
		shape,
		subnet,
		opts)
	return
}

func (s *InstanceResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetInstance(s.D.Id())
	return
}

func (s *InstanceResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateInstance(s.D.Id(), opts)
	return
}

func (s *InstanceResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Resource.AvailabilityDomain)
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("image", s.Resource.ImageID)
	s.D.Set("metadata", s.Resource.Metadata)
	s.D.Set("region", s.Resource.Region)
	s.D.Set("shape", s.Resource.Shape)
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
}

func (s *InstanceResourceCrud) Delete() (e error) {
	return s.Client.TerminateInstance(s.D.Id(), nil)
}
