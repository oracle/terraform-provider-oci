// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PublicIpDataSource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Read:     readPublicIpDataSource,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"private_ip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifetime": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scope": {
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
		},
	}
}

func readPublicIpDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &PublicIpDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type PublicIpDataSourceCrud struct {
	crud.BaseCrud
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.PublicIp
}

func (s *PublicIpDataSourceCrud) ID() string {
	return *s.Res.Id
}

func (s *PublicIpDataSourceCrud) Get() error {
	// public ip resource fetching strategies ordered by specificity
	if id, ok := s.D.GetOkExists("id"); ok {
		return s.getById(id.(string))
	}

	if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
		return s.getByPrivateIpId(privateIpId.(string))
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		return s.getByPublicIp(ipAddress.(string))
	}

	return errors.New("require at least an id, private_ip_id, or ip_address to get a public ip data source")
}

func (s *PublicIpDataSourceCrud) getById(id string) error {
	request := oci_core.GetPublicIpRequest{}
	request.PublicIpId = &id

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetPublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *PublicIpDataSourceCrud) getByPrivateIpId(privateIpId string) error {
	request := oci_core.GetPublicIpByPrivateIpIdRequest{}
	if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
		tmp := privateIpId.(string)
		request.PrivateIpId = &tmp
	}
	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")
	response, err := s.Client.GetPublicIpByPrivateIpId(context.Background(), request)

	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *PublicIpDataSourceCrud) getByPublicIp(ipAddress string) error {
	request := oci_core.GetPublicIpByIpAddressRequest{}
	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")
	response, err := s.Client.GetPublicIpByIpAddress(context.Background(), request)

	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *PublicIpDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	s.D.Set("lifetime", s.Res.Lifetime)

	if s.Res.PrivateIpId != nil {
		s.D.Set("private_ip_id", *s.Res.PrivateIpId)
	}

	s.D.Set("scope", s.Res.Scope)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
