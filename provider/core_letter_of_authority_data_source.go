// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LetterOfAuthorityDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLetterOfAuthority,
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
			// @CODEGEN 07/2018: Remove duplicated fields in computed that are also required
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

func readSingularLetterOfAuthority(d *schema.ResourceData, m interface{}) error {
	sync := &LetterOfAuthorityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type LetterOfAuthorityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectLetterOfAuthorityResponse
}

func (s *LetterOfAuthorityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LetterOfAuthorityDataSourceCrud) Get() error {
	request := oci_core.GetCrossConnectLetterOfAuthorityRequest{}

	if crossConnectId, ok := s.D.GetOkExists("cross_connect_id"); ok {
		tmp := crossConnectId.(string)
		request.CrossConnectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnectLetterOfAuthority(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LetterOfAuthorityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(crud.GenerateDataSourceID())

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
