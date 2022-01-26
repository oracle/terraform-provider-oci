// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"
)

func OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOpsiOperationsInsightsWarehouseResourceUsageSummary,
		Schema: map[string]*schema.Schema{
			"operations_insights_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"cpu_used": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_used_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func readSingularOpsiOperationsInsightsWarehouseResourceUsageSummary(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.SummarizeOperationsInsightsWarehouseResourceUsageResponse
}

func (s *OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSourceCrud) Get() error {
	request := oci_opsi.SummarizeOperationsInsightsWarehouseResourceUsageRequest{}

	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		request.OperationsInsightsWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.SummarizeOperationsInsightsWarehouseResourceUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CpuUsed != nil {
		s.D.Set("cpu_used", *s.Res.CpuUsed)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageUsedInGBs != nil {
		s.D.Set("storage_used_in_gbs", *s.Res.StorageUsedInGBs)
	}

	return nil
}
