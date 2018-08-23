// Copyright 2018 AT&T Intellectual Property.  All other rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Interact with the Drydock Bootaction API

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (msg *BootactionMessage) post(url string, key string) error {
	timeout, _ := time.ParseDuration("60s")

	api_request, err := buildRequest(url, key, msg)

	if err != nil {
		return fmt.Errorf("Error build API request: %s", err)
	}

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(api_request)

	if err != nil {
		return fmt.Errorf("Error sending the API request: %s\n", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Error response: %s", resp.Status)
	}

	return nil
}

func buildRequest(url string, key string, msg *BootactionMessage) (*http.Request, error) {
	body, err := json.Marshal(msg)

	if err != nil {
		return nil, fmt.Errorf("Error encoding message: %s\n", err)
	}

	bodyReader := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)

	if err != nil {
		return nil, fmt.Errorf("Error creating API request: %s\n", err)
	}

	req.Header.Add("X-Bootaction-Key", key)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
