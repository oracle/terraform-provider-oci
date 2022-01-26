// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v56/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const DefaultFilePartSize int64 = 128 * 1024 * 1024 // 128MB
const defaultNumberOfGoroutines = 10
const MaxPartSize int64 = 50 * 1024 * 1024 * 1024
const MaxCount int64 = 10000

type MultipartUploadData struct {
	NamespaceName       *string                                 `mandatory:"true"`
	BucketName          *string                                 `mandatory:"true"`
	ObjectName          *string                                 `mandatory:"true"`
	ObjectStorageClient *oci_object_storage.ObjectStorageClient `mandatory:"true"`
	SourcePath          *string                                 `mandatory:"true"`
	SourceInfo          *os.FileInfo                            `mandatory:"true"`
	CacheControl        *string
	ContentDisposition  *string
	ContentMD5          *string
	ContentType         *string
	ContentLanguage     *string
	ContentEncoding     *string
	StorageTier         oci_object_storage.StorageTierEnum
	OpcSseKmsKeyId      *string
	Metadata            map[string]interface{}
	OpcClientRequestID  *string
	RequestMetadata     common.RequestMetadata
}

type objectStorageUploadPartResponse struct {
	response   oci_object_storage.UploadPartResponse
	partNumber *int
	error      error
}

type objectStorageMultiPartUploadContext struct {
	client                  oci_object_storage.ObjectStorageClient
	sourceBlocks            chan objectStorageSourceBlock
	osUploadPartResponses   chan objectStorageUploadPartResponse
	wg                      *sync.WaitGroup
	multipartUploadResponse oci_object_storage.CreateMultipartUploadResponse
	multipartUploadRequest  oci_object_storage.CreateMultipartUploadRequest
}

type objectStorageSourceBlock struct {
	section     *io.SectionReader
	blockNumber *int
}

func resourceObjectStorageMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func resourceObjectStorageMapToOPCMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result["opc-meta-"+k] = v.(string)
	}
	return result
}

// The SDK will return all 'metadata' header keys as lowercase, regardless of how they're specified in the config.
//
// To avoid unnecessary diffs and updates, we need to ensure all config keys for 'metadata' are lowercase.
// This avoids a conflict where our config has a non-lowercase key (e.g. MyKey) while the state file has a lowercase
// key (e.g. mykey) from the SDK.
//
// If we ran a 'terraform plan' on this without any config changes, Terraform will detect a diff between state and
// config; even though nothing changed in the state file.
func validateLowerCaseKeysInMetadata(raw interface{}, fieldName string) ([]string, []error) {
	metadataMap, ok := raw.(map[string]interface{})
	if !ok {
		return nil, []error{fmt.Errorf("Could not convert '%s' to map during validation.", fieldName)}
	}

	errors := []error{}
	for key := range metadataMap {
		lowercaseKey := strings.ToLower(key)
		if key != lowercaseKey {
			errors = append(errors, fmt.Errorf("All '%s' keys must be lowercase. Please specify '%s' as '%s'", fieldName, key, lowercaseKey))
		}
	}

	return nil, errors
}

