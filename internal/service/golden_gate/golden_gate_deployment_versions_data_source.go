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

func GoldenGateDeploymentVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateDeploymentVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_version_collection": {
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
									"deployment_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_security_fix": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"ogg_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"release_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_released": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_supported_until": {
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

func readGoldenGateDeploymentVersions(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListDeploymentVersionsResponse
}

func (s *GoldenGateDeploymentVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentVersionsDataSourceCrud) Get() error {
	request := oci_golden_gate.ListDeploymentVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
		request.DeploymentType = oci_golden_gate.ListDeploymentVersionsDeploymentTypeEnum(deploymentType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListDeploymentVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeploymentVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateDeploymentVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateDeploymentVersionsDataSource-", GoldenGateDeploymentVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	deploymentVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DeploymentVersionSummaryToMap(item))
	}
	deploymentVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateDeploymentVersionsDataSource().Schema["deployment_version_collection"].Elem.(*schema.Resource).Schema)
		deploymentVersion["items"] = items
	}

	resources = append(resources, deploymentVersion)
	if err := s.D.Set("deployment_version_collection", resources); err != nil {
		return err
	}

	return nil
}

func DeploymentVersionSummaryToMap(obj oci_golden_gate.DeploymentVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["deployment_type"] = string(obj.DeploymentType)

	if obj.IsSecurityFix != nil {
		result["is_security_fix"] = bool(*obj.IsSecurityFix)
	}

	if obj.OggVersion != nil {
		result["ogg_version"] = string(*obj.OggVersion)
	}

	result["release_type"] = string(obj.ReleaseType)

	if obj.TimeReleased != nil {
		result["time_released"] = obj.TimeReleased.String()
	}

	if obj.TimeSupportedUntil != nil {
		result["time_supported_until"] = obj.TimeSupportedUntil.String()
	}

	return result
}
