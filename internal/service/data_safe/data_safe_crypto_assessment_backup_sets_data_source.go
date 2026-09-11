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

func DataSafeCryptoAssessmentBackupSetsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentBackupSetsWithContext,
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
			"backup_set_key": {
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
			"is_encrypted": {
				Type:     schema.TypeBool,
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
			"crypto_assessment_backup_set_collection": {
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
									"algorithm_observed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"backup_pieces": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"backup_set_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"backup_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cipher_mode_observed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_compressed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_encrypted": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"set_stamp": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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

func readDataSafeCryptoAssessmentBackupSetsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentBackupSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentBackupSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentBackupSetsResponse
}

func (s *DataSafeCryptoAssessmentBackupSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentBackupSetsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentBackupSetsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListCryptoAssessmentBackupSetsAccessLevelEnum(accessLevel.(string))
	}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessmentType, ok := s.D.GetOkExists("assessment_type"); ok {
		request.AssessmentType = oci_data_safe.CryptoAssessmentTypeEnum(assessmentType.(string))
	}

	if backupSetKey, ok := s.D.GetOkExists("backup_set_key"); ok {
		tmp := backupSetKey.(string)
		request.BackupSetKey = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if isEncrypted, ok := s.D.GetOkExists("is_encrypted"); ok {
		tmp := isEncrypted.(bool)
		request.IsEncrypted = &tmp
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

	response, err := s.Client.ListCryptoAssessmentBackupSets(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAssessmentBackupSets(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeCryptoAssessmentBackupSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentBackupSetsDataSource-", DataSafeCryptoAssessmentBackupSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessmentBackupSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentBackupSetSummaryToMap(item))
	}
	cryptoAssessmentBackupSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentBackupSetsDataSource().Schema["crypto_assessment_backup_set_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessmentBackupSet["items"] = items
	}

	resources = append(resources, cryptoAssessmentBackupSet)
	if err := s.D.Set("crypto_assessment_backup_set_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAssessmentBackupSetSummaryToMap(obj oci_data_safe.CryptoAssessmentBackupSetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AlgorithmObserved != nil {
		result["algorithm_observed"] = string(*obj.AlgorithmObserved)
	}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	if obj.BackupPieces != nil {
		result["backup_pieces"] = int(*obj.BackupPieces)
	}

	if obj.BackupSetKey != nil {
		result["backup_set_key"] = string(*obj.BackupSetKey)
	}

	result["backup_type"] = string(obj.BackupType)

	if obj.CipherModeObserved != nil {
		result["cipher_mode_observed"] = string(*obj.CipherModeObserved)
	}

	if obj.IsCompressed != nil {
		result["is_compressed"] = bool(*obj.IsCompressed)
	}

	if obj.IsEncrypted != nil {
		result["is_encrypted"] = bool(*obj.IsEncrypted)
	}

	if obj.SetStamp != nil {
		result["set_stamp"] = string(*obj.SetStamp)
	}

	if obj.SizeInGBs != nil {
		result["size_in_gbs"] = float64(*obj.SizeInGBs)
	}

	result["status"] = string(obj.Status)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastAssessed != nil {
		result["time_last_assessed"] = obj.TimeLastAssessed.String()
	}

	return result
}
