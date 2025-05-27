// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsWlsDomainServerBackupContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWlmsWlsDomainServerBackupContent,
		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"middleware": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"patches": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularWlmsWlsDomainServerBackupContent(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainServerBackupContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainServerBackupContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.GetWlsDomainServerBackupContentResponse
}

func (s *WlmsWlsDomainServerBackupContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainServerBackupContentDataSourceCrud) Get() error {
	request := oci_wlms.GetWlsDomainServerBackupContentRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	if serverId, ok := s.D.GetOkExists("server_id"); ok {
		tmp := serverId.(string)
		request.ServerId = &tmp
	}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.GetWlsDomainServerBackupContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WlmsWlsDomainServerBackupContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainServerBackupContentDataSource-", WlmsWlsDomainServerBackupContentDataSource(), s.D))
	switch v := (s.Res.BackupContent).(type) {
	case oci_wlms.BinaryBackupContent:
		s.D.Set("content_type", "BINARY")

		if v.Middleware != nil {
			s.D.Set("middleware", []interface{}{MiddlewareBinaryBackupContentToMap(v.Middleware)})
		} else {
			s.D.Set("middleware", nil)
		}
	default:
		log.Printf("[WARN] Received 'content_type' of unknown type %v", s.Res.BackupContent)
		return nil
	}

	return nil
}

func (s *WlmsWlsDomainServerBackupContentDataSourceCrud) mapToMiddlewareBackupPatch(fieldKeyFormat string) (oci_wlms.MiddlewareBackupPatch, error) {
	result := oci_wlms.MiddlewareBackupPatch{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func MiddlewareBackupPatchToMap(obj oci_wlms.MiddlewareBackupPatch) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *WlmsWlsDomainServerBackupContentDataSourceCrud) mapToMiddlewareBinaryBackupContent(fieldKeyFormat string) (oci_wlms.MiddlewareBinaryBackupContent, error) {
	result := oci_wlms.MiddlewareBinaryBackupContent{}

	if patches, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patches")); ok {
		interfaces := patches.([]interface{})
		tmp := make([]oci_wlms.MiddlewareBackupPatch, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patches"), stateDataIndex)
			converted, err := s.mapToMiddlewareBackupPatch(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patches")) {
			result.Patches = tmp
		}
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := version.(string)
		result.Version = &tmp
	}

	return result, nil
}

func MiddlewareBinaryBackupContentToMap(obj *oci_wlms.MiddlewareBinaryBackupContent) map[string]interface{} {
	result := map[string]interface{}{}

	patches := []interface{}{}
	for _, item := range obj.Patches {
		patches = append(patches, MiddlewareBackupPatchToMap(item))
	}
	result["patches"] = patches

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
