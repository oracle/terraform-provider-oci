// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/v58/filestorage"
)

func FileStorageExportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageExport,
		Read:     readFileStorageExport,
		Update:   updateFileStorageExport,
		Delete:   deleteFileStorageExport,
		Schema: map[string]*schema.Schema{
			// Required
			"export_set_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"export_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"access": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"anonymous_gid": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},
						"anonymous_uid": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},
						"identity_squash": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"require_privileged_source_port": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFileStorageExport(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readFileStorageExport(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageExport(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFileStorageExport(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageExportResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.Export
	DisableNotFoundRetries bool
}

func (s *FileStorageExportResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileStorageExportResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateCreating),
	}
}

func (s *FileStorageExportResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateActive),
	}
}

func (s *FileStorageExportResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateDeleting),
	}
}

func (s *FileStorageExportResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateDeleted),
	}
}

func (s *FileStorageExportResourceCrud) Create() error {
	request := oci_file_storage.CreateExportRequest{}

	if exportOptions, ok := s.D.GetOkExists("export_options"); ok {
		interfaces := exportOptions.([]interface{})
		tmp := make([]oci_file_storage.ClientOptions, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "export_options", stateDataIndex)
			converted, err := s.mapToClientOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("export_options") {
			request.ExportOptions = tmp
		}
	}

	if exportSetId, ok := s.D.GetOkExists("export_set_id"); ok {
		tmp := exportSetId.(string)
		request.ExportSetId = &tmp
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if path, ok := s.D.GetOkExists("path"); ok {
		tmp := path.(string)
		request.Path = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageExportResourceCrud) Get() error {
	request := oci_file_storage.GetExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	return nil
}

func (s *FileStorageExportResourceCrud) Update() error {
	request := oci_file_storage.UpdateExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	if exportOptions, ok := s.D.GetOkExists("export_options"); ok {
		interfaces := exportOptions.([]interface{})
		tmp := make([]oci_file_storage.ClientOptions, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "export_options", stateDataIndex)
			converted, err := s.mapToClientOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("export_options") {
			request.ExportOptions = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	return nil
}

func (s *FileStorageExportResourceCrud) Delete() error {
	request := oci_file_storage.DeleteExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteExport(context.Background(), request)
	return err
}

func (s *FileStorageExportResourceCrud) SetData() error {
	exportOptions := []interface{}{}
	for _, item := range s.Res.ExportOptions {
		exportOptions = append(exportOptions, ClientOptionsToMap(item))
	}
	s.D.Set("export_options", exportOptions)

	if s.Res.ExportSetId != nil {
		s.D.Set("export_set_id", *s.Res.ExportSetId)
	}

	if s.Res.FileSystemId != nil {
		s.D.Set("file_system_id", *s.Res.FileSystemId)
	}

	if s.Res.Path != nil {
		s.D.Set("path", *s.Res.Path)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *FileStorageExportResourceCrud) mapToClientOptions(fieldKeyFormat string) (oci_file_storage.ClientOptions, error) {
	result := oci_file_storage.ClientOptions{}

	if access, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access")); ok {
		result.Access = oci_file_storage.ClientOptionsAccessEnum(access.(string))
	}

	if anonymousGid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "anonymous_gid")); ok {
		tmp := anonymousGid.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert anonymousGid string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.AnonymousGid = &tmpInt64
	}

	if anonymousUid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "anonymous_uid")); ok {
		tmp := anonymousUid.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert anonymousUid string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.AnonymousUid = &tmpInt64
	}

	if identitySquash, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identity_squash")); ok {
		result.IdentitySquash = oci_file_storage.ClientOptionsIdentitySquashEnum(identitySquash.(string))
	}

	if requirePrivilegedSourcePort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "require_privileged_source_port")); ok {
		tmp := requirePrivilegedSourcePort.(bool)
		result.RequirePrivilegedSourcePort = &tmp
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		tmp := source.(string)
		result.Source = &tmp
	}

	return result, nil
}

func ClientOptionsToMap(obj oci_file_storage.ClientOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["access"] = string(obj.Access)

	if obj.AnonymousGid != nil {
		result["anonymous_gid"] = strconv.FormatInt(*obj.AnonymousGid, 10)
	}

	if obj.AnonymousUid != nil {
		result["anonymous_uid"] = strconv.FormatInt(*obj.AnonymousUid, 10)
	}

	result["identity_squash"] = string(obj.IdentitySquash)

	if obj.RequirePrivilegedSourcePort != nil {
		result["require_privileged_source_port"] = bool(*obj.RequirePrivilegedSourcePort)
	}

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	return result
}
