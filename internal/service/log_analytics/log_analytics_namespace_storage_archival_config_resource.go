// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceStorageArchivalConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsNamespaceStorageArchivalConfig,
		Read:     readLogAnalyticsNamespaceStorageArchivalConfig,
		Update:   updateLogAnalyticsNamespaceStorageArchivalConfig,
		Delete:   deleteLogAnalyticsNamespaceStorageArchivalConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"archiving_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"active_storage_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"archival_storage_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Computed
			"is_archiving_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsNamespaceStorageArchivalConfig(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageArchivalConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	return nil

}

func readLogAnalyticsNamespaceStorageArchivalConfig(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageArchivalConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsNamespaceStorageArchivalConfig(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageArchivalConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteLogAnalyticsNamespaceStorageArchivalConfig(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LogAnalyticsNamespaceStorageArchivalConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.Storage
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigResourceCrud) ID() string {
	return GetNamespaceStorageArchivalConfigCompositeId(s.D.Get("namespace").(string))
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigResourceCrud) Create() error {
	request := oci_log_analytics.UpdateStorageRequest{}

	if archivingConfiguration, ok := s.D.GetOkExists("archiving_configuration"); ok {
		if tmpList := archivingConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "archiving_configuration", 0)
			tmp, err := s.mapToArchivingConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ArchivingConfiguration = &tmp
		}
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateStorage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Storage
	return nil
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigResourceCrud) Get() error {
	request := oci_log_analytics.GetStorageRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	namespace, err := parseNamespaceStorageArchivalConfigCompositeId(s.D.Id())
	if err == nil {
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetStorage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Storage
	return nil
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigResourceCrud) Update() error {
	request := oci_log_analytics.UpdateStorageRequest{}

	if archivingConfiguration, ok := s.D.GetOkExists("archiving_configuration"); ok {
		if tmpList := archivingConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "archiving_configuration", 0)
			tmp, err := s.mapToArchivingConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ArchivingConfiguration = &tmp
		}
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateStorage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Storage
	return nil
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigResourceCrud) SetData() error {

	namespace, err := parseNamespaceStorageArchivalConfigCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ArchivingConfiguration != nil {
		s.D.Set("archiving_configuration", []interface{}{ArchivingConfigurationToMap(s.Res.ArchivingConfiguration)})
	} else {
		s.D.Set("archiving_configuration", nil)
	}

	if s.Res.IsArchivingEnabled != nil {
		s.D.Set("is_archiving_enabled", *s.Res.IsArchivingEnabled)
	}

	return nil
}

func GetNamespaceStorageArchivalConfigCompositeId(namespace string) string {
	namespace = url.PathEscape(namespace)
	compositeId := "namespaces/" + namespace + "/storage"
	return compositeId
}

func parseNamespaceStorageArchivalConfigCompositeId(compositeId string) (namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/storage", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])

	return
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigResourceCrud) mapToArchivingConfiguration(fieldKeyFormat string) (oci_log_analytics.ArchivingConfiguration, error) {
	result := oci_log_analytics.ArchivingConfiguration{}

	if activeStorageDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "active_storage_duration")); ok {
		tmp := activeStorageDuration.(string)
		result.ActiveStorageDuration = &tmp
	}

	if archivalStorageDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archival_storage_duration")); ok {
		tmp := archivalStorageDuration.(string)
		result.ArchivalStorageDuration = &tmp
	}

	return result, nil
}

func ArchivingConfigurationToMap(obj *oci_log_analytics.ArchivingConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveStorageDuration != nil {
		result["active_storage_duration"] = string(*obj.ActiveStorageDuration)
	}

	if obj.ArchivalStorageDuration != nil {
		result["archival_storage_duration"] = string(*obj.ArchivalStorageDuration)
	}

	return result
}
