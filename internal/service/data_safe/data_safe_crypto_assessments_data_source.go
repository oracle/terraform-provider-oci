// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assessment_id": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_assessment_scheduled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"posture_category": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_database_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"crypto_assessment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeCryptoAssessmentResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeCryptoAssessmentsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentsResponse
}

func (s *DataSafeCryptoAssessmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListCryptoAssessmentsAccessLevelEnum(accessLevel.(string))
	}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isAssessmentScheduled, ok := s.D.GetOkExists("is_assessment_scheduled"); ok {
		tmp := isAssessmentScheduled.(bool)
		request.IsAssessmentScheduled = &tmp
	}

	if postureCategory, ok := s.D.GetOkExists("posture_category"); ok {
		interfaces := postureCategory.([]interface{})
		tmp := make([]oci_data_safe.CryptoPostureCategoryEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.CryptoPostureCategoryEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("posture_category") {
			request.PostureCategory = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.CryptoAssessmentLifecycleStateEnum(state.(string))
	}

	if targetDatabaseGroupId, ok := s.D.GetOkExists("target_database_group_id"); ok {
		tmp := targetDatabaseGroupId.(string)
		request.TargetDatabaseGroupId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetIds, ok := s.D.GetOkExists("target_ids"); ok {
		interfaces := targetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("target_ids") {
			request.TargetIds = tmp
		}
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_data_safe.CryptoAssessmentTargetTypeEnum(targetType.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_data_safe.CryptoAssessmentTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListCryptoAssessments(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAssessments(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeCryptoAssessmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentsDataSource-", DataSafeCryptoAssessmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentSummaryToMap(item))
	}
	cryptoAssessment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentsDataSource().Schema["crypto_assessment_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessment["items"] = items
	}

	resources = append(resources, cryptoAssessment)
	if err := s.D.Set("crypto_assessment_collection", resources); err != nil {
		return err
	}

	return nil
}
