﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package identity

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azuredevops"
    "github.com/microsoft/azure-devops-go-api/azuredevops/delegatedauthorization"
    "github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("8a3d49b8-91f0-46ef-b33d-dda338c25db3")

type Client struct {
    Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client{
        Client: *client,
    }, nil
}

// [Preview API]
func (client *Client) CreateOrBindWithClaims(ctx context.Context, args CreateOrBindWithClaimsArgs) (*Identity, error) {
    if args.SourceIdentity == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "sourceIdentity"}
    }
    body, marshalErr := json.Marshal(*args.SourceIdentity)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("90ddfe71-171c-446c-bf3b-b597cd562afd")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Identity
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateOrBindWithClaims function
type CreateOrBindWithClaimsArgs struct {
    // (required)
    SourceIdentity *Identity
}

// [Preview API]
func (client *Client) GetDescriptorById(ctx context.Context, args GetDescriptorByIdArgs) (*string, error) {
    routeValues := make(map[string]string)
    if args.Id == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = (*args.Id).String()

    queryParams := url.Values{}
    if args.IsMasterId != nil {
        queryParams.Add("isMasterId", strconv.FormatBool(*args.IsMasterId))
    }
    locationId, _ := uuid.Parse("a230389a-94f2-496c-839f-c929787496dd")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetDescriptorById function
type GetDescriptorByIdArgs struct {
    // (required)
    Id *uuid.UUID
    // (optional)
    IsMasterId *bool
}

func (client *Client) CreateGroups(ctx context.Context, args CreateGroupsArgs) (*[]Identity, error) {
    if args.Container == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "container"}
    }
    body, marshalErr := json.Marshal(args.Container)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5966283b-4196-4d57-9211-1b68f41ec1c2")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateGroups function
type CreateGroupsArgs struct {
    // (required)
    Container interface{}
}

