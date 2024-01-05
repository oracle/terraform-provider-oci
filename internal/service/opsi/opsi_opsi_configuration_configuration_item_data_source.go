// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiOpsiConfigurationConfigurationItemDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOpsiOpsiConfigurationConfigurationItem,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_item_field": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"config_items_applicable_context": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"opsi_config_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"config_items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"applicable_contexts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"config_item_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"config_item_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"value_input_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_value_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"max_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"min_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"possible_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value_source_config": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularOpsiOpsiConfigurationConfigurationItem(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOpsiConfigurationConfigurationItemDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOpsiConfigurationConfigurationItemDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.SummarizeConfigurationItemsResponse
}

func (s *OpsiOpsiConfigurationConfigurationItemDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOpsiConfigurationConfigurationItemDataSourceCrud) Get() error {
	request := oci_opsi.SummarizeConfigurationItemsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configItemField, ok := s.D.GetOkExists("config_item_field"); ok {
		interfaces := configItemField.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.SummarizeConfigurationItemsConfigItemFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingSummarizeConfigurationItemsConfigItemFieldEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("config_item_field") {
			request.ConfigItemField = tmp2
		}
	}

	if configItemsApplicableContext, ok := s.D.GetOkExists("config_items_applicable_context"); ok {
		interfaces := configItemsApplicableContext.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("config_items_applicable_context") {
			request.ConfigItemsApplicableContext = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if opsiConfigType, ok := s.D.GetOkExists("opsi_config_type"); ok {
		request.OpsiConfigType = oci_opsi.SummarizeConfigurationItemsOpsiConfigTypeEnum(opsiConfigType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.SummarizeConfigurationItems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiOpsiConfigurationConfigurationItemDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiOpsiConfigurationConfigurationItemDataSource-", OpsiOpsiConfigurationConfigurationItemDataSource(), s.D))
	switch v := (s.Res.ConfigurationItemsCollection).(type) {
	case oci_opsi.UxConfigurationItemsCollection:
		s.D.Set("opsi_config_type", "UX_CONFIGURATION")
		var configItems = v.GetConfigItems()
		if configItems != nil {
			configItems := []interface{}{}
			for _, item := range v.ConfigItems {
				configItems = append(configItems, OpsiConfigurationConfigurationItemsToMap(item))
			}
			s.D.Set("config_items", configItems)
		}
	default:
		log.Printf("[WARN] Received 'opsi_config_type' of unknown type %v", s.Res.ConfigurationItemsCollection)
		return nil
	}

	return nil
}

func OpsiConfigurationConfigurationItemsToMap(obj oci_opsi.ConfigurationItemSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.BasicConfigurationItemSummary:
		result["config_item_type"] = "BASIC"

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}

		if v.DefaultValue != nil {
			result["default_value"] = string(*v.DefaultValue)
		}

		if v.ApplicableContexts != nil && len(v.ApplicableContexts) != 0 {
			result["applicable_contexts"] = v.ApplicableContexts
		}

		result["config_item_type"] = "BASIC"

		if v.Metadata != nil {
			metadataArray := []interface{}{}
			if metadataMap := ConfigurationItemMetadataToMap(&v.Metadata); metadataMap != nil {
				metadataArray = append(metadataArray, metadataMap)
			}
			result["metadata"] = metadataArray
		}

		result["value_source_config"], _ = oci_opsi.GetMappingConfigurationItemValueSourceConfigurationTypeEnum(string(v.ValueSourceConfig))

	default:
		log.Printf("[WARN] Received 'config_item_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *OpsiOpsiConfigurationConfigurationItemDataSourceCrud) mapToConfigurationItemAllowedValueDetails(fieldKeyFormat string) (oci_opsi.ConfigurationItemAllowedValueDetails, error) {
	var baseObject oci_opsi.ConfigurationItemAllowedValueDetails
	//discriminator
	allowedValueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_value_type"))
	var allowedValueType string
	if ok {
		allowedValueType = allowedValueTypeRaw.(string)
	} else {
		allowedValueType = "" // default value
	}
	switch strings.ToLower(allowedValueType) {
	case strings.ToLower("FREE_TEXT"):
		details := oci_opsi.ConfigurationItemFreeTextAllowedValueDetails{}
		baseObject = details
	case strings.ToLower("LIMIT"):
		details := oci_opsi.ConfigurationItemLimitAllowedValueDetails{}
		if maxValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_value")); ok {
			tmp := maxValue.(string)
			details.MaxValue = &tmp
		}
		if minValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_value")); ok {
			tmp := minValue.(string)
			details.MinValue = &tmp
		}
		baseObject = details
	case strings.ToLower("PICK"):
		details := oci_opsi.ConfigurationItemPickAllowedValueDetails{}
		if possibleValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "possible_values")); ok {
			interfaces := possibleValues.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "possible_values")) {
				details.PossibleValues = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown allowed_value_type '%v' was specified", allowedValueType)
	}
	return baseObject, nil
}

