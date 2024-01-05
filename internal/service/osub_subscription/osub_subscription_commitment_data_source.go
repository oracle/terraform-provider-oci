// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_subscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osub_subscription "github.com/oracle/oci-go-sdk/v65/osubsubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsubSubscriptionCommitmentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsubSubscriptionCommitment,
		Schema: map[string]*schema.Schema{
			"commitment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"x_one_gateway_subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"x_one_origin_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"available_amount": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"funded_allocation_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"quantity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"used_amount": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOsubSubscriptionCommitment(d *schema.ResourceData, m interface{}) error {
	sync := &OsubSubscriptionCommitmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CommitmentClient()

	return tfresource.ReadResource(sync)
}

type OsubSubscriptionCommitmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_subscription.CommitmentClient
	Res    *oci_osub_subscription.GetCommitmentResponse
}

func (s *OsubSubscriptionCommitmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubSubscriptionCommitmentDataSourceCrud) Get() error {
	request := oci_osub_subscription.GetCommitmentRequest{}

	if commitmentId, ok := s.D.GetOkExists("commitment_id"); ok {
		tmp := commitmentId.(string)
		request.CommitmentId = &tmp
	}

	if xOneGatewaySubscriptionId, ok := s.D.GetOkExists("x_one_gateway_subscription_id"); ok {
		tmp := xOneGatewaySubscriptionId.(string)
		request.XOneGatewaySubscriptionId = &tmp
	}

	if xOneOriginRegion, ok := s.D.GetOkExists("x_one_origin_region"); ok {
		tmp := xOneOriginRegion.(string)
		request.XOneOriginRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osub_subscription")

	response, err := s.Client.GetCommitment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsubSubscriptionCommitmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailableAmount != nil {
		s.D.Set("available_amount", *s.Res.AvailableAmount)
	}

	if s.Res.FundedAllocationValue != nil {
		s.D.Set("funded_allocation_value", *s.Res.FundedAllocationValue)
	}

	if s.Res.Quantity != nil {
		s.D.Set("quantity", *s.Res.Quantity)
	}

	if s.Res.TimeEnd != nil {
		s.D.Set("time_end", s.Res.TimeEnd.String())
	}

	if s.Res.TimeStart != nil {
		s.D.Set("time_start", s.Res.TimeStart.String())
	}

	if s.Res.UsedAmount != nil {
		s.D.Set("used_amount", *s.Res.UsedAmount)
	}

	return nil
}
