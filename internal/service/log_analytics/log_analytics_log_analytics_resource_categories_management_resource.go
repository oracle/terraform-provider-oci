// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"fmt"
	"net/url"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v56/loganalytics"
)

func LogAnalyticsLogAnalyticsResourceCategoriesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsResourceCategoriesManagement,
		Update:   updateLogAnalyticsLogAnalyticsResourceCategoriesManagement,
		Read:     readLogAnalyticsLogAnalyticsResourceCategoriesManagement,
		Delete:   deleteLogAnalyticsLogAnalyticsResourceCategoriesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"resource_categories": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional

			// Computed
		},
	}
}

func createLogAnalyticsLogAnalyticsResourceCategoriesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func updateLogAnalyticsLogAnalyticsResourceCategoriesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsResourceCategoriesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteLogAnalyticsLogAnalyticsResourceCategoriesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud) ID() string {
	return getLogAnalyticsResourceCategoriesManagementId(
		s.D.Get("namespace").(string), s.D.Get("resource_id").(string), s.D.Get("resource_type").(string))
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud) Create() error {
	if resourceCategories, ok := s.D.GetOkExists("resource_categories"); ok {
		interfaces := resourceCategories.(*schema.Set).List()
		updateItems := make([]oci_log_analytics.LogAnalyticsResourceCategory, len(interfaces))

		var ns string
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			ns = namespace.(string)
		}

		var resourceId string
		if resId, ok := s.D.GetOkExists("resource_id"); ok {
			resourceId = resId.(string)
		}

		var resourceType string
		if resType, ok := s.D.GetOkExists("resource_type"); ok {
			resourceType = resType.(string)
		}

		for i := range interfaces {
			updateItems[i] = toLogAnalyticsResourceCategory(resourceId, resourceType, interfaces[i].(string))
		}

		err := s.updateResourceCategories(ns, updateItems)
		if err != nil {
			return fmt.Errorf("failed to update resource categories, error: %v", err)
		}
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud) Update() error {
	if _, ok := s.D.GetOkExists("resource_categories"); ok && s.D.HasChange("resource_categories") {
		o, n := s.D.GetChange("resource_categories")

		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		oldResourceCategories := o.(*schema.Set)
		newResourceCategories := n.(*schema.Set)

		resourceCategoriesToRemove := oldResourceCategories.Difference(newResourceCategories).List()
		resourceCategoriesToUpdate := newResourceCategories.Difference(oldResourceCategories).List()

		var ns string
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			ns = namespace.(string)
		}

		var resourceId string
		if resId, ok := s.D.GetOkExists("resource_id"); ok {
			resourceId = resId.(string)
		}

		var resourceType string
		if resType, ok := s.D.GetOkExists("resource_type"); ok {
			resourceType = resType.(string)
		}

		if len(resourceCategoriesToRemove) != 0 {
			removeItems := make([]oci_log_analytics.LogAnalyticsResourceCategory, len(resourceCategoriesToRemove))
			for i := range resourceCategoriesToRemove {
				removeItems[i] = toLogAnalyticsResourceCategory(resourceId, resourceType, resourceCategoriesToRemove[i].(string))
			}
			err := s.removeResourceCategories(ns, removeItems)
			if err != nil {
				return fmt.Errorf("failed to remove resource categories, error: %v", err)
			}
		}

		if len(resourceCategoriesToUpdate) != 0 {
			updateItems := make([]oci_log_analytics.LogAnalyticsResourceCategory, len(resourceCategoriesToUpdate))
			for i := range resourceCategoriesToUpdate {
				updateItems[i] = toLogAnalyticsResourceCategory(resourceId, resourceType, resourceCategoriesToUpdate[i].(string))
			}
			err := s.updateResourceCategories(ns, updateItems)
			if err != nil {
				return fmt.Errorf("failed to update resource categories, error: %v", err)
			}
		}
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud) Delete() error {
	if resourceCategories, ok := s.D.GetOkExists("resource_categories"); ok {
		interfaces := resourceCategories.(*schema.Set).List()
		removeItems := make([]oci_log_analytics.LogAnalyticsResourceCategory, len(interfaces))

		var ns string
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			ns = namespace.(string)
		}

		var resourceId string
		if resId, ok := s.D.GetOkExists("resource_id"); ok {
			resourceId = resId.(string)
		}

		var resourceType string
		if resType, ok := s.D.GetOkExists("resource_type"); ok {
			resourceType = resType.(string)
		}

		for i := range interfaces {
			removeItems[i] = toLogAnalyticsResourceCategory(resourceId, resourceType, interfaces[i].(string))
		}

		err := s.removeResourceCategories(ns, removeItems)
		if err != nil {
			return fmt.Errorf("failed to remove resource categories, error: %v", err)
		}
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud) updateResourceCategories(namespace string, items []oci_log_analytics.LogAnalyticsResourceCategory) error {
	request := oci_log_analytics.UpdateResourceCategoriesRequest{}
	request.NamespaceName = &namespace
	request.UpdateResourceCategoriesDetails.Items = items
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")
	_, err := s.Client.UpdateResourceCategories(context.Background(), request)

	if err != nil {
		return err
	}
	return nil
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud) removeResourceCategories(namespace string, items []oci_log_analytics.LogAnalyticsResourceCategory) error {
	request := oci_log_analytics.RemoveResourceCategoriesRequest{}
	request.NamespaceName = &namespace
	request.RemoveResourceCategoriesDetails.Items = items
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")
	_, err := s.Client.RemoveResourceCategories(context.Background(), request)

	if err != nil {
		return err
	}
	return nil
}

func getLogAnalyticsResourceCategoriesManagementId(namespace string, resourceId string, resourceType string) string {
	namespace = url.PathEscape(namespace)
	resourceId = url.PathEscape(resourceId)
	resourceType = url.PathEscape(resourceType)
	resourceCategoriesManagementId := "namespaces/" + namespace + "/" + resourceId + "/" + resourceType + "/categories"
	return resourceCategoriesManagementId
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceCrud) SetData() error {
	if resourceCategories, ok := s.D.GetOkExists("resource_categories"); ok {
		interfaces := resourceCategories.(*schema.Set).List()

		checkItems := []interface{}{}
		for _, item := range interfaces {
			checkItems = append(checkItems, item)
		}

		s.D.Set("resource_categories", schema.NewSet(utils.LiteralTypeHashCodeForSets, checkItems))
	}

	return nil
}

func toLogAnalyticsResourceCategory(resourceId string, resourceType string, categoryName string) oci_log_analytics.LogAnalyticsResourceCategory {
	result := oci_log_analytics.LogAnalyticsResourceCategory{}

	result.ResourceId = &resourceId
	result.ResourceType = &resourceType
	result.CategoryName = &categoryName

	return result
}
