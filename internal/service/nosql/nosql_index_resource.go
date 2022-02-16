// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_nosql "github.com/oracle/oci-go-sdk/v58/nosql"
)

func NosqlIndexResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNosqlIndex,
		Read:     readNosqlIndex,
		Delete:   deleteNosqlIndex,
		Schema: map[string]*schema.Schema{
			// Required
			"keys": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"column_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"json_field_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"json_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"table_name_or_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_if_not_exists": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"table_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNosqlIndex(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlIndexResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.CreateResource(d, sync)
}

func readNosqlIndex(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlIndexResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.ReadResource(sync)
}

func deleteNosqlIndex(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlIndexResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NosqlIndexResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_nosql.NosqlClient
	Res                    *oci_nosql.Index
	DisableNotFoundRetries bool
}

func (s *NosqlIndexResourceCrud) ID() string {
	return *s.Res.Name
}

func (s *NosqlIndexResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_nosql.IndexLifecycleStateCreating),
	}
}

func (s *NosqlIndexResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_nosql.IndexLifecycleStateActive),
	}
}

func (s *NosqlIndexResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_nosql.IndexLifecycleStateDeleting),
	}
}

func (s *NosqlIndexResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_nosql.IndexLifecycleStateDeleted),
	}
}

func (s *NosqlIndexResourceCrud) Create() error {
	request := oci_nosql.CreateIndexRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isIfNotExists, ok := s.D.GetOkExists("is_if_not_exists"); ok {
		tmp := isIfNotExists.(bool)
		request.IsIfNotExists = &tmp
	}

	if keys, ok := s.D.GetOkExists("keys"); ok {
		interfaces := keys.([]interface{})
		tmp := make([]oci_nosql.IndexKey, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "keys", stateDataIndex)
			converted, err := s.mapToIndexKey(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("keys") {
			request.Keys = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.CreateIndex(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIndexFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql"), oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NosqlIndexResourceCrud) getIndexFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_nosql.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	indexId, err := indexWaitForWorkRequest(workId, "TABLE",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, indexId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_nosql.DeleteWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*indexId)

	return s.Get()
}

func indexWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "nosql", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_nosql.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func indexWaitForWorkRequest(wId *string, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_nosql.NosqlClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "nosql")
	retryPolicy.ShouldRetryOperation = indexWorkRequestShouldRetryFunc(timeout)

	response := oci_nosql.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_nosql.WorkRequestStatusInProgress),
			string(oci_nosql.WorkRequestStatusAccepted),
			string(oci_nosql.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_nosql.WorkRequestStatusSucceeded),
			string(oci_nosql.WorkRequestStatusFailed),
			string(oci_nosql.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_nosql.GetWorkRequestRequest{
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
		if strings.Contains(*res.EntityType, entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.WorkRequest.Status == oci_nosql.WorkRequestStatusFailed || response.WorkRequest.Status == oci_nosql.WorkRequestStatusCanceled {
		return nil, getErrorFromNosqlIndexWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromNosqlIndexWorkRequest(client *oci_nosql.NosqlClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_nosql.ListWorkRequestErrorsRequest{
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *NosqlIndexResourceCrud) Get() error {
	request := oci_nosql.GetIndexRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if indexName, ok := s.D.GetOkExists("name"); ok {
		tmp := indexName.(string)
		request.IndexName = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	index, table, err := parseIndexCompositeId(s.D.Id())
	if err == nil {
		request.IndexName = &index
		request.TableNameOrId = &table
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.GetIndex(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Index
	return nil
}

func (s *NosqlIndexResourceCrud) Delete() error {
	request := oci_nosql.DeleteIndexRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if indexName, ok := s.D.GetOkExists("name"); ok {
		tmp := indexName.(string)
		request.IndexName = &tmp
	}

	if isIfExists, ok := s.D.GetOkExists("is_if_exists"); ok {
		tmp := isIfExists.(bool)
		request.IsIfExists = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.DeleteIndex(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := indexWaitForWorkRequest(workId, "TABLE",
		oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NosqlIndexResourceCrud) SetData() error {

	indexName, tableNameOrId, err := parseIndexCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &indexName)
		s.D.Set("table_name_or_id", &tableNameOrId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	keys := []interface{}{}
	for _, item := range s.Res.Keys {
		keys = append(keys, IndexKeyToMap(item))
	}
	s.D.Set("keys", keys)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TableId != nil {
		s.D.Set("table_id", *s.Res.TableId)
	}

	if s.Res.TableName != nil {
		s.D.Set("table_name", *s.Res.TableName)
	}

	return nil
}

func GetIndexCompositeId(indexName string, tableNameOrId string) string {
	indexName = url.PathEscape(indexName)
	tableNameOrId = url.PathEscape(tableNameOrId)
	compositeId := "tables/" + tableNameOrId + "/indexes/" + indexName
	return compositeId
}

func parseIndexCompositeId(compositeId string) (indexName string, tableNameOrId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("tables/.*/indexes/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	tableNameOrId, _ = url.PathUnescape(parts[1])
	indexName, _ = url.PathUnescape(parts[3])

	return
}

func (s *NosqlIndexResourceCrud) mapToIndexKey(fieldKeyFormat string) (oci_nosql.IndexKey, error) {
	result := oci_nosql.IndexKey{}

	if columnName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "column_name")); ok {
		tmp := columnName.(string)
		result.ColumnName = &tmp
	}

	if jsonFieldType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "json_field_type")); ok {
		tmp := jsonFieldType.(string)
		result.JsonFieldType = &tmp
	}

	if jsonPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "json_path")); ok {
		tmp := jsonPath.(string)
		result.JsonPath = &tmp
	}

	return result, nil
}

func IndexKeyToMap(obj oci_nosql.IndexKey) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	if obj.JsonFieldType != nil {
		result["json_field_type"] = string(*obj.JsonFieldType)
	}

	if obj.JsonPath != nil {
		result["json_path"] = string(*obj.JsonPath)
	}

	return result
}

func IndexSummaryToMap(obj oci_nosql.IndexSummary) map[string]interface{} {
	result := map[string]interface{}{}

	keys := []interface{}{}
	for _, item := range obj.Keys {
		keys = append(keys, IndexKeyToMap(item))
	}
	result["keys"] = keys

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}