// Commented out since same func is present in opsi_opsi_configuration_resource
/* func ConfigurationItemAllowedValueDetailsToMap(obj *oci_opsi.ConfigurationItemAllowedValueDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_opsi.ConfigurationItemFreeTextAllowedValueDetails:
		result["allowed_value_type"] = "FREE_TEXT"
	case oci_opsi.ConfigurationItemLimitAllowedValueDetails:
		result["allowed_value_type"] = "LIMIT"

		if v.MaxValue != nil {
			result["max_value"] = string(*v.MaxValue)
		}

		if v.MinValue != nil {
			result["min_value"] = string(*v.MinValue)
		}
	case oci_opsi.ConfigurationItemPickAllowedValueDetails:
		result["allowed_value_type"] = "PICK"

		result["possible_values"] = v.PossibleValues
		result["possible_values"] = v.PossibleValues
	default:
		log.Printf("[WARN] Received 'allowed_value_type' of unknown type %v", *obj)
		return nil
	}

	return result
} */

func (s *OpsiOpsiConfigurationConfigurationItemDataSourceCrud) mapToConfigurationItemMetadata(fieldKeyFormat string) (oci_opsi.ConfigurationItemMetadata, error) {
	var baseObject oci_opsi.ConfigurationItemMetadata
	//discriminator
	configItemTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_item_type"))
	var configItemType string
	if ok {
		configItemType = configItemTypeRaw.(string)
	} else {
		configItemType = "" // default value
	}
	switch strings.ToLower(configItemType) {
	case strings.ToLower("BASIC"):
		details := oci_opsi.BasicConfigurationItemMetadata{}
		if dataType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_type")); ok {
			tmp := dataType.(string)
			details.DataType = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if unitDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit_details")); ok {
			if tmpList := unitDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "unit_details"), 0)
				tmp, err := s.mapToConfigurationItemUnitDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert unit_details, encountered error: %v", err)
				}
				details.UnitDetails = &tmp
			}
		}
		if valueInputDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_input_details")); ok {
			if tmpList := valueInputDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "value_input_details"), 0)
				tmp, err := s.mapToConfigurationItemAllowedValueDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert value_input_details, encountered error: %v", err)
				}
				details.ValueInputDetails = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_item_type '%v' was specified", configItemType)
	}
	return baseObject, nil
}

// Commented out since same func is present in opsi_opsi_configuration_resource
/* func ConfigurationItemMetadataToMap(obj *oci_opsi.ConfigurationItemMetadata) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_opsi.BasicConfigurationItemMetadata:
		result["config_item_type"] = "BASIC"

		if v.DataType != nil {
			result["data_type"] = string(*v.DataType)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.UnitDetails != nil {
			result["unit_details"] = []interface{}{ConfigurationItemUnitDetailsToMap(v.UnitDetails)}
		}

		if v.ValueInputDetails != nil {
			valueInputDetailsArray := []interface{}{}
			if valueInputDetailsMap := ConfigurationItemAllowedValueDetailsToMap(&v.ValueInputDetails); valueInputDetailsMap != nil {
				valueInputDetailsArray = append(valueInputDetailsArray, valueInputDetailsMap)
			}
			result["value_input_details"] = valueInputDetailsArray
		}
	default:
		log.Printf("[WARN] Received 'config_item_type' of unknown type %v", *obj)
		return nil
	}

	return result
} */

func ConfigurationItemSummaryToMap(obj oci_opsi.ConfigurationItemSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.BasicConfigurationItemSummary:
		result["config_item_type"] = "BASIC"

		result["applicable_contexts"] = v.ApplicableContexts
		result["applicable_contexts"] = v.ApplicableContexts

		if v.DefaultValue != nil {
			result["default_value"] = string(*v.DefaultValue)
		}

		if v.Metadata != nil {
			metadataArray := []interface{}{}
			if metadataMap := ConfigurationItemMetadataToMap(&v.Metadata); metadataMap != nil {
				metadataArray = append(metadataArray, metadataMap)
			}
			result["metadata"] = metadataArray
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}

		result["value_source_config"] = string(v.ValueSourceConfig)
	default:
		log.Printf("[WARN] Received 'config_item_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *OpsiOpsiConfigurationConfigurationItemDataSourceCrud) mapToConfigurationItemUnitDetails(fieldKeyFormat string) (oci_opsi.ConfigurationItemUnitDetails, error) {
	result := oci_opsi.ConfigurationItemUnitDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		tmp := unit.(string)
		result.Unit = &tmp
	}

	return result, nil
}

// Commented out since same func is present in opsi_opsi_configuration_resource
/* func ConfigurationItemUnitDetailsToMap(obj *oci_opsi.ConfigurationItemUnitDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	return result
} */
