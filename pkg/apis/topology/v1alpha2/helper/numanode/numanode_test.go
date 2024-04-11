/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package numanode

import (
	"testing"
)

func TestIDToName(t *testing.T) {
	testCases := []struct {
		name          string
		cellID        int
		expected      string
		expectedError bool
	}{
		{
			name:     "zero",
			cellID:   0,
			expected: "node-0",
		},
		{
			name:     "positive",
			cellID:   7,
			expected: "node-7",
		},
		{
			name:     "excessive",
			cellID:   65535,
			expected: "node-65535",
		},
		{
			name:          "negative",
			cellID:        -1,
			expected:      "",
			expectedError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IDToName(tt.cellID)
			gotErr := (err != nil)
			if gotErr != tt.expectedError {
				t.Fatalf("expected error=%v got=%v", tt.expectedError, gotErr)
			}
			if got != tt.expected {
				t.Fatalf("expected value %v got value %v", tt.expected, got)
			}
		})
	}
}

func TestNameToID(t *testing.T) {
	testCases := []struct {
		name          string
		cellName      string
		expected      int
		expectedError bool
	}{
		{
			name:     "zero",
			cellName: "node-0",
			expected: 0,
		},
		{
			name:     "positive",
			cellName: "node-7",
			expected: 7,
		},
		{
			name:     "excessive",
			cellName: "node-65535",
			expected: 65535,
		},
		{
			name:          "negative",
			cellName:      "node--1",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "wrong node id",
			cellName:      "cell-1",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "case node id (1)",
			cellName:      "Node-1",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "case node id (2)",
			cellName:      "NODE-2",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "case node id (3)",
			cellName:      "node_3",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "case node id (4)",
			cellName:      "node 4",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "case node id (5)",
			cellName:      "node-5-1",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "case node id (6)",
			cellName:      "node6",
			expected:      -1,
			expectedError: true,
		},
		{
			name:          "case node id (7)",
			cellName:      "--node-7",
			expected:      -1,
			expectedError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NameToID(tt.cellName)
			gotErr := (err != nil)
			if gotErr != tt.expectedError {
				t.Fatalf("expected error=%v got=%v", tt.expectedError, gotErr)
			}
			if got != tt.expected {
				t.Fatalf("expected value %v got value %v", tt.expected, got)
			}
		})
	}
}
