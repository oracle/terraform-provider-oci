// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

package cloud_guard

/*
import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cve_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cvss_score": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cvss_score_greater_than": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cvss_score_less_than": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"detector_rule_id_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"detector_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"risk_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"risk_level_greater_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"risk_level_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_collection": {
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
									"additional_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"os_info": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"open_ports_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"problem_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"risk_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_first_monitored": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_monitored": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vulnerability_count": {
										Type:     schema.TypeInt,
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

func readCloudGuardResources(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListResourcesResponse
}

func (s *CloudGuardResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardResourcesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListResourcesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_cloud_guard.ListResourcesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if cveId, ok := s.D.GetOkExists("cve_id"); ok {
		tmp := cveId.(string)
		request.CveId = &tmp
	}

	if cvssScore, ok := s.D.GetOkExists("cvss_score"); ok {
		tmp := cvssScore.(int)
		request.CvssScore = &tmp
	}

	if cvssScoreGreaterThan, ok := s.D.GetOkExists("cvss_score_greater_than"); ok {
		tmp := cvssScoreGreaterThan.(int)
		request.CvssScoreGreaterThan = &tmp
	}

	if cvssScoreLessThan, ok := s.D.GetOkExists("cvss_score_less_than"); ok {
		tmp := cvssScoreLessThan.(int)
		request.CvssScoreLessThan = &tmp
	}

	if detectorRuleIdList, ok := s.D.GetOkExists("detector_rule_id_list"); ok {
		interfaces := detectorRuleIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("detector_rule_id_list") {
			request.DetectorRuleIdList = tmp
		}
	}

	if detectorType, ok := s.D.GetOkExists("detector_type"); ok {
		request.DetectorType = oci_cloud_guard.ListResourcesDetectorTypeEnum(detectorType.(string))
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if riskLevel, ok := s.D.GetOkExists("risk_level"); ok {
		tmp := riskLevel.(string)
		request.RiskLevel = &tmp
	}

	if riskLevelGreaterThan, ok := s.D.GetOkExists("risk_level_greater_than"); ok {
		tmp := riskLevelGreaterThan.(string)
		request.RiskLevelGreaterThan = &tmp
	}

	if riskLevelLessThan, ok := s.D.GetOkExists("risk_level_less_than"); ok {
		tmp := riskLevelLessThan.(string)
		request.RiskLevelLessThan = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardResourcesDataSource-", CloudGuardResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	resource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceSummaryToMap(item))
	}
	resource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardResourcesDataSource().Schema["resource_collection"].Elem.(*schema.Resource).Schema)
		resource["items"] = items
	}

	resources = append(resources, resource)
	if err := s.D.Set("resource_collection", resources); err != nil {
		return err
	}

	return nil
}

func ResourceAdditionalDetailsToMap(obj *oci_cloud_guard.ResourceAdditionalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OsInfo != nil {
		result["os_info"] = string(*obj.OsInfo)
	}

	return result
}

func ResourceSummaryToMap(obj oci_cloud_guard.ResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.OpenPortsCount != nil {
		result["open_ports_count"] = int(*obj.OpenPortsCount)
	}

	if obj.ProblemCount != nil {
		result["problem_count"] = int(*obj.ProblemCount)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["risk_level"] = string(obj.RiskLevel)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetName != nil {
		result["target_name"] = string(*obj.TargetName)
	}

	if obj.VulnerabilityCount != nil {
		result["vulnerability_count"] = int(*obj.VulnerabilityCount)
	}

	return result
}*/
