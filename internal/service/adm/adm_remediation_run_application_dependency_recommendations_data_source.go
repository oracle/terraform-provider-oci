// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRunApplicationDependencyRecommendationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAdmRemediationRunApplicationDependencyRecommendations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"gav": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"purl": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remediation_run_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_dependency_recommendation_collection": {
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
									"application_dependency_node_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"gav": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"purl": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recommended_gav": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recommended_purl": {
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

func readAdmRemediationRunApplicationDependencyRecommendations(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunApplicationDependencyRecommendationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

type AdmRemediationRunApplicationDependencyRecommendationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_adm.ApplicationDependencyManagementClient
	Res    *oci_adm.ListApplicationDependencyRecommendationsResponse
}

func (s *AdmRemediationRunApplicationDependencyRecommendationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AdmRemediationRunApplicationDependencyRecommendationsDataSourceCrud) Get() error {
	request := oci_adm.ListApplicationDependencyRecommendationsRequest{}

	if gav, ok := s.D.GetOkExists("gav"); ok {
		tmp := gav.(string)
		request.Gav = &tmp
	}

	if purl, ok := s.D.GetOkExists("purl"); ok {
		tmp := purl.(string)
		request.Purl = &tmp
	}

	if remediationRunId, ok := s.D.GetOkExists("remediation_run_id"); ok {
		tmp := remediationRunId.(string)
		request.RemediationRunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "adm")

	response, err := s.Client.ListApplicationDependencyRecommendations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApplicationDependencyRecommendations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AdmRemediationRunApplicationDependencyRecommendationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AdmRemediationRunApplicationDependencyRecommendationsDataSource-", AdmRemediationRunApplicationDependencyRecommendationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	remediationRunApplicationDependencyRecommendation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ApplicationDependencyRecommendationSummaryToMap(item))
	}
	remediationRunApplicationDependencyRecommendation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AdmRemediationRunApplicationDependencyRecommendationsDataSource().Schema["application_dependency_recommendation_collection"].Elem.(*schema.Resource).Schema)
		remediationRunApplicationDependencyRecommendation["items"] = items
	}

	resources = append(resources, remediationRunApplicationDependencyRecommendation)
	if err := s.D.Set("application_dependency_recommendation_collection", resources); err != nil {
		return err
	}

	return nil
}

func ApplicationDependencyRecommendationSummaryToMap(obj oci_adm.ApplicationDependencyRecommendationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["application_dependency_node_ids"] = obj.ApplicationDependencyNodeIds

	if obj.Gav != nil {
		result["gav"] = string(*obj.Gav)
	}

	if obj.NodeId != nil {
		result["node_id"] = string(*obj.NodeId)
	}

	if obj.Purl != nil {
		result["purl"] = string(*obj.Purl)
	}

	if obj.RecommendedGav != nil {
		result["recommended_gav"] = string(*obj.RecommendedGav)
	}

	if obj.RecommendedPurl != nil {
		result["recommended_purl"] = string(*obj.RecommendedPurl)
	}

	return result
}
