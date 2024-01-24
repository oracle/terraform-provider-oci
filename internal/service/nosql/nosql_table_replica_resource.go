// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NosqlTableReplicaResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNosqlTableReplica,
		Read:     readNosqlTableReplica,
		Delete:   deleteNosqlTableReplica,
		Schema: map[string]*schema.Schema{
			// Required
			"region": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: anyChangeSuppressFunction,
			},
			"table_name_or_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: anyChangeSuppressFunction,
			},

			// Optional
			"compartment_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         false,
				DiffSuppressFunc: anyChangeSuppressFunction,
			},
			"max_read_units": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ForceNew:         false,
				DiffSuppressFunc: anyChangeSuppressFunction,
			},
			"max_write_units": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ForceNew:         false,
				DiffSuppressFunc: anyChangeSuppressFunction,
			},

			// Computed
		},
	}
}

func createNosqlTableReplica(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.CreateResource(d, sync)
}

func readNosqlTableReplica(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.ReadResource(sync)
}

func deleteNosqlTableReplica(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NosqlTableReplicaResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_nosql.NosqlClient
	Res                    *oci_nosql.Replica
	DisableNotFoundRetries bool
}

func (s *NosqlTableReplicaResourceCrud) ID() string {
	return *s.Res.Region
}

func (s *NosqlTableReplicaResourceCrud) Create() error {
	request := oci_nosql.CreateReplicaRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if maxReadUnits, ok := s.D.GetOkExists("max_read_units"); ok {
		tmp := maxReadUnits.(int)
		request.MaxReadUnits = &tmp
	}

	if maxWriteUnits, ok := s.D.GetOkExists("max_write_units"); ok {
		tmp := maxWriteUnits.(int)
		request.MaxWriteUnits = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.CreateReplica(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTableReplicaFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql"), oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NosqlTableReplicaResourceCrud) getTableReplicaFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_nosql.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	tableReplicaId, err := tableReplicaWaitForWorkRequest(workId, "TABLE",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, tableReplicaId)
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
	s.D.SetId(*tableReplicaId)

	return s.Get()
}

func tableReplicaWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func tableReplicaWaitForWorkRequest(wId *string, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_nosql.NosqlClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "nosql")
	retryPolicy.ShouldRetryOperation = tableReplicaWorkRequestShouldRetryFunc(timeout)

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
	if identifier == nil || response.Status == oci_nosql.WorkRequestStatusFailed || response.Status == oci_nosql.WorkRequestStatusCanceled {
		return nil, getErrorFromNosqlTableReplicaWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromNosqlTableReplicaWorkRequest(client *oci_nosql.NosqlClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum) error {
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

func (s *NosqlTableReplicaResourceCrud) Get() error {
	request := oci_nosql.GetTableRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	} else if s.D.Id() != "" {
		tmp := s.D.Id()
		request.TableNameOrId = &tmp
	}

	var regionName string
	if region, ok := s.D.GetOkExists("region"); ok {
		regionName = region.(string)
	}

	region, table, err := parseTableReplicaCompositeId(s.D.Id())
	if err == nil {
		regionName = region
		request.TableNameOrId = &table
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.GetTable(context.Background(), request)
	if err != nil {
		return err
	}

	for _, replica := range response.Table.Replicas {
		if strings.EqualFold(*replica.Region, regionName) {
			s.Res = &replica
			break
		}
	}

	return nil
}

func (s *NosqlTableReplicaResourceCrud) Delete() error {
	request := oci_nosql.DeleteReplicaRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.DeleteReplica(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := tableReplicaWaitForWorkRequest(workId, "TABLE",
		oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NosqlTableReplicaResourceCrud) SetData() error {
	region, tableNameOrId, err := parseTableReplicaCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("region", &region)
		s.D.Set("table_name_or_id", &tableNameOrId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}
	s.D.Set("max_write_units", *s.Res.MaxWriteUnits)

	return nil
}

func GetTableReplicaCompositeId(region string, tableNameOrId string) string {
	region = url.PathEscape(region)
	tableNameOrId = url.PathEscape(tableNameOrId)
	compositeId := "tables/" + tableNameOrId + "/replicas/" + region
	return compositeId
}

func parseTableReplicaCompositeId(compositeId string) (region string, tableNameOrId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("tables/.*/replicas/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	tableNameOrId, _ = url.PathUnescape(parts[1])
	region, _ = url.PathUnescape(parts[3])
	return
}

func anyChangeSuppressFunction(k string, old string, new string, d *schema.ResourceData) bool {
	if old == "" && new != "" {
		return false
	}
	return true
}
