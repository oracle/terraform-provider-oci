// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiOpsiConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["config_item_custom_status"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["config_item_field"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["config_items_applicable_context"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["opsi_config_field"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["opsi_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiOpsiConfigurationResource(), fieldMap, readSingularOpsiOpsiConfiguration)
}

func readSingularOpsiOpsiConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOpsiConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOpsiConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetOpsiConfigurationResponse
}

func (s *OpsiOpsiConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOpsiConfigurationDataSourceCrud) Get() error {
	request := oci_opsi.GetOpsiConfigurationRequest{}

	if configItemCustomStatus, ok := s.D.GetOkExists("config_item_custom_status"); ok {
		interfaces := configItemCustomStatus.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.GetOpsiConfigurationConfigItemCustomStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingGetOpsiConfigurationConfigItemCustomStatusEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp2) != 0 || s.D.HasChange("config_item_custom_status") {
			request.ConfigItemCustomStatus = tmp2
		}
	}

	if configItemField, ok := s.D.GetOkExists("config_item_field"); ok {
		interfaces := configItemField.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.GetOpsiConfigurationConfigItemFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingGetOpsiConfigurationConfigItemFieldEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp2) != 0 || s.D.HasChange("config_item_field") {
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

	if opsiConfigField, ok := s.D.GetOkExists("opsi_config_field"); ok {
		interfaces := opsiConfigField.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.GetOpsiConfigurationOpsiConfigFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingGetOpsiConfigurationOpsiConfigFieldEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp2) != 0 || s.D.HasChange("opsi_config_field") {
			request.OpsiConfigField = tmp2
		}
	}

	if opsiConfigurationId, ok := s.D.GetOkExists("opsi_configuration_id"); ok {
		tmp := opsiConfigurationId.(string)
		request.OpsiConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetOpsiConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiOpsiConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.OpsiConfiguration.GetId())
	switch v := (s.Res.OpsiConfiguration).(type) {
	case oci_opsi.OpsiUxConfiguration:
		s.D.Set("opsi_config_type", "UX_CONFIGURATION")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ConfigItems != nil {
			configItems := []interface{}{}
			for _, item := range v.ConfigItems {
				configItems = append(configItems, OpsiConfigurationConfigurationItemSummaryToMap(item))
			}
			s.D.Set("config_items", configItems)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.FreeformTags != nil {
			s.D.Set("freeform_tags", v.FreeformTags)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}
	default:
		log.Printf("[WARN] Received 'opsi_config_type' of unknown type %v", s.Res.OpsiConfiguration)
		return nil
	}

	return nil
}
