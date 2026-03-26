// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package demand_signal

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DemandSignalOccMetricAlarmsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDemandSignalOccMetricAlarmsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"occ_metric_alarm_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DemandSignalOccMetricAlarmResource()),
						},
					},
				},
			},
		},
	}
}

func readDemandSignalOccMetricAlarmsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DemandSignalOccMetricAlarmsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccMetricAlarmClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DemandSignalOccMetricAlarmsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_demand_signal.OccMetricAlarmClient
	Res    *oci_demand_signal.ListOccMetricAlarmsResponse
}

func (s *DemandSignalOccMetricAlarmsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DemandSignalOccMetricAlarmsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_demand_signal.ListOccMetricAlarmsRequest{}

	compartmentId := s.D.Get("compartment_id").(string)
	if compartmentId == "" {
		return fmt.Errorf("compartment_id must be provided and cannot be empty")
	}
	request.CompartmentId = &compartmentId

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isActive, ok := s.D.GetOkExists("is_active"); ok {
		tmp := isActive.(bool)
		request.IsActive = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "demand_signal")

	response, err := s.Client.ListOccMetricAlarms(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccMetricAlarms(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DemandSignalOccMetricAlarmsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DemandSignalOccMetricAlarmsDataSource-", DemandSignalOccMetricAlarmsDataSource(), s.D))
	resources := []map[string]interface{}{}
	occMetricAlarm := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccMetricAlarmSummaryToMap(item))
	}
	occMetricAlarm["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DemandSignalOccMetricAlarmsDataSource().Schema["occ_metric_alarm_collection"].Elem.(*schema.Resource).Schema)
		occMetricAlarm["items"] = items
	}

	resources = append(resources, occMetricAlarm)
	if err := s.D.Set("occ_metric_alarm_collection", resources); err != nil {
		return err
	}

	return nil
}
