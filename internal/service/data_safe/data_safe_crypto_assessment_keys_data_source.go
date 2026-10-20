// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentKeysDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentKeysWithContext,
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
			"feature": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_manager_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"key_type": {
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
			"crypto_assessment_key_collection": {
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
									"age": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"feature": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_cache": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"keystore_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"secondary_keystore_type": {
										Type:     schema.TypeString,
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
									"time_last_rotation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"wallet_location": {
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

func readDataSafeCryptoAssessmentKeysWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentKeysResponse
}

func (s *DataSafeCryptoAssessmentKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentKeysDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentKeysRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListCryptoAssessmentKeysAccessLevelEnum(accessLevel.(string))
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

	if feature, ok := s.D.GetOkExists("feature"); ok {
		request.Feature = oci_data_safe.ListCryptoAssessmentKeysFeatureEnum(feature.(string))
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if keyManagerType, ok := s.D.GetOkExists("key_manager_type"); ok {
		interfaces := keyManagerType.([]interface{})
		tmp := make([]oci_data_safe.CryptoKeystoreTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.CryptoKeystoreTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("key_manager_type") {
			request.KeyManagerType = tmp
		}
	}

	if keyType, ok := s.D.GetOkExists("key_type"); ok {
		request.KeyType = oci_data_safe.CryptoAssessmentKeySummaryKeyTypeEnum(keyType.(string))
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

	response, err := s.Client.ListCryptoAssessmentKeys(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAssessmentKeys(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeCryptoAssessmentKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentKeysDataSource-", DataSafeCryptoAssessmentKeysDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessmentKey := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentKeySummaryToMap(item))
	}
	cryptoAssessmentKey["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentKeysDataSource().Schema["crypto_assessment_key_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessmentKey["items"] = items
	}

	resources = append(resources, cryptoAssessmentKey)
	if err := s.D.Set("crypto_assessment_key_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAssessmentKeySummaryToMap(obj oci_data_safe.CryptoAssessmentKeySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Age != nil {
		result["age"] = strconv.FormatInt(*obj.Age, 10)
	}

	if obj.Algorithm != nil {
		result["algorithm"] = string(*obj.Algorithm)
	}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	result["feature"] = string(obj.Feature)

	result["key_cache"] = string(obj.KeyCache)

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	result["key_type"] = string(obj.KeyType)

	result["keystore_type"] = string(obj.KeystoreType)

	result["secondary_keystore_type"] = string(obj.SecondaryKeystoreType)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastAssessed != nil {
		result["time_last_assessed"] = obj.TimeLastAssessed.String()
	}

	if obj.TimeLastRotation != nil {
		result["time_last_rotation"] = obj.TimeLastRotation.String()
	}

	if obj.WalletLocation != nil {
		result["wallet_location"] = string(*obj.WalletLocation)
	}

	return result
}