func (client *Client) DeleteGroup(ctx context.Context, args DeleteGroupArgs) error {
    routeValues := make(map[string]string)
    if args.GroupId == nil || *args.GroupId == "" {
        return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *args.GroupId

    locationId, _ := uuid.Parse("5966283b-4196-4d57-9211-1b68f41ec1c2")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteGroup function
type DeleteGroupArgs struct {
    // (required)
    GroupId *string
}

func (client *Client) ListGroups(ctx context.Context, args ListGroupsArgs) (*[]Identity, error) {
    queryParams := url.Values{}
    if args.ScopeIds != nil {
        queryParams.Add("scopeIds", *args.ScopeIds)
    }
    if args.Recurse != nil {
        queryParams.Add("recurse", strconv.FormatBool(*args.Recurse))
    }
    if args.Deleted != nil {
        queryParams.Add("deleted", strconv.FormatBool(*args.Deleted))
    }
    if args.Properties != nil {
        queryParams.Add("properties", *args.Properties)
    }
    locationId, _ := uuid.Parse("5966283b-4196-4d57-9211-1b68f41ec1c2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ListGroups function
type ListGroupsArgs struct {
    // (optional)
    ScopeIds *string
    // (optional)
    Recurse *bool
    // (optional)
    Deleted *bool
    // (optional)
    Properties *string
}

func (client *Client) GetIdentityChanges(ctx context.Context, args GetIdentityChangesArgs) (*ChangedIdentities, error) {
    queryParams := url.Values{}
    if args.IdentitySequenceId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "identitySequenceId"}
    }
    queryParams.Add("identitySequenceId", strconv.Itoa(*args.IdentitySequenceId))
    if args.GroupSequenceId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "groupSequenceId"}
    }
    queryParams.Add("groupSequenceId", strconv.Itoa(*args.GroupSequenceId))
    if args.OrganizationIdentitySequenceId != nil {
        queryParams.Add("organizationIdentitySequenceId", strconv.Itoa(*args.OrganizationIdentitySequenceId))
    }
    if args.PageSize != nil {
        queryParams.Add("pageSize", strconv.Itoa(*args.PageSize))
    }
    if args.ScopeId != nil {
        queryParams.Add("scopeId", (*args.ScopeId).String())
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ChangedIdentities
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetIdentityChanges function
type GetIdentityChangesArgs struct {
    // (required)
    IdentitySequenceId *int
    // (required)
    GroupSequenceId *int
    // (optional)
    OrganizationIdentitySequenceId *int
    // (optional)
    PageSize *int
    // (optional)
    ScopeId *uuid.UUID
}

func (client *Client) GetUserIdentityIdsByDomainId(ctx context.Context, args GetUserIdentityIdsByDomainIdArgs) (*[]uuid.UUID, error) {
    queryParams := url.Values{}
    if args.DomainId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "domainId"}
    }
    queryParams.Add("domainId", (*args.DomainId).String())
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []uuid.UUID
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetUserIdentityIdsByDomainId function
type GetUserIdentityIdsByDomainIdArgs struct {
    // (required)
    DomainId *uuid.UUID
}

func (client *Client) ReadIdentities(ctx context.Context, args ReadIdentitiesArgs) (*[]Identity, error) {
    queryParams := url.Values{}
    if args.Descriptors != nil {
        queryParams.Add("descriptors", *args.Descriptors)
    }
    if args.IdentityIds != nil {
        queryParams.Add("identityIds", *args.IdentityIds)
    }
    if args.SubjectDescriptors != nil {
        queryParams.Add("subjectDescriptors", *args.SubjectDescriptors)
    }
    if args.SocialDescriptors != nil {
        queryParams.Add("socialDescriptors", *args.SocialDescriptors)
    }
    if args.SearchFilter != nil {
        queryParams.Add("searchFilter", *args.SearchFilter)
    }
    if args.FilterValue != nil {
        queryParams.Add("filterValue", *args.FilterValue)
    }
    if args.QueryMembership != nil {
        queryParams.Add("queryMembership", string(*args.QueryMembership))
    }
    if args.Properties != nil {
        queryParams.Add("properties", *args.Properties)
    }
    if args.IncludeRestrictedVisibility != nil {
        queryParams.Add("includeRestrictedVisibility", strconv.FormatBool(*args.IncludeRestrictedVisibility))
    }
    if args.Options != nil {
        queryParams.Add("options", string(*args.Options))
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadIdentities function
type ReadIdentitiesArgs struct {
    // (optional)
    Descriptors *string
    // (optional)
    IdentityIds *string
    // (optional)
    SubjectDescriptors *string
    // (optional)
    SocialDescriptors *string
    // (optional)
    SearchFilter *string
    // (optional)
    FilterValue *string
    // (optional)
    QueryMembership *QueryMembership
    // (optional)
    Properties *string
    // (optional)
    IncludeRestrictedVisibility *bool
    // (optional)
    Options *ReadIdentitiesOptions
}

func (client *Client) ReadIdentitiesByScope(ctx context.Context, args ReadIdentitiesByScopeArgs) (*[]Identity, error) {
    queryParams := url.Values{}
    if args.ScopeId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "scopeId"}
    }
    queryParams.Add("scopeId", (*args.ScopeId).String())
    if args.QueryMembership != nil {
        queryParams.Add("queryMembership", string(*args.QueryMembership))
    }
    if args.Properties != nil {
        queryParams.Add("properties", *args.Properties)
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadIdentitiesByScope function
type ReadIdentitiesByScopeArgs struct {
    // (required)
    ScopeId *uuid.UUID
    // (optional)
    QueryMembership *QueryMembership
    // (optional)
    Properties *string
}

func (client *Client) ReadIdentity(ctx context.Context, args ReadIdentityArgs) (*Identity, error) {
    routeValues := make(map[string]string)
    if args.IdentityId == nil || *args.IdentityId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "identityId"} 
    }
    routeValues["identityId"] = *args.IdentityId

    queryParams := url.Values{}
    if args.QueryMembership != nil {
        queryParams.Add("queryMembership", string(*args.QueryMembership))
    }
    if args.Properties != nil {
        queryParams.Add("properties", *args.Properties)
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Identity
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadIdentity function
type ReadIdentityArgs struct {
    // (required)
    IdentityId *string
    // (optional)
    QueryMembership *QueryMembership
    // (optional)
    Properties *string
}

func (client *Client) UpdateIdentities(ctx context.Context, args UpdateIdentitiesArgs) (*[]IdentityUpdateData, error) {
    if args.Identities == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "identities"}
    }
    body, marshalErr := json.Marshal(*args.Identities)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []IdentityUpdateData
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateIdentities function
type UpdateIdentitiesArgs struct {
    // (required)
    Identities *azuredevops.VssJsonCollectionWrapper
}

func (client *Client) UpdateIdentity(ctx context.Context, args UpdateIdentityArgs) error {
    if args.Identity == nil {
        return &azuredevops.ArgumentNilError{ArgumentName: "identity"}
    }
    routeValues := make(map[string]string)
    if args.IdentityId == nil {
        return &azuredevops.ArgumentNilError{ArgumentName: "identityId"} 
    }
    routeValues["identityId"] = (*args.IdentityId).String()

    body, marshalErr := json.Marshal(*args.Identity)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    _, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the UpdateIdentity function
type UpdateIdentityArgs struct {
    // (required)
    Identity *Identity
    // (required)
    IdentityId *uuid.UUID
}

func (client *Client) CreateIdentity(ctx context.Context, args CreateIdentityArgs) (*Identity, error) {
    if args.FrameworkIdentityInfo == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "frameworkIdentityInfo"}
    }
    body, marshalErr := json.Marshal(*args.FrameworkIdentityInfo)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("dd55f0eb-6ea2-4fe4-9ebe-919e7dd1dfb4")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Identity
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateIdentity function
type CreateIdentityArgs struct {
    // (required)
    FrameworkIdentityInfo *FrameworkIdentityInfo
}

// [Preview API]
func (client *Client) ReadIdentityBatch(ctx context.Context, args ReadIdentityBatchArgs) (*[]Identity, error) {
    if args.BatchInfo == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "batchInfo"}
    }
    body, marshalErr := json.Marshal(*args.BatchInfo)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("299e50df-fe45-4d3a-8b5b-a5836fac74dc")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadIdentityBatch function
type ReadIdentityBatchArgs struct {
    // (required)
    BatchInfo *IdentityBatchInfo
}

// [Preview API]
func (client *Client) GetIdentitySnapshot(ctx context.Context, args GetIdentitySnapshotArgs) (*IdentitySnapshot, error) {
    routeValues := make(map[string]string)
    if args.ScopeId == nil || *args.ScopeId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "scopeId"} 
    }
    routeValues["scopeId"] = *args.ScopeId

    locationId, _ := uuid.Parse("d56223df-8ccd-45c9-89b4-eddf692400d7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentitySnapshot
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetIdentitySnapshot function
type GetIdentitySnapshotArgs struct {
    // (required)
    ScopeId *string
}

// Read the max sequence id of all the identities.
func (client *Client) GetMaxSequenceId(ctx context.Context, args GetMaxSequenceIdArgs) (*uint64, error) {
    locationId, _ := uuid.Parse("e4a70778-cb2c-4e85-b7cc-3f3c7ae2d408")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue uint64
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetMaxSequenceId function
type GetMaxSequenceIdArgs struct {
}

// Read identity of the home tenant request user.
func (client *Client) GetSelf(ctx context.Context, args GetSelfArgs) (*IdentitySelf, error) {
    locationId, _ := uuid.Parse("4bb02b5b-c120-4be2-b68e-21f7c50a4b82")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentitySelf
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetSelf function
type GetSelfArgs struct {
}

// [Preview API]
func (client *Client) AddMember(ctx context.Context, args AddMemberArgs) (*bool, error) {
    routeValues := make(map[string]string)
    if args.ContainerId == nil || *args.ContainerId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = *args.ContainerId
    if args.MemberId == nil || *args.MemberId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "memberId"} 
    }
    routeValues["memberId"] = *args.MemberId

    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the AddMember function
type AddMemberArgs struct {
    // (required)
    ContainerId *string
    // (required)
    MemberId *string
}

// [Preview API]
func (client *Client) ReadMember(ctx context.Context, args ReadMemberArgs) (*string, error) {
    routeValues := make(map[string]string)
    if args.ContainerId == nil || *args.ContainerId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = *args.ContainerId
    if args.MemberId == nil || *args.MemberId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "memberId"} 
    }
    routeValues["memberId"] = *args.MemberId

    queryParams := url.Values{}
    if args.QueryMembership != nil {
        queryParams.Add("queryMembership", string(*args.QueryMembership))
    }
    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadMember function
type ReadMemberArgs struct {
    // (required)
    ContainerId *string
    // (required)
    MemberId *string
    // (optional)
    QueryMembership *QueryMembership
}

// [Preview API]
func (client *Client) ReadMembers(ctx context.Context, args ReadMembersArgs) (*[]string, error) {
    routeValues := make(map[string]string)
    if args.ContainerId == nil || *args.ContainerId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = *args.ContainerId

    queryParams := url.Values{}
    if args.QueryMembership != nil {
        queryParams.Add("queryMembership", string(*args.QueryMembership))
    }
    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadMembers function
type ReadMembersArgs struct {
    // (required)
    ContainerId *string
    // (optional)
    QueryMembership *QueryMembership
}

// [Preview API]
func (client *Client) RemoveMember(ctx context.Context, args RemoveMemberArgs) (*bool, error) {
    routeValues := make(map[string]string)
    if args.ContainerId == nil || *args.ContainerId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = *args.ContainerId
    if args.MemberId == nil || *args.MemberId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "memberId"} 
    }
    routeValues["memberId"] = *args.MemberId

    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the RemoveMember function
type RemoveMemberArgs struct {
    // (required)
    ContainerId *string
    // (required)
    MemberId *string
}

// [Preview API]
func (client *Client) ReadMemberOf(ctx context.Context, args ReadMemberOfArgs) (*string, error) {
    routeValues := make(map[string]string)
    if args.MemberId == nil || *args.MemberId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "memberId"} 
    }
    routeValues["memberId"] = *args.MemberId
    if args.ContainerId == nil || *args.ContainerId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = *args.ContainerId

    queryParams := url.Values{}
    if args.QueryMembership != nil {
        queryParams.Add("queryMembership", string(*args.QueryMembership))
    }
    locationId, _ := uuid.Parse("22865b02-9e4a-479e-9e18-e35b8803b8a0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadMemberOf function
type ReadMemberOfArgs struct {
    // (required)
    MemberId *string
    // (required)
    ContainerId *string
    // (optional)
    QueryMembership *QueryMembership
}

// [Preview API]
func (client *Client) ReadMembersOf(ctx context.Context, args ReadMembersOfArgs) (*[]string, error) {
    routeValues := make(map[string]string)
    if args.MemberId == nil || *args.MemberId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "memberId"} 
    }
    routeValues["memberId"] = *args.MemberId

    queryParams := url.Values{}
    if args.QueryMembership != nil {
        queryParams.Add("queryMembership", string(*args.QueryMembership))
    }
    locationId, _ := uuid.Parse("22865b02-9e4a-479e-9e18-e35b8803b8a0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the ReadMembersOf function
type ReadMembersOfArgs struct {
    // (required)
    MemberId *string
    // (optional)
    QueryMembership *QueryMembership
}

// [Preview API]
func (client *Client) CreateScope(ctx context.Context, args CreateScopeArgs) (*IdentityScope, error) {
    if args.Info == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "info"}
    }
    routeValues := make(map[string]string)
    if args.ScopeId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "scopeId"} 
    }
    routeValues["scopeId"] = (*args.ScopeId).String()

    body, marshalErr := json.Marshal(*args.Info)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityScope
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateScope function
type CreateScopeArgs struct {
    // (required)
    Info *CreateScopeInfo
    // (required)
    ScopeId *uuid.UUID
}

