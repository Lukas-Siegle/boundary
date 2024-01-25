// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scopes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"slices"
	"time"

	"github.com/hashicorp/boundary/api"
)

type Scope struct {
	Id                          string              `json:"id,omitempty"`
	ScopeId                     string              `json:"scope_id,omitempty"`
	Scope                       *ScopeInfo          `json:"scope,omitempty"`
	Name                        string              `json:"name,omitempty"`
	Description                 string              `json:"description,omitempty"`
	CreatedTime                 time.Time           `json:"created_time,omitempty"`
	UpdatedTime                 time.Time           `json:"updated_time,omitempty"`
	Version                     uint32              `json:"version,omitempty"`
	Type                        string              `json:"type,omitempty"`
	PrimaryAuthMethodId         string              `json:"primary_auth_method_id,omitempty"`
	AuthorizedActions           []string            `json:"authorized_actions,omitempty"`
	AuthorizedCollectionActions map[string][]string `json:"authorized_collection_actions,omitempty"`
	StoragePolicyId             string              `json:"storage_policy_id,omitempty"`

	response *api.Response
}

type ScopeReadResult struct {
	Item     *Scope
	response *api.Response
}

func (n ScopeReadResult) GetItem() *Scope {
	return n.Item
}

func (n ScopeReadResult) GetResponse() *api.Response {
	return n.response
}

type ScopeCreateResult = ScopeReadResult
type ScopeUpdateResult = ScopeReadResult

type ScopeDeleteResult struct {
	response *api.Response
}

// GetItem will always be nil for ScopeDeleteResult
func (n ScopeDeleteResult) GetItem() interface{} {
	return nil
}

func (n ScopeDeleteResult) GetResponse() *api.Response {
	return n.response
}

type ScopeListResult struct {
	Items        []*Scope `json:"items,omitempty"`
	EstItemCount uint     `json:"est_item_count,omitempty"`
	RemovedIds   []string `json:"removed_ids,omitempty"`
	ListToken    string   `json:"list_token,omitempty"`
	ResponseType string   `json:"response_type,omitempty"`
	response     *api.Response
}

func (n ScopeListResult) GetItems() []*Scope {
	return n.Items
}

func (n ScopeListResult) GetEstItemCount() uint {
	return n.EstItemCount
}

func (n ScopeListResult) GetRemovedIds() []string {
	return n.RemovedIds
}

func (n ScopeListResult) GetListToken() string {
	return n.ListToken
}

func (n ScopeListResult) GetResponseType() string {
	return n.ResponseType
}

func (n ScopeListResult) GetResponse() *api.Response {
	return n.response
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, scopeId string, opt ...Option) (*ScopeCreateResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts.postMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "POST", "scopes", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(ScopeCreateResult)
	target.Item = new(Scope)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Read(ctx context.Context, id string, opt ...Option) (*ScopeReadResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("scopes/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(ScopeReadResult)
	target.Item = new(Scope)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Update(ctx context.Context, id string, version uint32, opt ...Option) (*ScopeUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("scopes/%s", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(ScopeUpdateResult)
	target.Item = new(Scope)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Delete(ctx context.Context, id string, opt ...Option) (*ScopeDeleteResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("scopes/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &ScopeDeleteResult{
		response: resp,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, scopeId string, opt ...Option) (*ScopeListResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "GET", "scopes", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(ScopeListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	if target.ResponseType == "complete" || target.ResponseType == "" {
		return target, nil
	}
	// If there are more results, automatically fetch the rest of the results.
	// idToIndex keeps a map from the ID of an item to its index in target.Items.
	// This is used to update updated items in-place and remove deleted items
	// from the result after pagination is done.
	idToIndex := map[string]int{}
	for i, item := range target.Items {
		idToIndex[item.Id] = i
	}
	// Removed IDs in the response may contain duplicates,
	// maintain a set to avoid returning duplicates to the user.
	removedIds := map[string]struct{}{}
	for {
		req, err := c.client.NewRequest(ctx, "GET", "scopes", nil, apiOpts...)
		if err != nil {
			return nil, fmt.Errorf("error creating List request: %w", err)
		}

		opts.queryMap["list_token"] = target.ListToken
		if len(opts.queryMap) > 0 {
			q := url.Values{}
			for k, v := range opts.queryMap {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
		}

		resp, err := c.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error performing client request during List call: %w", err)
		}

		page := new(ScopeListResult)
		apiErr, err := resp.Decode(page)
		if err != nil {
			return nil, fmt.Errorf("error decoding List response: %w", err)
		}
		if apiErr != nil {
			return nil, apiErr
		}
		for _, item := range page.Items {
			if i, ok := idToIndex[item.Id]; ok {
				// Item has already been seen at index i, update in-place
				target.Items[i] = item
			} else {
				target.Items = append(target.Items, item)
				idToIndex[item.Id] = len(target.Items) - 1
			}
		}
		for _, removedId := range page.RemovedIds {
			removedIds[removedId] = struct{}{}
		}
		target.EstItemCount = page.EstItemCount
		target.ListToken = page.ListToken
		target.ResponseType = page.ResponseType
		target.response = resp
		if target.ResponseType == "complete" {
			break
		}
	}
	for _, removedId := range target.RemovedIds {
		if i, ok := idToIndex[removedId]; ok {
			// Remove the item at index i without preserving order
			// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
			target.Items[i] = target.Items[len(target.Items)-1]
			target.Items = target.Items[:len(target.Items)-1]
			// Update the index of the last element
			idToIndex[target.Items[i].Id] = i
		}
	}
	for deletedId := range removedIds {
		target.RemovedIds = append(target.RemovedIds, deletedId)
	}
	// Sort to make response deterministic
	slices.Sort(target.RemovedIds)
	// Since we paginated to the end, we can avoid confusion
	// for the user by setting the estimated item count to the
	// length of the items slice. If we don't set this here, it
	// will equal the value returned in the last response, which is
	// often much smaller than the total number returned.
	target.EstItemCount = uint(len(target.Items))
	// Sort the results again since in-place updates and deletes
	// may have shuffled items. We sort by created time descending
	// (most recently created first), same as the API.
	slices.SortFunc(target.Items, func(i, j *Scope) int {
		return j.CreatedTime.Compare(i.CreatedTime)
	})
	// Finally, since we made at least 2 requests to the server to fulfill this
	// function call, resp.Body and resp.Map will only contain the most recent response.
	// Overwrite them with the true response.
	target.response.Body.Reset()
	if err := json.NewEncoder(target.response.Body).Encode(target); err != nil {
		return nil, fmt.Errorf("error encoding final JSON list response: %w", err)
	}
	if err := json.Unmarshal(target.response.Body.Bytes(), &target.response.Map); err != nil {
		return nil, fmt.Errorf("error encoding final map list response: %w", err)
	}
	// Note: the HTTP response body is consumed by resp.Decode in the loop,
	// so it doesn't need to be updated (it will always be, and has always been, empty).
	return target, nil
}
