// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_dashboard

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_management_dashboard "github.com/oracle/oci-go-sdk/v65/managementdashboard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementDashboardManagementSavedSearchResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagementDashboardManagementSavedSearch,
		Read:     readManagementDashboardManagementSavedSearch,
		Update:   updateManagementDashboardManagementSavedSearch,
		Delete:   deleteManagementDashboardManagementSavedSearch,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_oob_saved_search": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"metadata_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nls": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},
			"provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"screen_image": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ui_config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},
			"widget_template": {
				Type:     schema.TypeString,
				Required: true,
			},
			"widget_vm": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"drilldown_config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeMap,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"parameters_config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func stringToMap(options string) (*interface{}, error) {
	var result interface{}
	var err error

	var obj interface{}
	err = json.Unmarshal([]byte(options), &obj)
	result = &obj

	return &result, err
}

func stringToArrayOfObjects(options string) ([]interface{}, error) {
	var arr []interface{}
	if err := json.Unmarshal([]byte(options), &arr); err != nil {
		return nil, err
	}

	// Optional validation: ensure every element is a JSON object
	for i, v := range arr {
		if _, ok := v.(map[string]interface{}); !ok {
			return nil, fmt.Errorf("element %d is not an object (got %T)", i, v)
		}
	}

	return arr, nil
}

func createManagementDashboardManagementSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()

	return tfresource.CreateResource(d, sync)
}

func readManagementDashboardManagementSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()

	return tfresource.ReadResource(sync)
}

func updateManagementDashboardManagementSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteManagementDashboardManagementSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ManagementDashboardManagementSavedSearchResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_management_dashboard.DashxApisClient
	Res                    *oci_management_dashboard.ManagementSavedSearch
	DisableNotFoundRetries bool
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) ID() string {
	log.Printf("[DEBUG] Resource ID set 2: %s", s.D.Id())
	// log.Printf("[DEBUG] Res: %s", s.Res)
	// log.Printf("[DEBUG] Res.id: %s", *s.Res.Id)

	if s.D != nil {
		return s.D.Id()
	}
	panic("resource ID not found")
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_management_dashboard.LifecycleStatesActive),
	}
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) Create() error {
	request := oci_management_dashboard.CreateManagementSavedSearchRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataConfig, ok := s.D.GetOkExists("data_config"); ok {
		tmp, err := stringToArrayOfObjects(dataConfig.(string))
		if err != nil {
			return err
		}
		request.DataConfig = tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drilldownConfig, ok := s.D.GetOkExists("drilldown_config"); ok {
		tmp, err := stringToArrayOfObjects(drilldownConfig.(string))
		if err != nil {
			return err
		}
		request.DrilldownConfig = tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isOobSavedSearch, ok := s.D.GetOkExists("is_oob_saved_search"); ok {
		tmp := isOobSavedSearch.(bool)
		request.IsOobSavedSearch = &tmp
	}

	if metadataVersion, ok := s.D.GetOkExists("metadata_version"); ok {
		tmp := metadataVersion.(string)
		request.MetadataVersion = &tmp
	}

	if nls, ok := s.D.GetOkExists("nls"); ok {
		tmp, err := stringToMap(nls.(string))
		if err != nil {
			return err
		}
		request.Nls = tmp
	}

	if parametersConfig, ok := s.D.GetOkExists("parameters_config"); ok {
		tmp, err := stringToArrayOfObjects(parametersConfig.(string))
		if err != nil {
			return err
		}
		request.ParametersConfig = tmp
	}

	if providerId, ok := s.D.GetOkExists("provider_id"); ok {
		tmp := providerId.(string)
		request.ProviderId = &tmp
	}

	if providerName, ok := s.D.GetOkExists("provider_name"); ok {
		tmp := providerName.(string)
		request.ProviderName = &tmp
	}

	if providerVersion, ok := s.D.GetOkExists("provider_version"); ok {
		tmp := providerVersion.(string)
		request.ProviderVersion = &tmp
	}

	if screenImage, ok := s.D.GetOkExists("screen_image"); ok {
		tmp := screenImage.(string)
		request.ScreenImage = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_management_dashboard.SavedSearchTypesEnum(type_.(string))
	}

	if uiConfig, ok := s.D.GetOkExists("ui_config"); ok {
		tmp, err := stringToMap(uiConfig.(string))
		if err != nil {
			return err
		}
		request.UiConfig = tmp
	}

	if widgetTemplate, ok := s.D.GetOkExists("widget_template"); ok {
		tmp := widgetTemplate.(string)
		request.WidgetTemplate = &tmp
	}

	if widgetVM, ok := s.D.GetOkExists("widget_vm"); ok {
		tmp := widgetVM.(string)
		request.WidgetVM = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_dashboard")

	response, err := s.Client.CreateManagementSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] CreateManagementSavedSearch response: %+v", response)
	var identifier *string = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
		log.Printf("[DEBUG] Resource ID set: %s", s.D.Id())
	} else {
		log.Printf("[DEBUG] Resource ID is nil")
	}

	return nil
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) Get() error {
	request := oci_management_dashboard.GetManagementSavedSearchRequest{}

	tmp := s.D.Id()
	request.ManagementSavedSearchId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_dashboard")

	response, err := s.Client.GetManagementSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementSavedSearch
	return nil
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_management_dashboard.UpdateManagementSavedSearchRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataConfig, ok := s.D.GetOkExists("data_config"); ok {
		tmp, err := stringToArrayOfObjects(dataConfig.(string))
		if err != nil {
			return err
		}
		request.DataConfig = tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drilldownConfig, ok := s.D.GetOkExists("drilldown_config"); ok {
		tmp, err := stringToArrayOfObjects(drilldownConfig.(string))
		if err != nil {
			return err
		}
		request.DrilldownConfig = tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isOobSavedSearch, ok := s.D.GetOkExists("is_oob_saved_search"); ok {
		tmp := isOobSavedSearch.(bool)
		request.IsOobSavedSearch = &tmp
	}

	tmp := s.D.Id()
	request.ManagementSavedSearchId = &tmp

	if metadataVersion, ok := s.D.GetOkExists("metadata_version"); ok {
		tmp := metadataVersion.(string)
		request.MetadataVersion = &tmp
	}

	if nls, ok := s.D.GetOkExists("nls"); ok {
		tmp, err := stringToMap(nls.(string))
		if err != nil {
			return err
		}
		request.Nls = tmp
	}

	if parametersConfig, ok := s.D.GetOkExists("parameters_config"); ok {
		tmp, err := stringToArrayOfObjects(parametersConfig.(string))
		if err != nil {
			return err
		}
		request.ParametersConfig = tmp
	}

	if providerId, ok := s.D.GetOkExists("provider_id"); ok {
		tmp := providerId.(string)
		request.ProviderId = &tmp
	}

	if providerName, ok := s.D.GetOkExists("provider_name"); ok {
		tmp := providerName.(string)
		request.ProviderName = &tmp
	}

	if providerVersion, ok := s.D.GetOkExists("provider_version"); ok {
		tmp := providerVersion.(string)
		request.ProviderVersion = &tmp
	}

	if screenImage, ok := s.D.GetOkExists("screen_image"); ok {
		tmp := screenImage.(string)
		request.ScreenImage = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_management_dashboard.SavedSearchTypesEnum(type_.(string))
	}

	if uiConfig, ok := s.D.GetOkExists("ui_config"); ok {
		tmp, err := stringToMap(uiConfig.(string))
		if err != nil {
			return err
		}
		request.UiConfig = tmp
	}

	if widgetTemplate, ok := s.D.GetOkExists("widget_template"); ok {
		tmp := widgetTemplate.(string)
		request.WidgetTemplate = &tmp
	}

	if widgetVM, ok := s.D.GetOkExists("widget_vm"); ok {
		tmp := widgetVM.(string)
		request.WidgetVM = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_dashboard")

	response, err := s.Client.UpdateManagementSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementSavedSearch
	return nil
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) Delete() error {
	request := oci_management_dashboard.DeleteManagementSavedSearchRequest{}

	tmp := s.D.Id()
	request.ManagementSavedSearchId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_dashboard")

	_, err := s.Client.DeleteManagementSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) SetData() error {
	if s.Res != nil {
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
			dataJSON, err := json.Marshal(s.Res.DataConfig)
			if err != nil {
				return err
			}
			if err := s.D.Set("data_config", string(dataJSON)); err != nil {
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
			drillJSON, err := json.Marshal(s.Res.DrilldownConfig)
			if err != nil {
				return err
			}
			if err := s.D.Set("drilldown_config", string(drillJSON)); err != nil {
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
			paraJSON, err := json.Marshal(s.Res.ParametersConfig)
			if err != nil {
				return err
			}
			if err := s.D.Set("parameters_config", string(paraJSON)); err != nil {
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
			uiConfigJson, err := json.Marshal(s.Res.UiConfig)
			if err != nil {
				return err
			}
			s.D.Set("ui_config", string(uiConfigJson))
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
	}

	return nil
}

func ManagementSavedSearchSummaryToMap(obj oci_management_dashboard.ManagementSavedSearchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.DataConfig != nil {
		dataJSON, err := json.Marshal(obj.DataConfig)
		if err != nil {
			log.Printf("[DEBUG] Error marshaling obj.DataConfig: %s", err)
		} else {
			result["data_config"] = string(dataJSON)
		}
	} else {
		result["data_config"] = nil
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsOobSavedSearch != nil {
		result["is_oob_saved_search"] = bool(*obj.IsOobSavedSearch)
	}

	if obj.MetadataVersion != nil {
		result["metadata_version"] = string(*obj.MetadataVersion)
	}

	if obj.Nls != nil {
		nlsJson, err := json.Marshal(obj.Nls)
		if err != nil {
			log.Printf("[DEBUG] Error marshaling obj.Nls: %s", err)
		} else {
			result["nls"] = string(nlsJson)
		}
	}

	if obj.ParametersConfig != nil {
		dataJSON, err := json.Marshal(obj.ParametersConfig) // obj.ParametersConfig is []interface{}
		if err != nil {
			log.Printf("[DEBUG] Error marshaling obj.ParametersConfig: %s", err)
		} else {
			result["parameters_config"] = string(dataJSON)
		}
	}

	if obj.ProviderId != nil {
		result["provider_id"] = string(*obj.ProviderId)
	}

	if obj.ProviderName != nil {
		result["provider_name"] = string(*obj.ProviderName)
	}

	if obj.ProviderVersion != nil {
		result["provider_version"] = string(*obj.ProviderVersion)
	}

	if obj.ScreenImage != nil {
		result["screen_image"] = string(*obj.ScreenImage)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	if obj.UiConfig != nil {
		uiConfigJson, err := json.Marshal(obj.UiConfig)
		if err != nil {
			log.Printf("[DEBUG] Error marshaling obj.UiConfig: %s", err)
		} else {
			result["ui_config"] = string(uiConfigJson)
		}
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	if obj.WidgetTemplate != nil {
		result["widget_template"] = string(*obj.WidgetTemplate)
	}

	if obj.WidgetVM != nil {
		result["widget_vm"] = string(*obj.WidgetVM)
	}

	return result
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) mapToobject(fieldKeyFormat string) (interface{}, error) {
	var tmp interface{}
	var err error
	if s.D.Get(fieldKeyFormat) != nil {
		tmp = s.D.Get(fieldKeyFormat)
	} else {
		tmp = make(map[string]interface{})
	}
	return tmp, err
}

func (s *ManagementDashboardManagementSavedSearchResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_management_dashboard.ChangeManagementSavedSearchesCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ManagementSavedSearchId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_dashboard")

	_, err := s.Client.ChangeManagementSavedSearchesCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
