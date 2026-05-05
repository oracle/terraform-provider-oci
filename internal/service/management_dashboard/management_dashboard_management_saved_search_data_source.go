// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_dashboard

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_dashboard "github.com/oracle/oci-go-sdk/v65/managementdashboard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementDashboardManagementSavedSearchDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["management_saved_search_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ManagementDashboardManagementSavedSearchResource(), fieldMap, readSingularManagementDashboardManagementSavedSearch)
}

func readSingularManagementDashboardManagementSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementSavedSearchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()

	return tfresource.ReadResource(sync)
}

type ManagementDashboardManagementSavedSearchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_dashboard.DashxApisClient
	Res    *oci_management_dashboard.GetManagementSavedSearchResponse
}

func (s *ManagementDashboardManagementSavedSearchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementDashboardManagementSavedSearchDataSourceCrud) Get() error {
	request := oci_management_dashboard.GetManagementSavedSearchRequest{}

	if managementSavedSearchId, ok := s.D.GetOkExists("management_saved_search_id"); ok {
		tmp := managementSavedSearchId.(string)
		request.ManagementSavedSearchId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_dashboard")

	response, err := s.Client.GetManagementSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementDashboardManagementSavedSearchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DataConfig == nil {
		if err := s.D.Set("data_config", nil); err != nil {
			return err
		}
	} else {
		b, err := json.Marshal(s.Res.DataConfig) // s.Res.DataConfig is []interface{}
		if err != nil {
			return err
		}
		if err := s.D.Set("data_config", string(b)); err != nil {
			return err
		}
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrilldownConfig == nil {
		if err := s.D.Set("drilldown_config", nil); err != nil {
			return err
		}
	} else {
		b, err := json.Marshal(s.Res.DrilldownConfig) // s.Res.DataConfig is []interface{}
		if err != nil {
			return err
		}
		if err := s.D.Set("drilldown_config", string(b)); err != nil {
			return err
		}
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOobSavedSearch != nil {
		s.D.Set("is_oob_saved_search", *s.Res.IsOobSavedSearch)
	}

	if s.Res.MetadataVersion != nil {
		s.D.Set("metadata_version", *s.Res.MetadataVersion)
	}

	if s.Res.Nls != nil {
		nlsJson, err := json.Marshal(s.Res.Nls)
		if err != nil {
			return err
		}
		s.D.Set("nls", string(nlsJson))
	} else {
		s.D.Set("nls", nil)
	}

	if s.Res.ParametersConfig == nil {
		if err := s.D.Set("parameters_config", nil); err != nil {
			return err
		}
	} else {
		b, err := json.Marshal(s.Res.ParametersConfig) // s.Res.DataConfig is []interface{}
		if err != nil {
			return err
		}
		if err := s.D.Set("parameters_config", string(b)); err != nil {
			return err
		}
	}

	if s.Res.ProviderId != nil {
		s.D.Set("provider_id", *s.Res.ProviderId)
	}

	if s.Res.ProviderName != nil {
		s.D.Set("provider_name", *s.Res.ProviderName)
	}

	if s.Res.ProviderVersion != nil {
		s.D.Set("provider_version", *s.Res.ProviderVersion)
	}

	if s.Res.ScreenImage != nil {
		s.D.Set("screen_image", *s.Res.ScreenImage)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UiConfig != nil {
		uiJson, err := json.Marshal(s.Res.UiConfig)
		if err != nil {
			return err
		}
		s.D.Set("ui_config", string(uiJson))
	} else {
		s.D.Set("ui_config", nil)
	}

	if s.Res.UpdatedBy != nil {
		s.D.Set("updated_by", *s.Res.UpdatedBy)
	}

	if s.Res.WidgetTemplate != nil {
		s.D.Set("widget_template", *s.Res.WidgetTemplate)
	}

	if s.Res.WidgetVM != nil {
		s.D.Set("widget_vm", *s.Res.WidgetVM)
	}

	return nil
}