func validateSourceValue(i interface{}, k string) (s []string, es []error) {
	v, ok := i.(string)
	if !ok {
		es = append(es, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	info, err := os.Stat(v)
	if err != nil {
		es = append(es, fmt.Errorf("cannot get file information for the specified source: %s", v))
		return
	}
	if info.Size() > MaxCount*MaxPartSize {
		es = append(es, fmt.Errorf("the specified source: %s file is too large", v))
	}
	return
}

// Borrowed from https://mijailovic.net/2017/05/09/error-handling-patterns-in-go/
func safeClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && *err == nil {
		*err = cerr
	}
}

func singlePartUpload(multipartUploadData MultipartUploadData) (string, error) {

	sourcePath := *multipartUploadData.SourcePath
	sourceInfo := *multipartUploadData.SourceInfo

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}

	defer safeClose(sourceFile, &err)

	tmpSize := sourceInfo.Size()

	putObjectRequest := &oci_object_storage.PutObjectRequest{
		CacheControl:       multipartUploadData.CacheControl,
		ContentDisposition: multipartUploadData.ContentDisposition,
		ContentEncoding:    multipartUploadData.ContentEncoding,
		ContentLanguage:    multipartUploadData.ContentLanguage,
		ContentType:        multipartUploadData.ContentType,
		BucketName:         multipartUploadData.BucketName,
		ContentLength:      &tmpSize,
		PutObjectBody:      ioutil.NopCloser(sourceFile),
		OpcMeta:            resourceObjectStorageMapToMetadata(multipartUploadData.Metadata),
		NamespaceName:      multipartUploadData.NamespaceName,
		ObjectName:         multipartUploadData.ObjectName,
		StorageTier:        PutObjectStorageTierEnumFromString(string(multipartUploadData.StorageTier)),
	}
	putObjectRequest.RequestMetadata.RetryPolicy = multipartUploadData.RequestMetadata.RetryPolicy

	_, err = multipartUploadData.ObjectStorageClient.PutObject(context.Background(), *putObjectRequest)
	if err != nil {
		return "", err
	}

	id := GetObjectCompositeId(*putObjectRequest.BucketName, *putObjectRequest.NamespaceName, *putObjectRequest.ObjectName)

	return id, nil
}

func StorageTierEnumFromString(s string) oci_object_storage.StorageTierEnum {
	switch s {
	case "Standard":
		return oci_object_storage.StorageTierStandard
	case "Archive":
		return oci_object_storage.StorageTierArchive
	case "InfrequentAccess":
		return oci_object_storage.StorageTierInfrequentAccess
	default:
		return ""
	}
}

func PutObjectStorageTierEnumFromString(s string) oci_object_storage.PutObjectStorageTierEnum {
	switch s {
	case "Standard":
		return oci_object_storage.PutObjectStorageTierStandard
	case "Archive":
		return oci_object_storage.PutObjectStorageTierArchive
	case "InfrequentAccess":
		return oci_object_storage.PutObjectStorageTierInfrequentaccess
	default:
		return ""
	}
}

func MultiPartUpload(multipartUploadData MultipartUploadData) (string, error) {

	sourceInfo := *multipartUploadData.SourceInfo

	if sourceInfo.Size() > DefaultFilePartSize {
		return multiPartUploadImpl(multipartUploadData)
	}

	return singlePartUpload(multipartUploadData)
}

