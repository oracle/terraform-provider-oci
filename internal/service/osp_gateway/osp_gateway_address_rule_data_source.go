// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v65/ospgateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OspGatewayAddressRuleDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOspGatewayAddressRule,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"country_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"address": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"fields": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"format": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"example": {
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
									"is_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"example": {
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
									"language": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"third_party_validation": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"contact": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"fields": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"format": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"example": {
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
									"is_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"example": {
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
									"language": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"tax": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"fields": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"format": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"example": {
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
									"is_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"example": {
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
									"language": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"value_set": {
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
		},
	}
}

func readSingularOspGatewayAddressRule(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewayAddressRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AddressRuleServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewayAddressRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.AddressRuleServiceClient
	Res    *oci_osp_gateway.GetAddressRuleResponse
}

func (s *OspGatewayAddressRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewayAddressRuleDataSourceCrud) Get() error {
	request := oci_osp_gateway.GetAddressRuleRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if countryCode, ok := s.D.GetOkExists("country_code"); ok {
		tmp := countryCode.(string)
		request.CountryCode = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osp_gateway")

	response, err := s.Client.GetAddressRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OspGatewayAddressRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OspGatewayAddressRuleDataSource-", OspGatewayAddressRuleDataSource(), s.D))

	if s.Res.Address != nil {
		s.D.Set("address", []interface{}{AddressTypeRuleToMap(s.Res.Address)})
	} else {
		s.D.Set("address", nil)
	}

	if s.Res.Contact != nil {
		s.D.Set("contact", []interface{}{ContactTypeRuleToMap(s.Res.Contact)})
	} else {
		s.D.Set("contact", nil)
	}

	if s.Res.Tax != nil {
		s.D.Set("tax", []interface{}{TaxTypeRuleToMap(s.Res.Tax)})
	} else {
		s.D.Set("tax", nil)
	}

	return nil
}

func AddressTypeRuleToMap(obj *oci_osp_gateway.AddressTypeRule) map[string]interface{} {
	result := map[string]interface{}{}

	fields := []interface{}{}
	for _, item := range obj.Fields {
		fields = append(fields, FieldToMap(item))
	}
	result["fields"] = fields

	result["third_party_validation"] = string(obj.ThirdPartyValidation)

	return result
}

func ContactTypeRuleToMap(obj *oci_osp_gateway.ContactTypeRule) map[string]interface{} {
	result := map[string]interface{}{}

	fields := []interface{}{}
	for _, item := range obj.Fields {
		fields = append(fields, FieldToMap(item))
	}
	result["fields"] = fields

	return result
}

func FieldToMap(obj oci_osp_gateway.Field) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Format != nil {
		result["format"] = []interface{}{FormatToMap(obj.Format)}
	}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.Label != nil {
		result["label"] = []interface{}{LabelToMap(obj.Label)}
	}

	if obj.Language != nil {
		result["language"] = string(*obj.Language)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func FormatToMap(obj *oci_osp_gateway.Format) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Example != nil {
		result["example"] = string(*obj.Example)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func LabelToMap(obj *oci_osp_gateway.Label) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Example != nil {
		result["example"] = string(*obj.Example)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func TaxTypeRuleToMap(obj *oci_osp_gateway.TaxTypeRule) map[string]interface{} {
	result := map[string]interface{}{}

	fields := []interface{}{}
	for _, item := range obj.Fields {
		fields = append(fields, FieldToMap(item))
	}
	result["fields"] = fields

	valueSet := []interface{}{}
	for _, item := range obj.ValueSet {
		valueSet = append(valueSet, ValueSetEntityToMap(item))
	}
	result["value_set"] = valueSet

	return result
}

func ValueSetEntityToMap(obj oci_osp_gateway.ValueSetEntity) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
