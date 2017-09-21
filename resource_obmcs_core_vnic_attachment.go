// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VnicAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVnicAttachment,
		Read:     readVnicAttachment,
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
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_public_ip": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"hostname_label": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"skip_source_dest_check": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
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
	sync.Client = m.(*baremetal.Client)
	return crud.CreateResource(d, sync)
}

func readVnicAttachment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.ReadResource(sync)
}

func deleteVnicAttachment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
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

	vnicOpts := SetCreateVnicOptions(s.D.Get("create_vnic_details"))

	s.Resource, e = s.Client.AttachVnic(instanceID, vnicOpts, vaOpts)

	return
}

func (s *VnicAttachmentResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetVnicAttachment(s.D.Id())
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
}

func (s *VnicAttachmentResourceCrud) Delete() (e error) {
	return s.Client.DetachVnic(s.D.Id(), nil)
}
