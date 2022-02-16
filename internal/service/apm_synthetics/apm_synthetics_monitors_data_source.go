// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v58/apmsynthetics"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsMonitorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmSyntheticsMonitors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitor_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"script_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitor_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApmSyntheticsMonitorResource()),
						},
					},
				},
			},
		},
	}
}

func readApmSyntheticsMonitors(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsMonitorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsMonitorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.ListMonitorsResponse
}

func (s *ApmSyntheticsMonitorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsMonitorsDataSourceCrud) Get() error {
	request := oci_apm_synthetics.ListMonitorsRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if monitorType, ok := s.D.GetOkExists("monitor_type"); ok {
		tmp := monitorType.(string)
		request.MonitorType = &tmp
	}

	if compositeId, ok := s.D.GetOkExists("script_id"); ok {
		tmp := compositeId.(string)
		scriptId, apmDomainId, err := parseScriptCompositeId(tmp)
		if err == nil {
			request.ScriptId = &scriptId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.ListMonitorsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.ListMonitors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMonitors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmSyntheticsMonitorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmSyntheticsMonitorsDataSource-", ApmSyntheticsMonitorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitor := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitorSummaryToMap(item))
	}
	monitor["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApmSyntheticsMonitorsDataSource().Schema["monitor_collection"].Elem.(*schema.Resource).Schema)
		monitor["items"] = items
	}

	resources = append(resources, monitor)
	if err := s.D.Set("monitor_collection", resources); err != nil {
		return err
	}

	return nil
}
