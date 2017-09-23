// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PrivateIPResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPrivateIP,
		Read:     readPrivateIP,
		Update:   updatePrivateIP,
		Delete:   deletePrivateIP,
		Schema: map[string]*schema.Schema{
			//Required
			"vnic_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			//Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			//Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_primary": {
				Type:     schema.TypeBool,
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
		},
	}
}

func createPrivateIP(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PrivateIPResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.CreateResource(d, sync)
}

func readPrivateIP(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PrivateIPResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.ReadResource(sync)
}

func updatePrivateIP(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PrivateIPResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.UpdateResource(d, sync)
}

func deletePrivateIP(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PrivateIPResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.DeleteResource(d, sync)
}

type PrivateIPResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.PrivateIP
}

func (s *PrivateIPResourceCrud) ID() string {
	return s.Res.ID
}

func (s *PrivateIPResourceCrud) Create() (e error) {
	vnicID := s.D.Get("vnic_id").(string)

	opts := &baremetal.CreatePrivateIPOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}
	hostnameLabel, ok := s.D.GetOk("hostname_label")
	if ok {
		opts.HostnameLabel = hostnameLabel.(string)
	}
	ipAddress, ok := s.D.GetOk("ip_address")
	if ok {
		opts.IPAddress = ipAddress.(string)
	}

	s.Res, e = s.Client.CreatePrivateIP(vnicID, opts)
	return
}

func (s *PrivateIPResourceCrud) Get() (e error) {
	res, e := s.Client.GetPrivateIP(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *PrivateIPResourceCrud) Update() (e error) {
	opts := &baremetal.UpdatePrivateIPOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}
	if hostnameLabel, ok := s.D.GetOk("hostname_label"); ok {
		opts.HostnameLabel = hostnameLabel.(string)
	}
	if vnicID, ok := s.D.GetOk("vnic_id"); ok {
		opts.VnicID = vnicID.(string)
	}
	s.Res, e = s.Client.UpdatePrivateIP(s.D.Id(), opts)
	return
}

func (s *PrivateIPResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("hostname_label", s.Res.HostnameLabel)
	s.D.Set("id", s.Res.ID)
	s.D.Set("ip_address", s.Res.IPAddress)
	s.D.Set("is_primary", s.Res.IsPrimary)
	s.D.Set("subnet_id", s.Res.SubnetID)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("vnic_id", s.Res.VnicID)
}

func (s *PrivateIPResourceCrud) Delete() (e error) {
	return s.Client.DeletePrivateIP(s.D.Id(), nil)
}
