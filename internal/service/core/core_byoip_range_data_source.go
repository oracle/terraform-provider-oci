// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
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
			"byoip_range_vcn_ipv6allocations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"byoip_range_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
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
			"ipv6cidr_block": {
				Type:     schema.TypeString,
				Computed: true,
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

	byoipRangeVcnIpv6Allocations := []interface{}{}
	for _, item := range s.Res.ByoipRangeVcnIpv6Allocations {
		byoipRangeVcnIpv6Allocations = append(byoipRangeVcnIpv6Allocations, ByoipRangeVcnIpv6AllocationSummaryToMap(item))
	}
	s.D.Set("byoip_range_vcn_ipv6allocations", byoipRangeVcnIpv6Allocations)

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

	if s.Res.Ipv6CidrBlock != nil {
		s.D.Set("ipv6cidr_block", *s.Res.Ipv6CidrBlock)
	}

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

func ByoipRangeVcnIpv6AllocationSummaryToMap(obj oci_core.ByoipRangeVcnIpv6AllocationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ByoipRangeId != nil {
		result["byoip_range_id"] = string(*obj.ByoipRangeId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Ipv6CidrBlock != nil {
		result["ipv6cidr_block"] = string(*obj.Ipv6CidrBlock)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}
