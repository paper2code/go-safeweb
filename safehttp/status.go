// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package safehttp

// StatusCode contains HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type StatusCode int

const (
	StatusOK                  StatusCode = 200 // RFC 7231, 6.3.1
	StatusMovedPermanently    StatusCode = 301 // RFC 7231, 6.4.2
	StatusBadRequest          StatusCode = 400 // RFC 7231, 6.5.1
	StatusUnauthorized        StatusCode = 401 // RFC 7231, 3.1
	StatusForbidden           StatusCode = 403 // RFC 7231, 6.5.3
	StatusInternalServerError StatusCode = 500 // RFC 7231, 6.6.1
)
