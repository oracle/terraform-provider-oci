// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardProblemEntitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardProblemEntities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"problem_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"problem_entity_collection": {
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
									"entity_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"problem_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"regions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"result_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_first_detected": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_detected": {
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

func readCloudGuardProblemEntities(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardProblemEntitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardProblemEntitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListProblemEntitiesResponse
}

func (s *CloudGuardProblemEntitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardProblemEntitiesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListProblemEntitiesRequest{}

	if problemId, ok := s.D.GetOkExists("problem_id"); ok {
		tmp := problemId.(string)
		request.ProblemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListProblemEntities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProblemEntities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardProblemEntitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardProblemEntitiesDataSource-", CloudGuardProblemEntitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	problemEntity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProblemEntitiesSummaryToMap(item))
	}
	problemEntity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardProblemEntitiesDataSource().Schema["problem_entity_collection"].Elem.(*schema.Resource).Schema)
		problemEntity["items"] = items
	}

	resources = append(resources, problemEntity)
	if err := s.D.Set("problem_entity_collection", resources); err != nil {
		return err
	}

	return nil
}

func EntitiesDetailsToMap(obj oci_cloud_guard.EntityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func ProblemEntitiesSummaryToMap(obj oci_cloud_guard.ProblemEntitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	entityDetails := []interface{}{}
	for _, item := range obj.EntityDetails {
		entityDetails = append(entityDetails, EntitiesDetailsToMap(item))
	}
	result["entity_details"] = entityDetails

	if obj.ProblemId != nil {
		result["problem_id"] = string(*obj.ProblemId)
	}

	result["regions"] = obj.Regions

	if obj.ResultUrl != nil {
		result["result_url"] = string(*obj.ResultUrl)
	}

	if obj.TimeFirstDetected != nil {
		result["time_first_detected"] = obj.TimeFirstDetected.String()
	}

	if obj.TimeLastDetected != nil {
		result["time_last_detected"] = obj.TimeLastDetected.String()
	}

	return result
}
