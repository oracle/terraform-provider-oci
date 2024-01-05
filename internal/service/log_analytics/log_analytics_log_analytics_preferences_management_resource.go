// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsLogAnalyticsPreferencesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsPreferencesManagement,
		Update:   updateLogAnalyticsLogAnalyticsPreferencesManagement,
		Read:     readLogAnalyticsLogAnalyticsPreferencesManagement,
		Delete:   deleteLogAnalyticsLogAnalyticsPreferencesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"items": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      prefItemsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createLogAnalyticsLogAnalyticsPreferencesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func updateLogAnalyticsLogAnalyticsPreferencesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsPreferencesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteLogAnalyticsLogAnalyticsPreferencesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) ID() string {
	return getLogAnalyticsPreferencesManagementId(s.D.Get("namespace").(string))
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) Create() error {
	if items, ok := s.D.GetOkExists("items"); ok {
		interfaces := items.(*schema.Set).List()
		updateItems := make([]oci_log_analytics.LogAnalyticsPreference, len(interfaces))

		for i := range interfaces {
			stateDataIndex := prefItemsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)

			converted, err := s.mapToLogAnalyticsPreference(fieldKeyFormat)
			if err != nil {
				return err
			}
			updateItems[i] = converted
		}

		var ns string
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			ns = namespace.(string)
		}

		err := s.updatePreferences(ns, updateItems)
		if err != nil {
			return fmt.Errorf("failed to update preferences, error: %v", err)
		}
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) Update() error {
	if _, ok := s.D.GetOkExists("items"); ok && s.D.HasChange("items") {
		o, n := s.D.GetChange("items")

		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		oldPreferences := o.(*schema.Set)
		newPreferences := n.(*schema.Set)

		preferencesToRemove := oldPreferences.Difference(newPreferences).List()
		preferencesToUpdate := newPreferences.Difference(oldPreferences).List()

		var ns string
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			ns = namespace.(string)
		}

		if len(preferencesToRemove) != 0 {
			removeItems := make([]oci_log_analytics.LogAnalyticsPreference, len(preferencesToRemove))
			for i := range preferencesToRemove {
				removeItems[i] = mapToLogAnalyticsPreferenceDetailsPreference(preferencesToRemove[i].(map[string]interface{}))
			}
			err := s.removePreferences(ns, removeItems)
			if err != nil {
				return fmt.Errorf("failed to remove preferences, error: %v", err)
			}
		}

		if len(preferencesToUpdate) != 0 {
			updateItems := make([]oci_log_analytics.LogAnalyticsPreference, len(preferencesToUpdate))
			for i := range preferencesToUpdate {
				updateItems[i] = mapToLogAnalyticsPreferenceDetailsPreference(preferencesToUpdate[i].(map[string]interface{}))
			}
			err := s.updatePreferences(ns, updateItems)
			if err != nil {
				return fmt.Errorf("failed to update preferences, error: %v", err)
			}
		}
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) Delete() error {
	if items, ok := s.D.GetOkExists("items"); ok {
		interfaces := items.(*schema.Set).List()
		removeItems := make([]oci_log_analytics.LogAnalyticsPreference, len(interfaces))

		for i := range interfaces {
			stateDataIndex := prefItemsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)

			converted, err := s.mapToLogAnalyticsPreference(fieldKeyFormat)
			if err != nil {
				return err
			}
			removeItems[i] = converted
		}

		var ns string
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			ns = namespace.(string)
		}

		err := s.removePreferences(ns, removeItems)
		if err != nil {
			return fmt.Errorf("failed to remove preferences, error: %v", err)
		}
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) updatePreferences(namespace string, items []oci_log_analytics.LogAnalyticsPreference) error {
	request := oci_log_analytics.UpdatePreferencesRequest{}
	request.NamespaceName = &namespace
	request.UpdatePreferencesDetails.Items = items
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")
	_, err := s.Client.UpdatePreferences(context.Background(), request)

	if err != nil {
		return err
	}
	return nil
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) removePreferences(namespace string, items []oci_log_analytics.LogAnalyticsPreference) error {
	request := oci_log_analytics.RemovePreferencesRequest{}
	request.NamespaceName = &namespace
	request.RemovePreferencesDetails.Items = items
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")
	_, err := s.Client.RemovePreferences(context.Background(), request)

	if err != nil {
		return err
	}
	return nil
}

func getLogAnalyticsPreferencesManagementId(namespace string) string {
	namespace = url.PathEscape(namespace)
	preferencesManagementId := "namespaces/" + namespace + "/preferences"
	return preferencesManagementId
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) SetData() error {
	if items, ok := s.D.GetOkExists("items"); ok {
		interfaces := items.(*schema.Set).List()

		checkItems := []interface{}{}
		for _, item := range interfaces {
			checkItems = append(checkItems, item.(map[string]interface{}))
		}

		s.D.Set("items", schema.NewSet(prefItemsHashCodeForSets, checkItems))
	}

	return nil
}

func prefItemsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if preferenceName, ok := m["name"]; ok && preferenceName != "" {
		buf.WriteString(fmt.Sprintf("%v-", preferenceName))
	}

	if preferenceValue, ok := m["value"]; ok && preferenceValue != "" {
		buf.WriteString(fmt.Sprintf("%v-", preferenceValue))
	}

	return utils.GetStringHashcode(buf.String())
}

func mapToLogAnalyticsPreferenceDetailsPreference(preference map[string]interface{}) oci_log_analytics.LogAnalyticsPreference {
	result := oci_log_analytics.LogAnalyticsPreference{}

	if preferenceName, ok := preference["name"]; ok && preferenceName != "" {
		tmp := preferenceName.(string)
		result.Name = &tmp
	}

	if preferenceValue, ok := preference["value"]; ok && preferenceValue != "" {
		tmp := preferenceValue.(string)
		result.Value = &tmp
	}

	return result
}

func (s *LogAnalyticsLogAnalyticsPreferencesManagementResourceCrud) mapToLogAnalyticsPreference(fieldKeyFormat string) (oci_log_analytics.LogAnalyticsPreference, error) {
	result := oci_log_analytics.LogAnalyticsPreference{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}
