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

func OnesubscriptionCommitmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionCommitments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscribed_service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"commitments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"available_amount": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"funded_allocation_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
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
				},
			},
		},
	}
}

func readOnesubscriptionCommitments(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionCommitmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CommitmentRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionCommitmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.CommitmentClient
	Res    *oci_onesubscription.ListCommitmentsResponse
}

func (s *OnesubscriptionCommitmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionCommitmentsDataSourceCrud) Get() error {
	request := oci_onesubscription.ListCommitmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscribedServiceId, ok := s.D.GetOkExists("subscribed_service_id"); ok {
		tmp := subscribedServiceId.(string)
		request.SubscribedServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

	response, err := s.Client.ListCommitments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCommitments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OnesubscriptionCommitmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionCommitmentsDataSource-", OnesubscriptionCommitmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		commitment := map[string]interface{}{
			"subscribed_service_id": *r.SubscribedServiceId,
		}

		if r.AvailableAmount != nil {
			commitment["available_amount"] = *r.AvailableAmount
		}

		if r.FundedAllocationValue != nil {
			commitment["funded_allocation_value"] = *r.FundedAllocationValue
		}

		if r.Id != nil {
			commitment["id"] = *r.Id
		}

		if r.Quantity != nil {
			commitment["quantity"] = *r.Quantity
		}

		if r.TimeEnd != nil {
			commitment["time_end"] = r.TimeEnd.String()
		}

		if r.TimeStart != nil {
			commitment["time_start"] = r.TimeStart.String()
		}

		if r.UsedAmount != nil {
			commitment["used_amount"] = *r.UsedAmount
		}

		resources = append(resources, commitment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionCommitmentsDataSource().Schema["commitments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("commitments", resources); err != nil {
		return err
	}

	return nil
}
