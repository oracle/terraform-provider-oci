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

func DataSafeCryptoAssessmentTdeObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentTdeObjectsWithContext,
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
			"assessment_type": {
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
			"encryption_observed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"encryption_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quantum_readiness": {
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
			"crypto_assessment_tde_object_collection": {
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
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"column_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"encryption_observed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mode_observed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantum_readiness": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"table_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tablespace_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_assessed": {
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

func readDataSafeCryptoAssessmentTdeObjectsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentTdeObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentTdeObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentTdeObjectsResponse
}

func (s *DataSafeCryptoAssessmentTdeObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentTdeObjectsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentTdeObjectsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListCryptoAssessmentTdeObjectsAccessLevelEnum(accessLevel.(string))
	}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessmentType, ok := s.D.GetOkExists("assessment_type"); ok {
		request.AssessmentType = oci_data_safe.CryptoAssessmentTypeEnum(assessmentType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if encryptionObserved, ok := s.D.GetOkExists("encryption_observed"); ok {
		interfaces := encryptionObserved.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("encryption_observed") {
			request.EncryptionObserved = tmp
		}
	}

	if encryptionStatus, ok := s.D.GetOkExists("encryption_status"); ok {
		request.EncryptionStatus = oci_data_safe.ListCryptoAssessmentTdeObjectsEncryptionStatusEnum(encryptionStatus.(string))
	}

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		request.ObjectType = oci_data_safe.ListCryptoAssessmentTdeObjectsObjectTypeEnum(objectType.(string))
	}

	if quantumReadiness, ok := s.D.GetOkExists("quantum_readiness"); ok {
		request.QuantumReadiness = oci_data_safe.ListCryptoAssessmentTdeObjectsQuantumReadinessEnum(quantumReadiness.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListCryptoAssessmentTdeObjects(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAssessmentTdeObjects(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeCryptoAssessmentTdeObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentTdeObjectsDataSource-", DataSafeCryptoAssessmentTdeObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessmentTdeObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentTdeObjectSummaryToMap(item))
	}
	cryptoAssessmentTdeObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentTdeObjectsDataSource().Schema["crypto_assessment_tde_object_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessmentTdeObject["items"] = items
	}

	resources = append(resources, cryptoAssessmentTdeObject)
	if err := s.D.Set("crypto_assessment_tde_object_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAssessmentTdeObjectSummaryToMap(obj oci_data_safe.CryptoAssessmentTdeObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	if obj.EncryptionObserved != nil {
		result["encryption_observed"] = string(*obj.EncryptionObserved)
	}

	if obj.ModeObserved != nil {
		result["mode_observed"] = string(*obj.ModeObserved)
	}

	result["quantum_readiness"] = string(obj.QuantumReadiness)

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SizeInGBs != nil {
		result["size_in_gbs"] = float64(*obj.SizeInGBs)
	}

	if obj.TableName != nil {
		result["table_name"] = string(*obj.TableName)
	}

	if obj.TablespaceName != nil {
		result["tablespace_name"] = string(*obj.TablespaceName)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeLastAssessed != nil {
		result["time_last_assessed"] = obj.TimeLastAssessed.String()
	}

	return result
}
