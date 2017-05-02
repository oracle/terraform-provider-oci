// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"log"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type InstanceResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.Instance

	// Computed fields
	public_ip  string
	private_ip string
}

func (s *InstanceResourceCrud) ID() string {
	return s.Resource.ID
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

func resourceInstanceMapToMetadata(rm map[string]interface{}) map[string]string {
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
	if hostnameLabel, ok := s.D.GetOk("hostname_label"); ok {
		opts.HostnameLabel = hostnameLabel.(string)
	}

	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceInstanceMapToMetadata(rawMetadata.(map[string]interface{}))
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

/*
 * Return the id of the first VNIC attached to this Instance.
 *
 * NOTE while the instance is still being created, calls to this function
 * can return  an error priort to the Vnic being attached.
 */
func (s *InstanceResourceCrud) getInstanceVnicId() (vnic_id string, e error) {
	compartmentID := s.Resource.CompartmentID

	opts := &baremetal.ListVnicAttachmentsOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	opts.AvailabilityDomain = s.Resource.AvailabilityDomain
	opts.InstanceID = s.Resource.ID

	var list *baremetal.ListVnicAttachments
	if list, e = s.Client.ListVnicAttachments(compartmentID, opts); e != nil {
		return "", e
	}

	if len(list.Attachments) < 1 {
		log.Printf("[DEBUG] GetInstanceVnicID - InstanceID: %q, State: %q, no vnic attachments: %q", s.Resource.ID, s.Resource.State, e)
		return "", e
	}

	return list.Attachments[0].VnicID, nil
}

/*
 * Return the public, private IP pair associated with the instance's first Vnic.
 *
 * NOTE while the instance is still being created, calls to this function
 * can return  an error priort to the Vnic being attached.
 */
func (s *InstanceResourceCrud) getInstanceIPs() (public_ip string, private_ip string, e error) {
	vnicID, e := s.getInstanceVnicId()
	if e != nil {
		return "", "", e
	}

	// Lookup Vnic by id
	vnic, e := s.Client.GetVnic(vnicID)
	if e != nil {
		return "", "", e
	}

	return vnic.PublicIPAddress, vnic.PrivateIPAddress, nil
}

func (s *InstanceResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetInstance(s.D.Id())

	if e != nil {
		return e
	}

	// Compute instance IPs through attached Vnic
	// (Not available while state==PROVISIONING)
	public_ip, private_ip, e2 := s.getInstanceIPs()
	if e2 != nil {
		log.Printf("[DEBUG] no vnic yet, skipping")
	}

	if public_ip != "" {
		s.public_ip = public_ip
	}
	if private_ip != "" {
		s.private_ip = private_ip
	}

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

	s.D.Set("public_ip", s.public_ip)
	s.D.Set("private_ip", s.private_ip)
}

func (s *InstanceResourceCrud) Delete() (e error) {
	return s.Client.TerminateInstance(s.D.Id(), nil)
}
