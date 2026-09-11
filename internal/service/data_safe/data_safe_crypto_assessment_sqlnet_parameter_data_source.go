// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentSqlnetParameterDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDataSafeCryptoAssessmentSqlnetParameterWithContext,
		Schema: map[string]*schema.Schema{
			"crypto_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parameter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quantum_readiness": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantum_readiness": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"source": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDataSafeCryptoAssessmentSqlnetParameterWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentSqlnetParameterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentSqlnetParameterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetCryptoAssessmentSqlnetParametersResponse
}

func (s *DataSafeCryptoAssessmentSqlnetParameterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentSqlnetParameterDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetCryptoAssessmentSqlnetParametersRequest{}

	if cryptoAssessmentId, ok := s.D.GetOkExists("crypto_assessment_id"); ok {
		tmp := cryptoAssessmentId.(string)
		request.CryptoAssessmentId = &tmp
	}

	if parameter, ok := s.D.GetOkExists("parameter"); ok {
		tmp := parameter.(string)
		request.Parameter = &tmp
	}

	if quantumReadiness, ok := s.D.GetOkExists("quantum_readiness"); ok {
		request.QuantumReadiness = oci_data_safe.GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum(quantumReadiness.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetCryptoAssessmentSqlnetParameters(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeCryptoAssessmentSqlnetParameterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentSqlnetParameterDataSource-", DataSafeCryptoAssessmentSqlnetParameterDataSource(), s.D))

	parameters := []interface{}{}
	for _, item := range s.Res.Parameters {
		parameters = append(parameters, CryptoSqlnetParameterToMap(item))
	}
	s.D.Set("parameters", parameters)

	if s.Res.Source != nil {
		s.D.Set("source", []interface{}{CryptoSqlnetParameterSourceToMap(s.Res.Source)})
	} else {
		s.D.Set("source", nil)
	}

	return nil
}

func CryptoSqlnetParameterToMap(obj oci_data_safe.CryptoSqlnetParameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["quantum_readiness"] = string(obj.QuantumReadiness)

	if obj.Value != nil {
		result["value"] = []interface{}{CryptoSqlnetParameterValueToMap(obj.Value)}
	}

	return result
}

func CryptoSqlnetParameterSourceToMap(obj *oci_data_safe.CryptoSqlnetParameterSource) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	return result
}

func CryptoSqlnetParameterValueToMap(obj *oci_data_safe.CryptoSqlnetParameterValue) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = fmt.Sprint(*obj.Value)
	}

	return result
}
