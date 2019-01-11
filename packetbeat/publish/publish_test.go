// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build !integration

package publish

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
)

func testEvent() beat.Event {
	return beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type": "test",
			"src":  &common.Endpoint{},
			"dst":  &common.Endpoint{},
		},
	}
}

// Test that FilterEvent detects events that do not contain the required fields
// and returns error.
func TestFilterEvent(t *testing.T) {
	var testCases = []struct {
		f   func() beat.Event
		err string
	}{
		{testEvent, ""},
		{
			func() beat.Event {
				e := testEvent()
				e.Fields["@timestamp"] = time.Now()
				return e
			},
			"duplicate '@timestamp'",
		},
		{
			func() beat.Event {
				e := testEvent()
				e.Timestamp = time.Time{}
				return e
			},
			"missing '@timestamp'",
		},
		{
			func() beat.Event {
				e := testEvent()
				delete(e.Fields, "type")
				return e
			},
			"missing 'type'",
		},
		{
			func() beat.Event {
				e := testEvent()
				e.Fields["type"] = 123
				return e
			},
			"invalid 'type'",
		},
	}

	for _, test := range testCases {
		event := test.f()
		assert.Regexp(t, test.err, validateEvent(&event))
	}
}

func TestDirectionOut(t *testing.T) {
	processor := transProcessor{
		localIPStrings: []string{"192.145.2.4"},
		ignoreOutgoing: false,
		name:           "test",
	}

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type": "test",
			"src": &common.Endpoint{
				IP:     "192.145.2.4",
				Port:   3267,
				Domain: "server1",
				Process: common.Process{
					Args: []string{"proc1", "start"},
					Name: "proc1",
				},
			},
			"dst": &common.Endpoint{
				IP:     "192.145.2.5",
				Port:   32232,
				Domain: "server2",
				Process: common.Process{
					Args: []string{"proc2", "start"},
					Name: "proc2",
				},
			},
		},
	}

	if res, _ := processor.Run(&event); res == nil {
		t.Fatalf("event has been filtered out")
	}
	clientIP, _ := event.GetValue("client.ip")
	assert.Equal(t, "192.145.2.4", clientIP)
	dir, _ := event.GetValue("network.direction")
	assert.Equal(t, "outgoing", dir)
}

func TestDirectionIn(t *testing.T) {
	processor := transProcessor{
		localIPStrings: []string{"192.145.2.5"},
		ignoreOutgoing: false,
		name:           "test",
	}

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type": "test",
			"src": &common.Endpoint{
				IP:     "192.145.2.4",
				Port:   3267,
				Domain: "server1",
				Process: common.Process{
					Args: []string{"proc1", "start"},
					Name: "proc1",
				},
			},
			"dst": &common.Endpoint{
				IP:     "192.145.2.5",
				Port:   32232,
				Domain: "server2",
				Process: common.Process{
					Args: []string{"proc2", "start"},
					Name: "proc2",
				},
			},
		},
	}

	if res, _ := processor.Run(&event); res == nil {
		t.Fatalf("event has been filtered out")
	}
	clientIP, _ := event.GetValue("client.ip")
	assert.Equal(t, "192.145.2.4", clientIP)
	dir, _ := event.GetValue("network.direction")
	assert.Equal(t, "incoming", dir)
}

func TestNoDirection(t *testing.T) {
	processor := transProcessor{
		localIPStrings: []string{"192.145.2.6"},
		ignoreOutgoing: false,
		name:           "test",
	}

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type": "test",
			"src": &common.Endpoint{
				IP:     "192.145.2.4",
				Port:   3267,
				Domain: "server1",
				Process: common.Process{
					Args: []string{"proc1", "start"},
					Name: "proc1",
				},
			},
			"dst": &common.Endpoint{
				IP:     "192.145.2.5",
				Port:   32232,
				Domain: "server2",
				Process: common.Process{
					Args: []string{"proc2", "start"},
					Name: "proc2",
				},
			},
		},
	}

	if res, _ := processor.Run(&event); res == nil {
		t.Fatalf("event has been filtered out")
	}
	clientIP, _ := event.GetValue("client.ip")
	assert.Equal(t, "192.145.2.4", clientIP)
	dir, _ := event.GetValue("network.direction")
	assert.Nil(t, dir)
}
