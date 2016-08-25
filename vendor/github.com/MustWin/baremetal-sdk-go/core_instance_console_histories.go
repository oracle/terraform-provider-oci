package baremetal

import (
	"net/http"
	"time"
)

// ConsoleHistoryMetadata describes console history metadata
//
// See: https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#ConsoleHistoryMetadata
type ConsoleHistoryMetadata struct {
	ETaggedResource
	AvailabilityDomain string    `json:"availabilityDomain"`
	CompartmentID      string    `json:"compartmentId"`
	DisplayName        string    `json:"displayName"`
	ID                 string    `json:"id"`
	InstanceID         string    `json:"instanceId"`
	State              string    `json:"state"`
	TimeCreated        time.Time `json:"TimeCreated"`
}

// InstanceConsoleHistoriesMetadataList contains a list of Console History Metadata
type ListInstanceConsoleHistoriesMetadatas struct {
	ResourceContainer
	InstanceConsoleHistoriesMetadatas []ConsoleHistoryMetadata
}

func (l *ListInstanceConsoleHistoriesMetadatas) GetList() interface{} {
	return &l.InstanceConsoleHistoriesMetadatas
}

// ListConsoleHistories shows the metadata for the specified compartment or instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listConsoleHistories
func (c *Client) ListConsoleHistories(compartmentID string, opts ...Options) (icHistories *ListInstanceConsoleHistoriesMetadatas, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceInstanceConsoleHistories,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	icHistories = &ListInstanceConsoleHistoriesMetadatas{}
	e = resp.unmarshal(icHistories)
	return
}

// CaptureConsoleHistory captures the most recent serial console data (up to a megabyte) for the specified instance.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#captureConsoleHistory
func (c *Client) CaptureConsoleHistory(instanceID string, opts ...Options) (icHistory *ConsoleHistoryMetadata, e error) {
	createRequest := struct {
		InstanceID string `json:"instanceId"`
	}{
		InstanceID: instanceID,
	}

	reqOpts := &sdkRequestOptions{
		body:    createRequest,
		name:    resourceInstanceConsoleHistories,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	icHistory = &ConsoleHistoryMetadata{}
	e = resp.unmarshal(icHistory)
	return
}

// GetConsoleHistory shows the metadata for the specified console history
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getConsoleHistory
func (c *Client) GetConsoleHistory(instanceID string, opts ...Options) (consoleHistoryMetadata *ConsoleHistoryMetadata, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceInstanceConsoleHistories,
		options: opts,
		ids:     urlParts{instanceID},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	consoleHistoryMetadata = &ConsoleHistoryMetadata{}
	e = resp.unmarshal(consoleHistoryMetadata)
	return
}

// DeleteConsoleHistory deletes the specified console history metadata and the console history data
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteConsoleHistory
func (c *Client) DeleteConsoleHistory(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceInstanceConsoleHistories,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ShowConsoleHistoryData gets the actual console history data (not the metadata).
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#showConsoleHistoryData
func (c *Client) ShowConsoleHistoryData(instanceConsoleHistoryID string, opts ...Options) (consoleHistoryData string, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceInstanceConsoleHistories,
		options: opts,
		ids:     urlParts{instanceConsoleHistoryID, dataURLPart},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	consoleHistoryData = string(resp.body[:])

	return
}
