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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
