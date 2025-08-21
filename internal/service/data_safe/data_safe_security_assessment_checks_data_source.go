// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityAssessmentChecksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentChecks,
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
			"contains_references": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"contains_severity": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"suggested_severity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"checks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DataSafeSecurityAssessmentCheckResource(),
			},
		},
	}
}

func readDataSafeSecurityAssessmentChecks(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentChecksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentChecksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListChecksResponse
}

func (s *DataSafeSecurityAssessmentChecksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentChecksDataSourceCrud) Get() error {
	request := oci_data_safe.ListChecksRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListChecksAccessLevelEnum(accessLevel.(string))
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if containsReferences, ok := s.D.GetOkExists("contains_references"); ok {
		interfaces := containsReferences.([]interface{})
		tmp := make([]oci_data_safe.SecurityAssessmentReferencesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.SecurityAssessmentReferencesEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("contains_references") {
			request.ContainsReferences = tmp
		}
	}

	if containsSeverity, ok := s.D.GetOkExists("contains_severity"); ok {
		interfaces := containsSeverity.([]interface{})
		tmp := make([]oci_data_safe.ListChecksContainsSeverityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListChecksContainsSeverityEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("contains_severity") {
			request.ContainsSeverity = tmp
		}
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	if suggestedSeverity, ok := s.D.GetOkExists("suggested_severity"); ok {
		request.SuggestedSeverity = oci_data_safe.ListChecksSuggestedSeverityEnum(suggestedSeverity.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListChecks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListChecks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentChecksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentChecksDataSource-", DataSafeSecurityAssessmentChecksDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		securityAssessmentCheck := map[string]interface{}{}

		if r.Category != nil {
			securityAssessmentCheck["category"] = *r.Category
		}

		if r.Key != nil {
			securityAssessmentCheck["key"] = *r.Key
		}

		if r.Oneline != nil {
			securityAssessmentCheck["oneline"] = *r.Oneline
		}

		if r.References != nil {
			securityAssessmentCheck["references"] = []interface{}{ReferencesToMap(r.References)}
		} else {
			securityAssessmentCheck["references"] = nil
		}

		if r.Remarks != nil {
			securityAssessmentCheck["remarks"] = *r.Remarks
		}

		securityAssessmentCheck["suggested_severity"] = r.SuggestedSeverity

		if r.Title != nil {
			securityAssessmentCheck["title"] = *r.Title
		}

		resources = append(resources, securityAssessmentCheck)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeSecurityAssessmentChecksDataSource().Schema["checks"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("checks", resources); err != nil {
		return err
	}

	return nil
}
