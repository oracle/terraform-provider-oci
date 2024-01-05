// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"
)

func DatacatalogDataAssetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatacatalogDataAsset,
		Read:     readDatacatalogDataAsset,
		Update:   updateDatacatalogDataAsset,
		Delete:   deleteDatacatalogDataAsset,
		Schema: map[string]*schema.Schema{
			// Required
			"catalog_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"properties": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: propertiesDiffSuppressFunction,
				Elem:             schema.TypeString,
			},

			// Computed
			"created_by_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_harvested": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_by_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatacatalogDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.CreateResource(d, sync)
}

func readDatacatalogDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

func updateDatacatalogDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatacatalogDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatacatalogDataAssetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datacatalog.DataCatalogClient
	Res                    *oci_datacatalog.DataAsset
	DisableNotFoundRetries bool
}

func (s *DatacatalogDataAssetResourceCrud) ID() string {
	return *s.Res.Key
}

func (s *DatacatalogDataAssetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateCreating),
	}
}

func (s *DatacatalogDataAssetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateActive),
	}
}

func (s *DatacatalogDataAssetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleting),
	}
}

func (s *DatacatalogDataAssetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleted),
	}
}

func (s *DatacatalogDataAssetResourceCrud) Create() error {
	request := oci_datacatalog.CreateDataAssetRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		convertedProperties, err := mapToProperties(properties.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.Properties = convertedProperties
	}

	if typeKey, ok := s.D.GetOkExists("type_key"); ok {
		tmp := typeKey.(string)
		request.TypeKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.CreateDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *DatacatalogDataAssetResourceCrud) Get() error {
	request := oci_datacatalog.GetDataAssetRequest{}
	tmp := s.D.Id()
	request.DataAssetKey = &tmp
	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]oci_datacatalog.GetDataAssetFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacatalog.GetDataAssetFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	catalogId, dataAssetKey, err := parseDataAssetCompositeId(s.D.Id())
	if err == nil {
		request.CatalogId = &catalogId
		request.DataAssetKey = &dataAssetKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.GetDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *DatacatalogDataAssetResourceCrud) Update() error {
	request := oci_datacatalog.UpdateDataAssetRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	tmp := s.D.Id()
	request.DataAssetKey = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		convertedProperties, err := mapToProperties(properties.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.Properties = convertedProperties
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.UpdateDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *DatacatalogDataAssetResourceCrud) Delete() error {
	request := oci_datacatalog.DeleteDataAssetRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if dataAssetKey, ok := s.D.GetOkExists("key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	_, err := s.Client.DeleteDataAsset(context.Background(), request)
	return err
}

func (s *DatacatalogDataAssetResourceCrud) SetData() error {

	catalogId, dataAssetKey, err := parseDataAssetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("catalog_id", &catalogId)
		s.D.SetId(dataAssetKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CatalogId != nil {
		s.D.Set("catalog_id", *s.Res.CatalogId)
	}

	if s.Res.CreatedById != nil {
		s.D.Set("created_by_id", *s.Res.CreatedById)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalKey != nil {
		s.D.Set("external_key", *s.Res.ExternalKey)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Properties != nil {
		s.D.Set("properties", propertiesToMap(s.Res.Properties))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeHarvested == nil {
		s.D.Set("time_harvested", "null")
	}

	if s.Res.TimeHarvested != nil {
		s.D.Set("time_harvested", s.Res.TimeHarvested.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TypeKey != nil {
		s.D.Set("type_key", *s.Res.TypeKey)
	}

	if s.Res.UpdatedById != nil {
		s.D.Set("updated_by_id", *s.Res.UpdatedById)
	}

	if s.Res.Uri != nil {
		s.D.Set("uri", *s.Res.Uri)
	}

	return nil
}

func GetDataAssetCompositeId(catalogId string, dataAssetKey string) string {
	catalogId = url.PathEscape(catalogId)
	dataAssetKey = url.PathEscape(dataAssetKey)
	compositeId := "catalogs/" + catalogId + "/dataAssets/" + dataAssetKey
	return compositeId
}

func parseDataAssetCompositeId(compositeId string) (catalogId string, dataAssetKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("catalogs/.*/dataAssets/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	catalogId, _ = url.PathUnescape(parts[1])
	dataAssetKey, _ = url.PathUnescape(parts[3])

	return
}

func DataAssetSummaryToMap(obj oci_datacatalog.DataAssetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CatalogId != nil {
		result["catalog_id"] = string(*obj.CatalogId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalKey != nil {
		result["external_key"] = string(*obj.ExternalKey)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TypeKey != nil {
		result["type_key"] = string(*obj.TypeKey)
	}

	if obj.Uri != nil {
		result["uri"] = string(*obj.Uri)
	}

	return result
}

func propertiesToMap(properties map[string]map[string]string) map[string]interface{} {
	var rtn = make(map[string]interface{})
	if len(properties) > 0 {
		for namespace, keys := range properties {
			for key, value := range keys {
				rtn[namespace+"."+key] = value
			}
		}
	}
	return rtn
}

func mapToProperties(rawMap map[string]interface{}) (map[string]map[string]string, error) {
	properties := map[string]map[string]string{}
	if len(rawMap) > 0 {
		for key, value := range rawMap {
			var keyComponents = strings.Split(key, ".")
			if len(keyComponents) != 2 {
				return nil, fmt.Errorf("invalid key structure found %s", key)
			}
			var namespace = keyComponents[0]
			if _, ok := properties[namespace]; !ok {
				properties[namespace] = map[string]string{}
			}
			properties[namespace][keyComponents[1]] = value.(string)
		}
	}
	return properties, nil
}

//
//func getDataAssetCompositeId(dataAssetKey string, catalogId string) string {
//	dataAssetKey = url.PathEscape(dataAssetKey)
//	catalogId = url.PathEscape(catalogId)
//	compositeId := "catalogs/" + catalogId + "/dataAssets/" + dataAssetKey
//	return compositeId
//}
//
//func parseDataAssetCompositeId(compositeId string) (dataAssetKey string, catalogId string, err error) {
//	parts := strings.Split(compositeId, "/")
//	match, _ := regexp.MatchString("catalogs/.*/dataAssets/.*", compositeId)
//	if !match || len(parts) != 4 {
//		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
//		return
//	}
//	catalogId, _ = url.PathUnescape(parts[1])
//	dataAssetKey, _ = url.PathUnescape(parts[3])
//
//	return
//}
