// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OnesubscriptionCommitmentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOnesubscriptionCommitment,
		Schema: map[string]*schema.Schema{
			"commitment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"subscribed_service_id": {
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

func readSingularOnesubscriptionCommitment(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionCommitmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CommitmentRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionCommitmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.CommitmentClient
	Res    *oci_onesubscription.GetCommitmentResponse
}

func (s *OnesubscriptionCommitmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionCommitmentDataSourceCrud) Get() error {
	request := oci_onesubscription.GetCommitmentRequest{}

	if commitmentId, ok := s.D.GetOkExists("commitment_id"); ok {
		tmp := commitmentId.(string)
		request.CommitmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

	response, err := s.Client.GetCommitment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OnesubscriptionCommitmentDataSourceCrud) SetData() error {
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

	if s.Res.SubscribedServiceId != nil {
		s.D.Set("subscribed_service_id", *s.Res.SubscribedServiceId)
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
