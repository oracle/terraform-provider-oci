// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreByoipRangeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreByoipRange,
		Schema: map[string]*schema.Schema{
			"byoip_range_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"cidr_block": {
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
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_advertised": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_validated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_withdrawn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"validation_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreByoipRange(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoipRangeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreByoipRangeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetByoipRangeResponse
}

func (s *CoreByoipRangeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreByoipRangeDataSourceCrud) Get() error {
	request := oci_core.GetByoipRangeRequest{}

	if byoipRangeId, ok := s.D.GetOkExists("byoip_range_id"); ok {
		tmp := byoipRangeId.(string)
		request.ByoipRangeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetByoipRange(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreByoipRangeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAdvertised != nil {
		s.D.Set("time_advertised", s.Res.TimeAdvertised.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeValidated != nil {
		s.D.Set("time_validated", s.Res.TimeValidated.String())
	}

	if s.Res.TimeWithdrawn != nil {
		s.D.Set("time_withdrawn", s.Res.TimeWithdrawn.String())
	}

	if s.Res.ValidationToken != nil {
		s.D.Set("validation_token", *s.Res.ValidationToken)
	}

	return nil
}