// [Preview API]
func (client *Client) DeleteScope(ctx context.Context, args DeleteScopeArgs) error {
    routeValues := make(map[string]string)
    if args.ScopeId == nil {
        return &azuredevops.ArgumentNilError{ArgumentName: "scopeId"} 
    }
    routeValues["scopeId"] = (*args.ScopeId).String()

    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteScope function
type DeleteScopeArgs struct {
    // (required)
    ScopeId *uuid.UUID
}

// [Preview API]
func (client *Client) GetScopeById(ctx context.Context, args GetScopeByIdArgs) (*IdentityScope, error) {
    routeValues := make(map[string]string)
    if args.ScopeId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "scopeId"} 
    }
    routeValues["scopeId"] = (*args.ScopeId).String()

    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityScope
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetScopeById function
type GetScopeByIdArgs struct {
    // (required)
    ScopeId *uuid.UUID
}

// [Preview API]
func (client *Client) GetScopeByName(ctx context.Context, args GetScopeByNameArgs) (*IdentityScope, error) {
    queryParams := url.Values{}
    if args.ScopeName == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "scopeName"}
    }
    queryParams.Add("scopeName", *args.ScopeName)
    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityScope
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetScopeByName function
type GetScopeByNameArgs struct {
    // (required)
    ScopeName *string
}

