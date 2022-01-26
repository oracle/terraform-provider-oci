// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v56/datasafe"
)

func DataSafeSecurityAssessmentFindingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentFindings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"finding_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"findings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"assessment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"details": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"references": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"cis": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gdpr": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"stig": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"remarks": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeSecurityAssessmentFindings(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentFindingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentFindingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListFindingsResponse
}

func (s *DataSafeSecurityAssessmentFindingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentFindingsDataSourceCrud) Get() error {
	request := oci_data_safe.ListFindingsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListFindingsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if findingKey, ok := s.D.GetOkExists("finding_key"); ok {
		tmp := findingKey.(string)
		request.FindingKey = &tmp
	}

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_data_safe.ListFindingsSeverityEnum(severity.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListFindings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFindings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentFindingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentFindingsDataSource-", DataSafeSecurityAssessmentFindingsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		securityAssessmentFinding := map[string]interface{}{}

		if r.AssessmentId != nil {
			securityAssessmentFinding["assessment_id"] = *r.AssessmentId
		}

		if r.Details != nil {
			securityAssessmentFinding["details"] = []interface{}{}
		} else {
			securityAssessmentFinding["details"] = nil
		}

		if r.Key != nil {
			securityAssessmentFinding["key"] = *r.Key
		}

		if r.References != nil {
			securityAssessmentFinding["references"] = []interface{}{FindingsReferencesToMap(r.References)}
		} else {
			securityAssessmentFinding["references"] = nil
		}

		if r.Remarks != nil {
			securityAssessmentFinding["remarks"] = *r.Remarks
		}

		securityAssessmentFinding["severity"] = r.Severity

		if r.Summary != nil {
			securityAssessmentFinding["summary"] = *r.Summary
		}

		if r.TargetId != nil {
			securityAssessmentFinding["target_id"] = *r.TargetId
		}

		if r.Title != nil {
			securityAssessmentFinding["title"] = *r.Title
		}

		resources = append(resources, securityAssessmentFinding)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeSecurityAssessmentFindingsDataSource().Schema["findings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("findings", resources); err != nil {
		return err
	}

	return nil
}

func FindingsReferencesToMap(obj *oci_data_safe.References) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cis != nil {
		result["cis"] = string(*obj.Cis)
	}

	if obj.Gdpr != nil {
		result["gdpr"] = string(*obj.Gdpr)
	}

	if obj.Stig != nil {
		result["stig"] = string(*obj.Stig)
	}

	return result
}
