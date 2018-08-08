// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func ExportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createExport,
		Read:     readExport,
		Update:   updateExport,
		Delete:   deleteExport,
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
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},
						"anonymous_uid": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
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

func createExport(d *schema.ResourceData, m interface{}) error {
	sync := &ExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return CreateResource(d, sync)
}

func readExport(d *schema.ResourceData, m interface{}) error {
	sync := &ExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return ReadResource(sync)
}

func updateExport(d *schema.ResourceData, m interface{}) error {
	sync := &ExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return UpdateResource(d, sync)
}

func deleteExport(d *schema.ResourceData, m interface{}) error {
	sync := &ExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ExportResourceCrud struct {
	BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.Export
	DisableNotFoundRetries bool
}

func (s *ExportResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ExportResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateCreating),
	}
}

func (s *ExportResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateActive),
	}
}

func (s *ExportResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateDeleting),
	}
}

func (s *ExportResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateDeleted),
	}
}

func (s *ExportResourceCrud) Create() error {
	request := oci_file_storage.CreateExportRequest{}

	request.ExportOptions = []oci_file_storage.ClientOptions{}
	if exportOptions, ok := s.D.GetOkExists("export_options"); ok {
		interfaces := exportOptions.([]interface{})
		tmp := make([]oci_file_storage.ClientOptions, len(interfaces))
		for i := range interfaces {
			fieldKeyPrefix := fmt.Sprintf("export_options.%d", i)
			converted, err := mapToClientOptions(fieldKeyPrefix+".%s", s.D)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.ExportOptions = tmp
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	return nil
}

func (s *ExportResourceCrud) Get() error {
	request := oci_file_storage.GetExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	return nil
}

func (s *ExportResourceCrud) Update() error {
	request := oci_file_storage.UpdateExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	request.ExportOptions = []oci_file_storage.ClientOptions{}
	if exportOptions, ok := s.D.GetOkExists("export_options"); ok {
		interfaces := exportOptions.([]interface{})
		tmp := make([]oci_file_storage.ClientOptions, len(interfaces))
		for i := range interfaces {
			fieldKeyPrefix := fmt.Sprintf("export_options.%d", i)
			converted, err := mapToClientOptions(fieldKeyPrefix+".%s", s.D)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.ExportOptions = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	return nil
}

func (s *ExportResourceCrud) Delete() error {
	request := oci_file_storage.DeleteExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteExport(context.Background(), request)
	return err
}

func (s *ExportResourceCrud) SetData() error {
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

func mapToClientOptions(fieldKeyFormat string, d *schema.ResourceData) (oci_file_storage.ClientOptions, error) {
	result := oci_file_storage.ClientOptions{}

	if access, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access")); ok {
		tmp := oci_file_storage.ClientOptionsAccessEnum(access.(string))
		result.Access = tmp
	}

	if anonymousGid, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "anonymous_gid")); ok {
		tmp := anonymousGid.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert anonymousGid string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.AnonymousGid = &tmpInt64
	}

	if anonymousUid, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "anonymous_uid")); ok {
		tmp := anonymousUid.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert anonymousUid string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.AnonymousUid = &tmpInt64
	}

	if identitySquash, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identity_squash")); ok {
		tmp := oci_file_storage.ClientOptionsIdentitySquashEnum(identitySquash.(string))
		result.IdentitySquash = tmp
	}

	if requirePrivilegedSourcePort, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "require_privileged_source_port")); ok {
		tmp := requirePrivilegedSourcePort.(bool)
		result.RequirePrivilegedSourcePort = &tmp
	}

	if source, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
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