func multiPartUploadImpl(multipartUploadData MultipartUploadData) (string, error) {

	multipartUploadRequest := &oci_object_storage.CreateMultipartUploadRequest{
		NamespaceName:   multipartUploadData.NamespaceName,
		BucketName:      multipartUploadData.BucketName,
		RequestMetadata: multipartUploadData.RequestMetadata,
		CreateMultipartUploadDetails: oci_object_storage.CreateMultipartUploadDetails{
			CacheControl:       multipartUploadData.CacheControl,
			ContentDisposition: multipartUploadData.ContentDisposition,
			ContentEncoding:    multipartUploadData.ContentEncoding,
			ContentLanguage:    multipartUploadData.ContentLanguage,
			ContentType:        multipartUploadData.ContentType,
			Object:             multipartUploadData.ObjectName,
			StorageTier:        multipartUploadData.StorageTier,
			Metadata:           resourceObjectStorageMapToOPCMetadata(multipartUploadData.Metadata),
		},
	}
	source := multipartUploadData.SourcePath
	client := multipartUploadData.ObjectStorageClient

	file, err := os.Open(*source)
	if err != nil {
		return "", fmt.Errorf("error opening source file for upload \"%v\": %s", source, err)
	}
	defer safeClose(file, &err)

	sourceBlocks, err := objectMultiPartSplit(file)
	if err != nil {
		return "", fmt.Errorf("error splitting source file for upload \"%v\": %s", source, err)
	}

	multipartUploadResponse, err := client.CreateMultipartUpload(context.Background(), *multipartUploadRequest)
	if err != nil {
		return "", fmt.Errorf("error creating object in the Oracle cloud \"%v\": %s", source, err)
	}

	workerCount := defaultNumberOfGoroutines

	osUploadPartResponses := make(chan objectStorageUploadPartResponse, len(sourceBlocks))
	sourceBlocksChan := make(chan objectStorageSourceBlock, len(sourceBlocks))

	wg := &sync.WaitGroup{}
	wg.Add(len(sourceBlocks))

	for _, sourceBlock := range sourceBlocks {
		sourceBlocksChan <- sourceBlock
	}
	close(sourceBlocksChan)

	for i := 0; i < workerCount; i++ {
		go uploadPartsWorker(objectStorageMultiPartUploadContext{
			client:                  *client,
			wg:                      wg,
			multipartUploadResponse: multipartUploadResponse,
			multipartUploadRequest:  *multipartUploadRequest,
			sourceBlocks:            sourceBlocksChan,
			osUploadPartResponses:   osUploadPartResponses,
		})
	}

	wg.Wait()

	close(osUploadPartResponses)

	commitMultipartUploadPartDetails := make([]oci_object_storage.CommitMultipartUploadPartDetails, len(sourceBlocks))

	osResponseIndex := 0
	var uploadPartRespErr error
	for osUploadPartResponse := range osUploadPartResponses {
		if osUploadPartResponse.error != nil {
			uploadPartRespErr = osUploadPartResponse.error
			break
		}

		commitMultipartUploadPartDetails[osResponseIndex] = oci_object_storage.CommitMultipartUploadPartDetails{
			PartNum: osUploadPartResponse.partNumber,
			Etag:    osUploadPartResponse.response.ETag,
		}
		osResponseIndex++
	}

	if uploadPartRespErr != nil {
		// just aborting the multi upload for now; but the service itself will handle the same request again
		abortMultipartUploadRequest := oci_object_storage.AbortMultipartUploadRequest{
			NamespaceName:      multipartUploadResponse.Namespace,
			BucketName:         multipartUploadResponse.Bucket,
			ObjectName:         multipartUploadResponse.Object,
			UploadId:           multipartUploadResponse.MultipartUpload.UploadId,
			OpcClientRequestId: multipartUploadResponse.OpcClientRequestId,
			RequestMetadata:    multipartUploadRequest.RequestMetadata,
		}

		_, err := client.AbortMultipartUpload(context.Background(), abortMultipartUploadRequest)

		if err != nil {
			log.Println("[WARN] Aborting the multi part upload failed")
		}

		return "", fmt.Errorf("failed to upload object parts of \"%v\" to the Oracle cloud: %s", source, uploadPartRespErr)
	}

	commitMultipartUploadRequest := oci_object_storage.CommitMultipartUploadRequest{
		UploadId:           multipartUploadResponse.MultipartUpload.UploadId,
		NamespaceName:      multipartUploadResponse.Namespace,
		BucketName:         multipartUploadResponse.Bucket,
		ObjectName:         multipartUploadResponse.Object,
		OpcClientRequestId: multipartUploadResponse.OpcClientRequestId,
		RequestMetadata:    multipartUploadRequest.RequestMetadata,
	}
	commitMultipartUploadRequest.PartsToCommit = commitMultipartUploadPartDetails

	_, err = client.CommitMultipartUpload(context.Background(), commitMultipartUploadRequest)
	if err != nil {
		return "", fmt.Errorf("failed to commit multi part upload of \"%v\" to the service: %s", source, err)
	}

	id := GetObjectCompositeId(*commitMultipartUploadRequest.BucketName, *commitMultipartUploadRequest.NamespaceName, *commitMultipartUploadRequest.ObjectName)

	return id, nil
}

