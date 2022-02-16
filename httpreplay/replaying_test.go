// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Run with a command something like:
//   go test -run TestTransform -tags replay

//go:build replay
// +build replay

package httpreplay

import (
	"encoding/json"
	"testing"
)

type transformTestCase struct {
	Description string
	Interaction
	Request
	Response
	ExpectedResponse Response
}

type testObj map[string]string

func TestTransform(t *testing.T) {

	testCases := []transformTestCase{
		transformTestCase{
			Description: `Empty body -> stay empty.`,
			Interaction: Interaction{
				Request:  Request{Body: ""},
				Response: Response{Body: ""},
			},
			Request:          Request{Body: ""},
			Response:         Response{Body: ""},
			ExpectedResponse: Response{Body: ""},
		},
		transformTestCase{
			Description: `Some unchanging fields -> stay unchanging`,
			Interaction: Interaction{
				Request:  Request{Body: `{"displayName": "unchanged"}`},
				Response: Response{Body: `{"displayName": "unchanged"}`},
			},
			Request:          Request{Body: `{"displayName": "unchanged"}`},
			Response:         Response{Body: `{"displayName": "unchanged"}`},
			ExpectedResponse: Response{Body: `{"displayName": "unchanged"}`},
		},
		transformTestCase{
			Description: `Change to the displayName -> change in the result`,
			Interaction: Interaction{
				Request:  Request{Body: `{"displayName": "before"}`},
				Response: Response{Body: `{"displayName": "before"}`},
			},
			Request:          Request{Body: `{"displayName": "after"}`},
			Response:         Response{Body: `{"displayName": "before"}`},
			ExpectedResponse: Response{Body: `{"displayName": "after"}`},
		},
		transformTestCase{
			Description: `Displayname adds " (block storage)" -> still change the result`,
			Interaction: Interaction{
				Request:  Request{Body: `{"displayName": "before"}`},
				Response: Response{Body: `{"displayName": "before (block storage)"}`},
			},
			Request:          Request{Body: `{"displayName": "afterr"}`},
			Response:         Response{Body: `{"displayName": "before (block storage)"}`},
			ExpectedResponse: Response{Body: `{"displayName": "afterr (block storage)"}`},
		},
		transformTestCase{
			Description: `Displayname adds " (block storage)" -> no change if first part different length`,
			Interaction: Interaction{
				Request:  Request{Body: `{"displayName": "longerkey"}`},
				Response: Response{Body: `{"displayName": "longerkey (block storage)"}`},
			},
			Request:          Request{Body: `{"displayName": "short"}`},
			Response:         Response{Body: `{"displayName": "longerkey (block storage)"}`},
			ExpectedResponse: Response{Body: `{"displayName": "longerkey (block storage)"}`},
		},
	}

	for _, testCase := range testCases {
		var err error

		unm := func(in string, field string) interface{} {
			if err != nil {
				return nil
			}
			var result interface{}
			if result, err = unmarshal([]byte(in)); err != nil {
				t.Errorf("unable to unmarshal test %v %v %v", testCase.Description, field, err)
				return nil
			}
			return result
		}
		testCase.Interaction.Request.BodyParsed = unm(testCase.Interaction.Request.Body, "Interaction.Request")
		testCase.Interaction.Response.BodyParsed = unm(testCase.Interaction.Response.Body, "Interaction.Response")
		testCase.Request.BodyParsed = unm(testCase.Request.Body, "Request")
		testCase.Response.BodyParsed = unm(testCase.Response.Body, "Response")
		testCase.ExpectedResponse.BodyParsed = unm(testCase.ExpectedResponse.Body, "ExpectedResponse")
		if err != nil {
			return
		}
		transformer(&testCase.Request, testCase.Interaction, &testCase.Response)

		// Re-marshal both responses to get them easily comparable
		m := func(in interface{}, field string) string {
			if err != nil {
				return ""
			}
			var result []byte
			if result, err = json.Marshal(in); err != nil {
				t.Error("unable to marshal test %v %v %v", testCase.Description, field, err)
			}
			return string(result)
		}
		response := m(testCase.Response.BodyParsed, "Response")
		expectedResponse := m(testCase.ExpectedResponse.BodyParsed, "ExpectedResponse")
		if response != expectedResponse {
			t.Errorf("test <<%v>>: actual response <<%v>> does not match expected <<%v>>", testCase.Description, response, expectedResponse)
		}
	}
}
