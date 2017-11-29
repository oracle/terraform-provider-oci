// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"

	"log"
)

func VnicAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVnicAttachment,
		Read:     readVnicAttachment,
		Update:   updateVnicAttachment,
		Delete:   deleteVnicAttachment,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_vnic_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem:     createVnicDetailsSchema,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_tag": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVnicAttachment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readVnicAttachment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func updateVnicAttachment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.UpdateResource(sync.D, sync)
}

func deleteVnicAttachment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
	return crud.DeleteResource(d, sync)
}

type VnicAttachmentResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.VnicAttachment
}

func (s *VnicAttachmentResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *VnicAttachmentResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceAttaching}
}

func (s *VnicAttachmentResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAttached}
}

func (s *VnicAttachmentResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDetaching}
}

func (s *VnicAttachmentResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDetached}
}

func (s *VnicAttachmentResourceCrud) Create() (e error) {
	instanceID := s.D.Get("instance_id").(string)

	vaOpts := &baremetal.AttachVnicOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		vaOpts.DisplayName = displayName.(string)
	}

	vnicOpts := SetCreateVnicOptions(s.D.Get("create_vnic_details").([]interface{}))

	s.Resource, e = s.Client.AttachVnic(instanceID, vnicOpts, vaOpts)
	return
}

func (s *VnicAttachmentResourceCrud) Update() (e error) {
	// The VNIC ID is also available at s.D.Get("vnic_id"). However,
	// the VnicAttachment resources must be fetched anyway in order to update
	// the state data after the update call.
	s.Resource, e = s.Client.GetVnicAttachment(s.D.Id())
	if e != nil {
		return
	}

	opts := SetUpdateVnicOptions(s.D.Get("create_vnic_details").([]interface{}))
	_, e = s.Client.UpdateVnic(s.Resource.VnicID, opts)
	return
}

func (s *VnicAttachmentResourceCrud) Get() (e error) {
	res, e := s.Client.GetVnicAttachment(s.D.Id())
	if e == nil {
		s.Resource = res
	}
	return
}

func (s *VnicAttachmentResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Resource.AvailabilityDomain)
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("instance_id", s.Resource.InstanceID)
	s.D.Set("state", s.Resource.State)
	s.D.Set("subnet_id", s.Resource.SubnetID)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
	s.D.Set("vlan_tag", s.Resource.VlanTag)
	s.D.Set("vnic_id", s.Resource.VnicID)

	vnic, err := s.Client.GetVnic(s.Resource.VnicID)
	if vnic == nil {
		// VNIC might not be found when attaching or detaching.
		log.Printf("[DEBUG] VNIC not found during VNIC Attachment refresh. (VNIC ID: %q, Error: %q)", s.Resource.VnicID, err)
		return
	}

	RefreshCreateVnicDetails(s.D, vnic)
}

func (s *VnicAttachmentResourceCrud) Delete() (e error) {
	return s.Client.DetachVnic(s.D.Id(), nil)
}
