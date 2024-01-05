// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetDiagnosesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetDiagnoses,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleet_diagnosis_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"resource_diagnosis": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readJmsFleetDiagnoses(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetDiagnosesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetDiagnosesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListFleetDiagnosesResponse
}

func (s *JmsFleetDiagnosesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetDiagnosesDataSourceCrud) Get() error {
	request := oci_jms.ListFleetDiagnosesRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListFleetDiagnoses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFleetDiagnoses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetDiagnosesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetDiagnosesDataSource-", JmsFleetDiagnosesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetDiagnose := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FleetDiagnosisSummaryToMap(item))
	}
	fleetDiagnose["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetDiagnosesDataSource().Schema["fleet_diagnosis_collection"].Elem.(*schema.Resource).Schema)
		fleetDiagnose["items"] = items
	}

	resources = append(resources, fleetDiagnose)
	if err := s.D.Set("fleet_diagnosis_collection", resources); err != nil {
		return err
	}

	return nil
}

func FleetDiagnosisSummaryToMap(obj oci_jms.FleetDiagnosisSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ResourceDiagnosis != nil {
		result["resource_diagnosis"] = string(*obj.ResourceDiagnosis)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	result["resource_state"] = string(obj.ResourceState)

	result["resource_type"] = string(obj.ResourceType)

	return result
}
