// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiChargebackPlansDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readOpsiChargebackPlansWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"chargebackplan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"chargeback_plan_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpsiChargebackPlanResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiChargebackPlansWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OpsiChargebackPlansDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OpsiChargebackPlansDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListChargebackPlansResponse
}

func (s *OpsiChargebackPlansDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiChargebackPlansDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_opsi.ListChargebackPlansRequest{}

	if chargebackplanId, ok := s.D.GetOkExists("chargebackplan_id"); ok {
		tmp := chargebackplanId.(string)
		request.ChargebackplanId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListChargebackPlans(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListChargebackPlans(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiChargebackPlansDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiChargebackPlansDataSource-", OpsiChargebackPlansDataSource(), s.D))
	resources := []map[string]interface{}{}
	chargebackPlan := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ChargebackPlanSummaryToMap(item))
	}
	chargebackPlan["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiChargebackPlansDataSource().Schema["chargeback_plan_collection"].Elem.(*schema.Resource).Schema)
		chargebackPlan["items"] = items
	}

	resources = append(resources, chargebackPlan)
	if err := s.D.Set("chargeback_plan_collection", resources); err != nil {
		return err
	}

	return nil
}
