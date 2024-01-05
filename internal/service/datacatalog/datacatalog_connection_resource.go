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

func DatacatalogConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatacatalogConnection,
		Read:     readDatacatalogConnection,
		Update:   updateDatacatalogConnection,
		Delete:   deleteDatacatalogConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"catalog_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_asset_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": {
				Type:             schema.TypeMap,
				Required:         true,
				DiffSuppressFunc: propertiesDiffSuppressFunction,
				Elem:             schema.TypeString,
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
			"enc_properties": {
				Type:             schema.TypeMap,
				Optional:         true,
				Sensitive:        true,
				DiffSuppressFunc: encPropertiesDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_status_updated": {
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

func createDatacatalogConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.CreateResource(d, sync)
}

func readDatacatalogConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

func updateDatacatalogConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatacatalogConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatacatalogConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datacatalog.DataCatalogClient
	Res                    *oci_datacatalog.Connection
	DisableNotFoundRetries bool
}

func (s *DatacatalogConnectionResourceCrud) ID() string {
	return *s.Res.Key
}

func (s *DatacatalogConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateCreating),
	}
}

func (s *DatacatalogConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateActive),
	}
}

func (s *DatacatalogConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleting),
	}
}

func (s *DatacatalogConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleted),
	}
}

func (s *DatacatalogConnectionResourceCrud) Create() error {
	request := oci_datacatalog.CreateConnectionRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if dataAssetKey, ok := s.D.GetOkExists("data_asset_key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if encProperties, ok := s.D.GetOkExists("enc_properties"); ok {
		convertedProperties, err := mapToProperties(encProperties.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.EncProperties = convertedProperties
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
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

	response, err := s.Client.CreateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DatacatalogConnectionResourceCrud) Get() error {
	request := oci_datacatalog.GetConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionKey = &tmp

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if dataAssetKey, ok := s.D.GetOkExists("data_asset_key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]oci_datacatalog.GetConnectionFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacatalog.GetConnectionFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	catalogId, connectionKey, dataAssetKey, err := parseConnectionCompositeId(s.D.Id())
	if err == nil {
		request.CatalogId = &catalogId
		request.ConnectionKey = &connectionKey
		request.DataAssetKey = &dataAssetKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DatacatalogConnectionResourceCrud) Update() error {
	request := oci_datacatalog.UpdateConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionKey = &tmp

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if dataAssetKey, ok := s.D.GetOkExists("data_asset_key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if encProperties, ok := s.D.GetOkExists("enc_properties"); ok {
		convertedProperties, err := mapToProperties(encProperties.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.EncProperties = convertedProperties
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		convertedProperties, err := mapToProperties(properties.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.Properties = convertedProperties
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.UpdateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DatacatalogConnectionResourceCrud) Delete() error {
	request := oci_datacatalog.DeleteConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionKey = &tmp

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if connectionKey, ok := s.D.GetOkExists("key"); ok {
		tmp := connectionKey.(string)
		request.ConnectionKey = &tmp
	}

	if dataAssetKey, ok := s.D.GetOkExists("data_asset_key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	_, err := s.Client.DeleteConnection(context.Background(), request)
	return err
}

func (s *DatacatalogConnectionResourceCrud) SetData() error {

	catalogId, connectionKey, dataAssetKey, err := parseConnectionCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("catalog_id", &catalogId)
		s.D.Set("data_asset_key", &dataAssetKey)
		s.D.SetId(connectionKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CreatedById != nil {
		s.D.Set("created_by_id", *s.Res.CreatedById)
	}

	if s.Res.DataAssetKey != nil {
		s.D.Set("data_asset_key", *s.Res.DataAssetKey)
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

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Properties != nil {
		s.D.Set("properties", propertiesToMap(s.Res.Properties))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeStatusUpdated != nil {
		s.D.Set("time_status_updated", s.Res.TimeStatusUpdated.String())
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

func GetConnectionCompositeId(catalogId string, connectionKey string, dataAssetKey string) string {
	catalogId = url.PathEscape(catalogId)
	connectionKey = url.PathEscape(connectionKey)
	dataAssetKey = url.PathEscape(dataAssetKey)
	compositeId := "catalogs/" + catalogId + "/dataAssets/" + dataAssetKey + "/connections/" + connectionKey
	return compositeId
}

func parseConnectionCompositeId(compositeId string) (catalogId string, connectionKey string, dataAssetKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("catalogs/.*/dataAssets/.*/connections/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	catalogId, _ = url.PathUnescape(parts[1])
	dataAssetKey, _ = url.PathUnescape(parts[3])
	connectionKey, _ = url.PathUnescape(parts[5])

	return
}

func ConnectionSummaryToMap(obj oci_datacatalog.ConnectionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataAssetKey != nil {
		result["data_asset_key"] = string(*obj.DataAssetKey)
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

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
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
