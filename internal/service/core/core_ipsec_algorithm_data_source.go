// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreIpsecAlgorithmDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreIpsecAlgorithm,
		Schema: map[string]*schema.Schema{
			// Computed
			"allowed_phase_one_parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"authentication_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dh_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"encryption_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"allowed_phase_two_parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"authentication_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"encryption_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"pfs_dh_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"default_phase_one_parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"default_authentication_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_dh_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_encryption_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"default_phase_two_parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"default_authentication_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_encryption_algorithms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_pfs_dh_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularCoreIpsecAlgorithm(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpsecAlgorithmDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpsecAlgorithmDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetAllowedIkeIPSecParametersResponse
}

func (s *CoreIpsecAlgorithmDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpsecAlgorithmDataSourceCrud) Get() error {
	request := oci_core.GetAllowedIkeIPSecParametersRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetAllowedIkeIPSecParameters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreIpsecAlgorithmDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpsecAlgorithmDataSource-", CoreIpsecAlgorithmDataSource(), s.D))

	if s.Res.AllowedPhaseOneParameters != nil {
		s.D.Set("allowed_phase_one_parameters", []interface{}{AllowedPhaseOneParametersToMap(s.Res.AllowedPhaseOneParameters)})
	} else {
		s.D.Set("allowed_phase_one_parameters", nil)
	}

	if s.Res.AllowedPhaseTwoParameters != nil {
		s.D.Set("allowed_phase_two_parameters", []interface{}{AllowedPhaseTwoParametersToMap(s.Res.AllowedPhaseTwoParameters)})
	} else {
		s.D.Set("allowed_phase_two_parameters", nil)
	}

	if s.Res.DefaultPhaseOneParameters != nil {
		s.D.Set("default_phase_one_parameters", []interface{}{DefaultPhaseOneParametersToMap(s.Res.DefaultPhaseOneParameters)})
	} else {
		s.D.Set("default_phase_one_parameters", nil)
	}

	if s.Res.DefaultPhaseTwoParameters != nil {
		s.D.Set("default_phase_two_parameters", []interface{}{DefaultPhaseTwoParametersToMap(s.Res.DefaultPhaseTwoParameters)})
	} else {
		s.D.Set("default_phase_two_parameters", nil)
	}

	return nil
}

func AllowedPhaseOneParametersToMap(obj *oci_core.AllowedPhaseOneParameters) map[string]interface{} {
	result := map[string]interface{}{}

	result["authentication_algorithms"] = obj.AuthenticationAlgorithms

	result["dh_groups"] = obj.DhGroups

	result["encryption_algorithms"] = obj.EncryptionAlgorithms

	return result
}

func AllowedPhaseTwoParametersToMap(obj *oci_core.AllowedPhaseTwoParameters) map[string]interface{} {
	result := map[string]interface{}{}

	result["authentication_algorithms"] = obj.AuthenticationAlgorithms

	result["encryption_algorithms"] = obj.EncryptionAlgorithms

	result["pfs_dh_groups"] = obj.PfsDhGroups

	return result
}

func DefaultPhaseOneParametersToMap(obj *oci_core.DefaultPhaseOneParameters) map[string]interface{} {
	result := map[string]interface{}{}

	result["default_authentication_algorithms"] = obj.DefaultAuthenticationAlgorithms

	result["default_dh_groups"] = obj.DefaultDhGroups

	result["default_encryption_algorithms"] = obj.DefaultEncryptionAlgorithms

	return result
}

func DefaultPhaseTwoParametersToMap(obj *oci_core.DefaultPhaseTwoParameters) map[string]interface{} {
	result := map[string]interface{}{}

	result["default_authentication_algorithms"] = obj.DefaultAuthenticationAlgorithms

	result["default_encryption_algorithms"] = obj.DefaultEncryptionAlgorithms

	if obj.DefaultPfsDhGroup != nil {
		result["default_pfs_dh_group"] = string(*obj.DefaultPfsDhGroup)
	}

	return result
}
