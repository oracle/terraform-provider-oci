// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateDeploymentTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateDeploymentTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ogg_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_type_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connection_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"default_username": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deployment_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ogg_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_technologies": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"supported_capabilities": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"supported_technologies_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_technologies": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func readGoldenGateDeploymentTypes(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListDeploymentTypesResponse
}

func (s *GoldenGateDeploymentTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentTypesDataSourceCrud) Get() error {
	request := oci_golden_gate.ListDeploymentTypesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
		request.DeploymentType = oci_golden_gate.ListDeploymentTypesDeploymentTypeEnum(deploymentType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if oggVersion, ok := s.D.GetOkExists("ogg_version"); ok {
		tmp := oggVersion.(string)
		request.OggVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListDeploymentTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeploymentTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateDeploymentTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateDeploymentTypesDataSource-", GoldenGateDeploymentTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	deploymentType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DeploymentTypeSummaryToMap(item))
	}
	deploymentType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateDeploymentTypesDataSource().Schema["deployment_type_collection"].Elem.(*schema.Resource).Schema)
		deploymentType["items"] = items
	}

	resources = append(resources, deploymentType)
	if err := s.D.Set("deployment_type_collection", resources); err != nil {
		return err
	}

	return nil
}

func DeploymentTypeSummaryToMap(obj oci_golden_gate.DeploymentTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["category"] = string(obj.Category)

	result["connection_types"] = obj.ConnectionTypes

	if obj.DefaultUsername != nil {
		result["default_username"] = string(*obj.DefaultUsername)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.OggVersion != nil {
		result["ogg_version"] = string(*obj.OggVersion)
	}

	result["source_technologies"] = obj.SourceTechnologies

	result["supported_capabilities"] = obj.SupportedCapabilities

	if obj.SupportedTechnologiesUrl != nil {
		result["supported_technologies_url"] = string(*obj.SupportedTechnologiesUrl)
	}

	result["target_technologies"] = obj.TargetTechnologies

	return result
}