func objectMultiPartSplit(file *os.File) ([]objectStorageSourceBlock, error) {

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get FileInfo for the source %q: %s", file.Name(), err)
	}

	offsets, limits, err := SplitSizeToOffsetsAndLimits(info.Size())
	sourceBlocks := make([]objectStorageSourceBlock, len(offsets))
	for index := 0; index < len(offsets); index++ {
		tmpIndex := index + 1
		sourceBlocks[index] = objectStorageSourceBlock{
			section:     io.NewSectionReader(file, offsets[index], limits[index]),
			blockNumber: &tmpIndex,
		}
	}

	return sourceBlocks, nil
}

func SplitSizeToOffsetsAndLimits(infoSize int64) ([]int64, []int64, error) {
	partSize := DefaultFilePartSize
	remainingPart := int64(0)

	totalNumber := infoSize / partSize
	if infoSize%partSize > 0 {
		remainingPart = 1
	}

	if totalNumber+remainingPart > MaxCount {
		remainingPart = 0
		if infoSize%MaxCount > 0 {
			remainingPart = 1
		}

		partSize = infoSize/MaxCount + remainingPart

		if partSize > MaxPartSize {
			return nil, nil, fmt.Errorf("the %v size of the source object is more than the service limit", infoSize)
		}
	}

	totalNumber = infoSize / partSize
	if infoSize%partSize > 0 {
		totalNumber++
	}

	offsets := make([]int64, totalNumber)
	limits := make([]int64, totalNumber)
	index := 0
	for offset := int64(0); offset < infoSize; offset += partSize {
		tmpLimit := infoSize - offset
		if partSize < tmpLimit {
			tmpLimit = partSize
		}
		offsets[index] = offset
		limits[index] = tmpLimit
		index++
	}

	return offsets, limits, nil
}

func uploadPartsWorker(ctx objectStorageMultiPartUploadContext) {
	for sourceBlock := range ctx.sourceBlocks {

		block := make([]byte, sourceBlock.section.Size())
		_, err := sourceBlock.section.Read(block)
		if err != nil && err != io.EOF {
			if sourceBlock.blockNumber != nil {
				log.Printf("[ERROR] failed to read source file section %v: %s\n", *sourceBlock.blockNumber, err)
			}

			osUploadPartResponse := &objectStorageUploadPartResponse{
				error: err,
			}
			ctx.osUploadPartResponses <- *osUploadPartResponse
			ctx.wg.Done()
			continue
		}
		tmpLength := int64(len(block))
		uploadPartRequest := &oci_object_storage.UploadPartRequest{
			UploadId:        ctx.multipartUploadResponse.UploadId,
			ObjectName:      ctx.multipartUploadResponse.Object,
			NamespaceName:   ctx.multipartUploadResponse.Namespace,
			BucketName:      ctx.multipartUploadResponse.Bucket,
			RequestMetadata: ctx.multipartUploadRequest.RequestMetadata,
			ContentLength:   &tmpLength,
			UploadPartBody:  ioutil.NopCloser(bytes.NewBuffer(block)),
			UploadPartNum:   sourceBlock.blockNumber,
		}

		uploadPartResponse, err := ctx.client.UploadPart(context.Background(), *uploadPartRequest)

		osUploadPartResponse := &objectStorageUploadPartResponse{
			response:   uploadPartResponse,
			error:      err,
			partNumber: uploadPartRequest.UploadPartNum,
		}

		ctx.osUploadPartResponses <- *osUploadPartResponse
		ctx.wg.Done()
	}
}

func (s *ObjectStorageObjectResourceCrud) createSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceObjectStorageClient, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceObjectStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceObjectStorageClient
	}
	s.SourceRegionClient.SetRegion(region)

	return nil
}

