// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsOnPremiseVantagePointWorkersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmSyntheticsOnPremiseVantagePointWorkers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"capability": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_premise_vantage_point_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"worker_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApmSyntheticsOnPremiseVantagePointWorkerResource()),
						},
					},
				},
			},
		},
	}
}

func readApmSyntheticsOnPremiseVantagePointWorkers(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointWorkersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsOnPremiseVantagePointWorkersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.ListWorkersResponse
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkersDataSourceCrud) Get() error {
	request := oci_apm_synthetics.ListWorkersRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if capability, ok := s.D.GetOkExists("capability"); ok {
		tmp := capability.(string)
		request.Capability = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if compositeId, ok := s.D.GetOkExists("on_premise_vantage_point_id"); ok {
		tmp := compositeId.(string)
		onPremiseVantagePointId, apmDomainId, err := parseOnPremiseVantagePointCompositeId(tmp)
		if err == nil {
			request.OnPremiseVantagePointId = &onPremiseVantagePointId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current on_premise_vantage_point_id: %s", compositeId)
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.ListWorkersStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.ListWorkers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWorkers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmSyntheticsOnPremiseVantagePointWorkersDataSource-", ApmSyntheticsOnPremiseVantagePointWorkersDataSource(), s.D))
	resources := []map[string]interface{}{}
	onPremiseVantagePointWorker := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WorkerSummaryToMap(item))
	}
	onPremiseVantagePointWorker["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApmSyntheticsOnPremiseVantagePointWorkersDataSource().Schema["worker_collection"].Elem.(*schema.Resource).Schema)
		onPremiseVantagePointWorker["items"] = items
	}

	resources = append(resources, onPremiseVantagePointWorker)
	if err := s.D.Set("worker_collection", resources); err != nil {
		return err
	}

	return nil
}