// [Preview API]
func (client *Client) UpdateScope(ctx context.Context, args UpdateScopeArgs) error {
    if args.PatchDocument == nil {
        return &azuredevops.ArgumentNilError{ArgumentName: "patchDocument"}
    }
    routeValues := make(map[string]string)
    if args.ScopeId == nil {
        return &azuredevops.ArgumentNilError{ArgumentName: "scopeId"} 
    }
    routeValues["scopeId"] = (*args.ScopeId).String()

    body, marshalErr := json.Marshal(*args.PatchDocument)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the UpdateScope function
type UpdateScopeArgs struct {
    // (required)
    PatchDocument *[]webapi.JsonPatchOperation
    // (required)
    ScopeId *uuid.UUID
}

// [Preview API]
func (client *Client) GetSignedInToken(ctx context.Context, args GetSignedInTokenArgs) (*delegatedauthorization.AccessTokenResult, error) {
    locationId, _ := uuid.Parse("6074ff18-aaad-4abb-a41e-5c75f6178057")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue delegatedauthorization.AccessTokenResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetSignedInToken function
type GetSignedInTokenArgs struct {
}

// [Preview API]
func (client *Client) GetSignoutToken(ctx context.Context, args GetSignoutTokenArgs) (*delegatedauthorization.AccessTokenResult, error) {
    locationId, _ := uuid.Parse("be39e83c-7529-45e9-9c67-0410885880da")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue delegatedauthorization.AccessTokenResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetSignoutToken function
type GetSignoutTokenArgs struct {
}

// [Preview API]
func (client *Client) GetTenant(ctx context.Context, args GetTenantArgs) (*TenantInfo, error) {
    routeValues := make(map[string]string)
    if args.TenantId == nil || *args.TenantId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "tenantId"} 
    }
    routeValues["tenantId"] = *args.TenantId

    locationId, _ := uuid.Parse("5f0a1723-2e2c-4c31-8cae-002d01bdd592")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TenantInfo
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTenant function
type GetTenantArgs struct {
    // (required)
    TenantId *string
}
