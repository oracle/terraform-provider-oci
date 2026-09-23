// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiModelDiscoveriesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGenerativeAiModelDiscoveriesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"api_capability": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"capability": {
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
			"is_dedicated_retired": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_deprecated": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_on_demand_retired": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"model_access": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"realm": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"serving_mode": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"model_discovery_collection": {
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
									"api_capability": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"availability": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"realm": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"serving_modes": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"supported_replacements": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"time_dedicated_retired": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_deprecated": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_on_demand_retired": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"capabilities": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"modality_support": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"input": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"output": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"model_access": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"model_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameters": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"maximum": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"minimum": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"vendor": {
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

func readGenerativeAiModelDiscoveriesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiModelDiscoveriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiModelDiscoveriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.ListModelDiscoveryResponse
}

func (s *GenerativeAiModelDiscoveriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiModelDiscoveriesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.ListModelDiscoveryRequest{}

	if apiCapability, ok := s.D.GetOkExists("api_capability"); ok {
		interfaces := apiCapability.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("api_capability") {
			request.ApiCapability = tmp
		}
	}

	if capability, ok := s.D.GetOkExists("capability"); ok {
		interfaces := capability.([]interface{})
		tmp := make([]oci_generative_ai.ModelCapabilityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_generative_ai.ModelCapabilityEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("capability") {
			request.Capability = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isDedicatedRetired, ok := s.D.GetOkExists("is_dedicated_retired"); ok {
		tmp := isDedicatedRetired.(bool)
		request.IsDedicatedRetired = &tmp
	}

	if isDeprecated, ok := s.D.GetOkExists("is_deprecated"); ok {
		tmp := isDeprecated.(bool)
		request.IsDeprecated = &tmp
	}

	if isOnDemandRetired, ok := s.D.GetOkExists("is_on_demand_retired"); ok {
		tmp := isOnDemandRetired.(bool)
		request.IsOnDemandRetired = &tmp
	}

	if modelAccess, ok := s.D.GetOkExists("model_access"); ok {
		interfaces := modelAccess.([]interface{})
		tmp := make([]oci_generative_ai.ListModelDiscoveryModelAccessEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_generative_ai.ListModelDiscoveryModelAccessEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("model_access") {
			request.ModelAccess = tmp
		}
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if realm, ok := s.D.GetOkExists("realm"); ok {
		interfaces := realm.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("realm") {
			request.Realm = tmp
		}
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		interfaces := region.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("region") {
			request.Region = tmp
		}
	}

	if servingMode, ok := s.D.GetOkExists("serving_mode"); ok {
		interfaces := servingMode.([]interface{})
		tmp := make([]oci_generative_ai.ListModelDiscoveryServingModeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_generative_ai.ListModelDiscoveryServingModeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("serving_mode") {
			request.ServingMode = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.ListModelDiscovery(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelDiscovery(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiModelDiscoveriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiModelDiscoveriesDataSource-", GenerativeAiModelDiscoveriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	modelDiscovery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ModelDiscoveryToMap(item))
	}
	modelDiscovery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiModelDiscoveriesDataSource().Schema["model_discovery_collection"].Elem.(*schema.Resource).Schema)
		modelDiscovery["items"] = items
	}

	resources = append(resources, modelDiscovery)
	if err := s.D.Set("model_discovery_collection", resources); err != nil {
		return err
	}

	return nil
}

func AvailabilityToMap(obj oci_generative_ai.Availability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Realm != nil {
		result["realm"] = string(*obj.Realm)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["serving_modes"] = obj.ServingModes

	result["supported_replacements"] = obj.SupportedReplacements

	if obj.TimeDedicatedRetired != nil {
		result["time_dedicated_retired"] = obj.TimeDedicatedRetired.String()
	}

	if obj.TimeDeprecated != nil {
		result["time_deprecated"] = obj.TimeDeprecated.String()
	}

	if obj.TimeOnDemandRetired != nil {
		result["time_on_demand_retired"] = obj.TimeOnDemandRetired.String()
	}

	return result
}

func ModelDiscoveryToMap(obj oci_generative_ai.ModelDiscovery) map[string]interface{} {
	result := map[string]interface{}{}

	result["api_capability"] = obj.ApiCapability

	availability := []interface{}{}
	for _, item := range obj.Availability {
		availability = append(availability, AvailabilityToMap(item))
	}
	result["availability"] = availability

	result["capabilities"] = obj.Capabilities

	modalitySupport := []interface{}{}
	for _, item := range obj.ModalitySupport {
		modalitySupport = append(modalitySupport, ModelModalitySupportToMap(item))
	}
	result["modality_support"] = modalitySupport

	result["model_access"] = string(obj.ModelAccess)

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	parameters := []interface{}{}
	for _, item := range obj.Parameters {
		parameters = append(parameters, ParameterToMap(item))
	}
	result["parameters"] = parameters

	if obj.Vendor != nil {
		result["vendor"] = string(*obj.Vendor)
	}

	return result
}

func ModelModalitySupportToMap(obj oci_generative_ai.ModelModalitySupport) map[string]interface{} {
	result := map[string]interface{}{}

	result["input"] = string(obj.Input)

	result["output"] = string(obj.Output)

	return result
}

func ParameterToMap(obj oci_generative_ai.Parameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Maximum != nil {
		result["maximum"] = float64(*obj.Maximum)
	}

	if obj.Minimum != nil {
		result["minimum"] = float64(*obj.Minimum)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
