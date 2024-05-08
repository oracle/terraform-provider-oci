// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

package cloud_guard

/*
import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardResourceVulnerabilitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardResourceVulnerabilities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cve_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"risk_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_vulnerability_collection": {
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
									"cvss_score": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
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
									"package_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cause": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"location": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"package_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"remediation": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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
									"time_first_detected": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_detected": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_modified": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_published": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
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

func readCloudGuardResourceVulnerabilities(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResourceVulnerabilitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardResourceVulnerabilitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListResourceVulnerabilitiesResponse
}

func (s *CloudGuardResourceVulnerabilitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardResourceVulnerabilitiesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListResourceVulnerabilitiesRequest{}

	if cveId, ok := s.D.GetOkExists("cve_id"); ok {
		tmp := cveId.(string)
		request.CveId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if riskLevel, ok := s.D.GetOkExists("risk_level"); ok {
		tmp := riskLevel.(string)
		request.RiskLevel = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListResourceVulnerabilities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourceVulnerabilities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardResourceVulnerabilitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardResourceVulnerabilitiesDataSource-", CloudGuardResourceVulnerabilitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	resourceVulnerability := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceVulnerabilitySummaryToMap(item))
	}
	resourceVulnerability["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardResourceVulnerabilitiesDataSource().Schema["resource_vulnerability_collection"].Elem.(*schema.Resource).Schema)
		resourceVulnerability["items"] = items
	}

	resources = append(resources, resourceVulnerability)
	if err := s.D.Set("resource_vulnerability_collection", resources); err != nil {
		return err
	}

	return nil
}

func PackageDetailToMap(obj oci_cloud_guard.PackageDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cause != nil {
		result["cause"] = string(*obj.Cause)
	}

	if obj.Location != nil {
		result["location"] = string(*obj.Location)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PackageType != nil {
		result["package_type"] = string(*obj.PackageType)
	}

	if obj.Remediation != nil {
		result["remediation"] = string(*obj.Remediation)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func ResourceVulnerabilitySummaryToMap(obj oci_cloud_guard.ResourceVulnerabilitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["risk_level"] = string(obj.RiskLevel)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}*/
