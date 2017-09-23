// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ImageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createImage,
		Read:     readImage,
		Update:   updateImage,
		Delete:   deleteImage,
		Schema: map[string]*schema.Schema{
			"base_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_image_allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
			"operating_system": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}

type ImageResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Image
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
	res, e := s.Client.GetImage(s.D.Id())
	if e == nil {
		s.Res = res
	}
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
