// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreLetterOfAuthorityDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreLetterOfAuthority,
		Schema: map[string]*schema.Schema{
			"cross_connect_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"authorized_entity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"facility_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_issued": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreLetterOfAuthority(d *schema.ResourceData, m interface{}) error {
	sync := &CoreLetterOfAuthorityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreLetterOfAuthorityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectLetterOfAuthorityResponse
}

func (s *CoreLetterOfAuthorityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreLetterOfAuthorityDataSourceCrud) Get() error {
	request := oci_core.GetCrossConnectLetterOfAuthorityRequest{}

	if crossConnectId, ok := s.D.GetOkExists("cross_connect_id"); ok {
		tmp := crossConnectId.(string)
		request.CrossConnectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnectLetterOfAuthority(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreLetterOfAuthorityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreLetterOfAuthorityDataSource-", CoreLetterOfAuthorityDataSource(), s.D))

	if s.Res.AuthorizedEntityName != nil {
		s.D.Set("authorized_entity_name", *s.Res.AuthorizedEntityName)
	}

	s.D.Set("circuit_type", s.Res.CircuitType)

	if s.Res.FacilityLocation != nil {
		s.D.Set("facility_location", *s.Res.FacilityLocation)
	}

	if s.Res.PortName != nil {
		s.D.Set("port_name", *s.Res.PortName)
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.String())
	}

	if s.Res.TimeIssued != nil {
		s.D.Set("time_issued", s.Res.TimeIssued.String())
	}

	return nil
}
