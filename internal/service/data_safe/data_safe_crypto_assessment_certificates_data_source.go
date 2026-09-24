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

func DataSafeCryptoAssessmentCertificatesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentCertificatesWithContext,
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
			"certificate_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"days_to_expiry": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expiry_bucket": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"signature_algorithm": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"crypto_assessment_certificate_collection": {
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
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"assessment_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"certificate_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"days_to_expiry": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"expiry_bucket": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"issuer": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_key_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"serial_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"signature_algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subject": {
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
									"time_valid_from": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
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

func readDataSafeCryptoAssessmentCertificatesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentCertificatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentCertificatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentCertificatesResponse
}

func (s *DataSafeCryptoAssessmentCertificatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentCertificatesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentCertificatesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListCryptoAssessmentCertificatesAccessLevelEnum(accessLevel.(string))
	}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessmentType, ok := s.D.GetOkExists("assessment_type"); ok {
		request.AssessmentType = oci_data_safe.CryptoAssessmentTypeEnum(assessmentType.(string))
	}

	if certificateType, ok := s.D.GetOkExists("certificate_type"); ok {
		interfaces := certificateType.([]interface{})
		tmp := make([]oci_data_safe.CryptoAssessmentCertificateSummaryCertificateTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.CryptoAssessmentCertificateSummaryCertificateTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("certificate_type") {
			request.CertificateType = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if daysToExpiry, ok := s.D.GetOkExists("days_to_expiry"); ok {
		tmp := daysToExpiry.(int)
		request.DaysToExpiry = &tmp
	}

	if expiryBucket, ok := s.D.GetOkExists("expiry_bucket"); ok {
		tmp := expiryBucket.(string)
		request.ExpiryBucket = &tmp
	}

	if publicKeyType, ok := s.D.GetOkExists("public_key_type"); ok {
		interfaces := publicKeyType.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("public_key_type") {
			request.PublicKeyType = tmp
		}
	}

	if signatureAlgorithm, ok := s.D.GetOkExists("signature_algorithm"); ok {
		interfaces := signatureAlgorithm.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("signature_algorithm") {
			request.SignatureAlgorithm = tmp
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		interfaces := status.([]interface{})
		tmp := make([]oci_data_safe.CryptoAssessmentCertificateSummaryStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.CryptoAssessmentCertificateSummaryStatusEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("status") {
			request.Status = tmp
		}
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

	response, err := s.Client.ListCryptoAssessmentCertificates(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAssessmentCertificates(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeCryptoAssessmentCertificatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentCertificatesDataSource-", DataSafeCryptoAssessmentCertificatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessmentCertificate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentCertificateSummaryToMap(item))
	}
	cryptoAssessmentCertificate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentCertificatesDataSource().Schema["crypto_assessment_certificate_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessmentCertificate["items"] = items
	}

	resources = append(resources, cryptoAssessmentCertificate)
	if err := s.D.Set("crypto_assessment_certificate_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAssessmentCertificateSummaryToMap(obj oci_data_safe.CryptoAssessmentCertificateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Age != nil {
		result["age"] = strconv.FormatInt(*obj.Age, 10)
	}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	result["assessment_type"] = string(obj.AssessmentType)

	result["certificate_type"] = string(obj.CertificateType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DaysToExpiry != nil {
		result["days_to_expiry"] = int(*obj.DaysToExpiry)
	}

	if obj.ExpiryBucket != nil {
		result["expiry_bucket"] = string(*obj.ExpiryBucket)
	}

	if obj.Issuer != nil {
		result["issuer"] = string(*obj.Issuer)
	}

	if obj.PublicKeyType != nil {
		result["public_key_type"] = string(*obj.PublicKeyType)
	}

	if obj.SerialNumber != nil {
		result["serial_number"] = string(*obj.SerialNumber)
	}

	if obj.SignatureAlgorithm != nil {
		result["signature_algorithm"] = string(*obj.SignatureAlgorithm)
	}

	result["status"] = string(obj.Status)

	if obj.Subject != nil {
		result["subject"] = string(*obj.Subject)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeLastAssessed != nil {
		result["time_last_assessed"] = obj.TimeLastAssessed.String()
	}

	if obj.TimeValidFrom != nil {
		result["time_valid_from"] = obj.TimeValidFrom.String()
	}

	if obj.TimeValidUntil != nil {
		result["time_valid_until"] = obj.TimeValidUntil.String()
	}

	if obj.WalletLocation != nil {
		result["wallet_location"] = string(*obj.WalletLocation)
	}

	return result
}
