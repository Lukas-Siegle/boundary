// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apptokens

type Grant struct {
	Raw       string     `json:"raw,omitempty"`
	Canonical string     `json:"canonical,omitempty"`
	Json      *GrantJson `json:"json,omitempty"`
}
