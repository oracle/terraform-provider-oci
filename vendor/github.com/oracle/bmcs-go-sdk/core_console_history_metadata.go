// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"strconv"
	"time"
)

// ConsoleHistoryMetadata describes console history metadata
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/
type ConsoleHistoryMetadata struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AvailabilityDomain string    `json:"availabilityDomain"`
	CompartmentID      string    `json:"compartmentId"`
	DisplayName        string    `json:"displayName"`
	ID                 string    `json:"id"`
	InstanceID         string    `json:"instanceId"`
	State              string    `json:"lifecycleState"`
	TimeCreated        time.Time `json:"TimeCreated"`
}

// ConsoleHistoryData contains all or part of an instance console history
// snapshot.  If BytesRemaining is greater than zero, Data is only part of the
// total history.  The remainder may be fetched on subsequent calls to
// ShowConsoleHistoryData, populating Offset and Limit options.
type ConsoleHistoryData struct {
	BytesRemaining int
	Data           string
}

// ListConsoleHistories contains a list of Console History Metadata
type ListConsoleHistories struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	ConsoleHistories []ConsoleHistoryMetadata
}

func (l *ListConsoleHistories) GetList() interface{} {
	return &l.ConsoleHistories
}

// ListConsoleHistories shows the metadata for the specified compartment or instance
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/ListConsoleHistories
func (c *Client) ListConsoleHistories(compartmentID string, opts *ListConsoleHistoriesOptions) (icHistories *ListConsoleHistories, e error) {
	required := listOCIDRequirement{CompartmentID: compartmentID}

	details := &requestDetails{
		name:     resourceInstanceConsoleHistories,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	icHistories = &ListConsoleHistories{}
	e = resp.unmarshal(icHistories)
	return
}

// CaptureConsoleHistory captures the most recent serial console data (up to a megabyte) for the specified instance.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/CaptureConsoleHistory
func (c *Client) CaptureConsoleHistory(instanceID string, opts *RetryTokenOptions) (icHistory *ConsoleHistoryMetadata, e error) {
	required := struct {
		InstanceID string `header:"-" json:"instanceId" url:"-"`
	}{
		InstanceID: instanceID,
	}

	details := &requestDetails{
		name:     resourceInstanceConsoleHistories,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	icHistory = &ConsoleHistoryMetadata{}
	e = resp.unmarshal(icHistory)
	return
}

// GetConsoleHistory shows the metadata for the specified console history
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/GetConsoleHistory
func (c *Client) GetConsoleHistory(instanceID string) (consoleHistoryMetadata *ConsoleHistoryMetadata, e error) {
	details := &requestDetails{
		name: resourceInstanceConsoleHistories,
		ids:  urlParts{instanceID},
	}
	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	consoleHistoryMetadata = &ConsoleHistoryMetadata{}
	e = resp.unmarshal(consoleHistoryMetadata)
	return
}

// ShowConsoleHistoryData gets the actual console history data (not the metadata).
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/GetConsoleHistoryContent
func (c *Client) ShowConsoleHistoryData(instanceConsoleHistoryID string, opts *ConsoleHistoryDataOptions) (hist *ConsoleHistoryData, e error) {
	details := &requestDetails{
		name:     resourceInstanceConsoleHistories,
		ids:      urlParts{instanceConsoleHistoryID, dataURLPart},
		optional: opts,
	}
	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	hist = &ConsoleHistoryData{Data: string(resp.body[:])}
	s := resp.header.Get(headerBytesRemaining)

	if s != "" {
		if hist.BytesRemaining, e = strconv.Atoi(s); e != nil {
			return
		}
	}

	return
}