func copyObjectWaitForWorkRequest(wId *string, entityType string, timeout time.Duration, disableFoundRetries bool, client *oci_object_storage.ObjectStorageClient) error {

	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "object_storage")
	retryPolicy.ShouldRetryOperation = objectStorageWorkRequestShouldRetryFunc(timeout)

	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_object_storage.WorkRequestStatusAccepted),
			string(oci_object_storage.WorkRequestStatusInProgress),
			string(oci_object_storage.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_object_storage.WorkRequestSummaryStatusCompleted),
			string(oci_object_storage.WorkRequestSummaryStatusCanceled),
			string(oci_object_storage.WorkRequestStatusFailed),
		},
		Refresh: func() (interface{}, string, error) {
			getWorkRequestRequest := oci_object_storage.GetWorkRequestRequest{}
			getWorkRequestRequest.WorkRequestId = wId
			getWorkRequestRequest.RequestMetadata.RetryPolicy = retryPolicy
			workRequestResponse, err := client.GetWorkRequest(context.Background(), getWorkRequestRequest)
			wr := &workRequestResponse.WorkRequest
			return workRequestResponse, string(wr.Status), err
		},
		Timeout: timeout,
	}

	// Set PollInterval to 1 for replay mode.
	if httpreplay.ShouldRetryImmediately() {
		stateConf.PollInterval = 1
	}

	wrr, e := stateConf.WaitForState()
	if e != nil {
		return fmt.Errorf("work request did not succeed, workId: %s, entity: %s. Message: %s", *wId, entityType, e)
	}

	wr := wrr.(oci_object_storage.GetWorkRequestResponse).WorkRequest
	if wr.Status == oci_object_storage.WorkRequestStatusFailed {
		errorMessage, _ := getObjectStorageErrorFromWorkRequest(wId, client, disableFoundRetries)
		return fmt.Errorf("work request did not succeed, workId: %s, entity: %s. Message: %s", *wId, entityType, errorMessage)
	}

	return nil

}

func objectStorageWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		//Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		//Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "object_storage", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if objectRes, ok := response.Response.(oci_object_storage.GetWorkRequestResponse); ok {
			return objectRes.TimeFinished == nil
		}
		return false
	}
}

func getObjectStorageErrorFromWorkRequest(workRequestId *string, client *oci_object_storage.ObjectStorageClient, disableFoundAutoRetries bool) (string, error) {
	req := oci_object_storage.ListWorkRequestErrorsRequest{}
	req.WorkRequestId = workRequestId
	req.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(disableFoundAutoRetries, "object_storage")
	res, err := client.ListWorkRequestErrors(context.Background(), req)

	if err != nil {
		return "", err
	}

	allErrs := make([]string, 0)
	for _, errs := range res.Items {
		allErrs = append(allErrs, *errs.Message)
	}

	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage, nil
}

func DeleteAllObjectVersions(client *oci_object_storage.ObjectStorageClient, bucket string, namespace string, prefix string) error {
	request := oci_object_storage.ListObjectVersionsRequest{}

	request.BucketName = &bucket
	request.NamespaceName = &namespace

	if prefix != "" {
		request.Prefix = &prefix
	}

	response, err := client.ListObjectVersions(context.Background(), request)
	if err != nil {
		return err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

		listResponse, err := client.ListObjectVersions(context.Background(), request)
		if err != nil {
			return err
		}
		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	var errors []string
	for _, objectVersion := range response.Items {

		deleteObjectVersionRequest := oci_object_storage.DeleteObjectRequest{}
		deleteObjectVersionRequest.BucketName = &bucket
		deleteObjectVersionRequest.NamespaceName = &namespace
		deleteObjectVersionRequest.ObjectName = objectVersion.Name
		deleteObjectVersionRequest.VersionId = objectVersion.VersionId

		deleteObjectVersionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

		_, err := client.DeleteObject(context.Background(), deleteObjectVersionRequest)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", errors)
	}

	return nil
}
