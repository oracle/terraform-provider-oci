// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DrgAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDrgAttachment,
		Read:     readDrgAttachment,
		Update:   updateDrgAttachment,
		Delete:   deleteDrgAttachment,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDrgAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readDrgAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateDrgAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(sync.D, sync)
}

func deleteDrgAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}

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
	res, e := s.Client.GetDrgAttachment(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *DrgAttachmentResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateDrgAttachment(s.D.Id(), opts)
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
