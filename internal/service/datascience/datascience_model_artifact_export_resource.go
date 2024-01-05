// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelArtifactExportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceModelArtifactExport,
		Read:     readDatascienceModelArtifactExport,
		Delete:   deleteDatascienceModelArtifactExport,
		Schema: map[string]*schema.Schema{
			"model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"artifact_source_type": { //Enum
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ORACLE_OBJECT_STORAGE",
				}, true),
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_object_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDatascienceModelArtifactExport(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelArtifactExportResourceCreate{}
	sync.D = d
	sync.DisableNotFoundRetries = true
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceModelArtifactExport(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatascienceModelArtifactExport(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatascienceModelArtifactExportResourceCreate struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.ExportModelArtifactResponse
	DisableNotFoundRetries bool
}

func (s *DatascienceModelArtifactExportResourceCreate) ID() string {
	return *s.Res.OpcRequestId
}

func (s *DatascienceModelArtifactExportResourceCreate) SetData() error {
	return nil
}

func (s *DatascienceModelArtifactExportResourceCreate) Create() error {
	request := oci_datascience.ExportModelArtifactRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	tmp, err := s.mapToArtifactExportDetails()
	if err != nil {
		return fmt.Errorf("unable to convert artifact_export_details, encountered error: %v", err)
	}
	request.ExportModelArtifactDetails = oci_datascience.ExportModelArtifactDetails{}
	request.ExportModelArtifactDetails.ArtifactExportDetails = &tmp
	request.OpcRetryToken = oci_common.String(oci_common.RetryToken())
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	var emptyResponse oci_datascience.ExportModelArtifactResponse
	response, err := s.Client.ExportModelArtifact(context.Background(), request)
	if response != emptyResponse && response.RawResponse.StatusCode == 409 {
		s.Res = &response
		return nil
	}
	if err != nil {
		return err
	}

	s.Res = &response

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, ExportModelArtifactWorkRequestErr := modelArtifactExportWaitForWorkRequest(workId, "model-artifact",
		oci_datascience.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
	return ExportModelArtifactWorkRequestErr
}

func (s *DatascienceModelArtifactExportResourceCreate) mapToArtifactExportDetails() (oci_datascience.ArtifactExportDetails, error) {
	var baseObject oci_datascience.ArtifactExportDetails

	//discriminator
	artifactSourceTypeRaw, ok := s.D.GetOkExists("artifact_source_type")
	var artifactSourceType string
	if ok {
		artifactSourceType = artifactSourceTypeRaw.(string)
	} else {
		artifactSourceType = "" // default value
	}
	switch strings.ToLower(artifactSourceType) {
	case strings.ToLower("ORACLE_OBJECT_STORAGE"):
		details := oci_datascience.ArtifactExportDetailsObjectStorage{}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if sourceBucket, ok := s.D.GetOkExists("source_bucket"); ok {
			tmp := sourceBucket.(string)
			details.SourceBucket = &tmp
		}
		if sourceObjectName, ok := s.D.GetOkExists("source_object_name"); ok {
			tmp := sourceObjectName.(string)
			details.SourceObjectName = &tmp
		}
		if sourceRegion, ok := s.D.GetOkExists("source_region"); ok {
			tmp := sourceRegion.(string)
			details.SourceRegion = &tmp
		}
		baseObject = &details
	default:
		log.Printf("[WARN] Received 'artifact_source_type' of unknown type %v", artifactSourceType)
		return baseObject, nil
	}
	return baseObject, nil
}

func modelArtifactExportWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_datascience.WorkRequestStatusInProgress),
			string(oci_datascience.WorkRequestStatusAccepted),
			string(oci_datascience.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datascience.WorkRequestStatusSucceeded),
			string(oci_datascience.WorkRequestStatusFailed),
			string(oci_datascience.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datascience.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_datascience.WorkRequestStatusFailed || response.Status == oci_datascience.WorkRequestStatusCanceled {
		return nil, getErrorFromDatascienceModelArtifactExportWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func modelArtifactExportWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datascience", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datascience.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func getErrorFromDatascienceModelArtifactExportWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datascience.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("Work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}
