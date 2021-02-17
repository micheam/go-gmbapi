package gmbapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// AccountAccess ...
type AccountAccess struct {
	client *Client
}

// AccountAccess ...
func (c *Client) AccountAccess() *AccountAccess {
	return &AccountAccess{client: c}
}

// List ...
func (a *AccountAccess) List(ctx context.Context, params url.Values) ([]*Account, error) {
	var list = make([]*Account, 0)
	var next *string = nil
	for {
		result, err := a.list(ctx, next, params)
		if err != nil {
			return nil, err
		}
		list = append(list, result.Accounts...)
		next = result.NextPageToken
		if next == nil {
			break
		}
	}
	return list, nil
}

// Get return the specified account. Returns ErrNotFound if the
// account does not exist or if the caller does not have access rights to it.
//
// name : 'accounts/${account_id}'
func (a *AccountAccess) Get(ctx context.Context, name string) (*Account, error) {
	// TODO(micheam): QPS Limit with "golang.org/x/time/rate"
	b, err := a.client.doRequest(
		ctx, time.Now(), http.MethodGet,
		strings.Join([]string{BaseEndpoint, name}, "/"), nil, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest accounts.get: %w", err)
	}
	var acc = new(Account)
	if err := json.Unmarshal(b, acc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return acc, nil
}

func (a *AccountAccess) list(ctx context.Context, nextPageToken *string, params url.Values) (*AccountList, error) {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	if nextPageToken != nil {
		params.Add("pageToken", *nextPageToken)
	}
	b, err := a.client.doRequest(ctx, time.Now(), http.MethodGet, BaseEndpoint+"/accounts/", nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest accounts.list: %w", err)
	}
	var dat = new(AccountList)
	if err := json.Unmarshal(b, dat); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return dat, nil
}

// AccountList : accounts.list
type AccountList struct {
	Accounts      []*Account `json:"accounts"`
	NextPageToken *string    `json:"nextPageToken"`
}

// AccountID is a identifier of account
type AccountID string

// Account is a data for Account Resource of Google My Business API.
type Account struct {
	Name             string           `json:"name"` //  the resource name, in the format 'accounts/{accountId}'.
	AccountName      string           `json:"accountName"`
	Type             AccountType      `json:"type"`
	Role             AccountRole      `json:"role"`
	State            AccountState     `json:"state"`
	AccountNumber    string           `json:"accountNumber"`
	PermissionLevel  PermissionLevel  `json:"permissionLevel"`
	OrganizationInfo OrganizationInfo `json:"organizationInfo"`
}

func (a *Account) ID() AccountID {
	s := strings.Split(a.Name, "/")
	return AccountID(s[len(s)-1])
}

var ErrInvalidAccountName = errors.New("invalid account name")

// TODO(micheam): move under package gmbapi
func ParseAccountName(s string) (accountName string, err error) {
	ss := strings.Split(s, "/")
	if len(ss) > 0 {
		if ss[0] != "accounts" {
			return "", ErrInvalidAccountName
		}
		return s, nil
	}
	return fmt.Sprintf("accounts/%s", s), nil
}

/*
 * AccountType
 */

// AccountType ...
type AccountType int

// definition of AccountTypes
const (
	AccountTypeUnspecified AccountType = iota
	AccountTypePersonal
	AccountTypeLocationGroup
	AccountTypeUserGroup
	AccountTypeOrganization
)

func (a AccountType) String() string {
	return [...]string{
		"ACCOUNT_TYPE_UNSPECIFIED",
		"PERSONAL",
		"LOCATION_GROUP",
		"USER_GROUP",
		"ORGANIZATION",
	}[a]
}

// UnmarshalJSON ...
func (a *AccountType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = AccountTypeUnspecified
	case "personal":
		*a = AccountTypePersonal
	case "location_group":
		*a = AccountTypeLocationGroup
	case "user_group":
		*a = AccountTypeUserGroup
	case "organization":
		*a = AccountTypeOrganization
	}
	return nil
}

// MarshalJSON ...
func (a *AccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

/*
 * AccountRole
 */

// AccountRole ...
type AccountRole int

// Definition of AccountRole
const (
	AccountRoleUnspecified AccountRole = iota
	AccountRoleOwner
	AccountRoleCoOwner
	AccountRoleManager
	AccountRoleCommunityManager
)

func (a AccountRole) String() string {
	return [...]string{
		"ACCOUNT_ROLE_UNSPECIFIED",
		"OWNER",
		"CO_OWNER",
		"MANAGER",
		"COMMUNITY_MANAGER",
	}[a]
}

// UnmarshalJSON ...
func (a *AccountRole) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = AccountRoleUnspecified
	case "owner":
		*a = AccountRoleOwner
	case "co_owner":
		*a = AccountRoleCoOwner
	case "manager":
		*a = AccountRoleManager
	case "community_manager":
		*a = AccountRoleCommunityManager
	}
	return nil
}

// MarshalJSON ...
func (a *AccountRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

/*
 * AccountState
 */

// AccountState  ...
type AccountState int

// Definition of AccountState
const (
	AccountStateUnspecified AccountState = iota
	AccountStateVerified
	AccountStateUnverified
	AccountStateVerificationRequested
)

func (a AccountState) String() string {
	return [...]string{
		"ACCOUNT_STATUS_UNSPECIFIED",
		"VERIFIED",
		"UNVERIFIED",
		"VERIFICATION_REQUESTED",
	}[a]
}

// UnmarshalJSON ...
func (a *AccountState) UnmarshalJSON(b []byte) error {
	var m map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	switch strings.ToLower(m["state"]) {
	default:
		*a = AccountStateUnspecified
	case "verified":
		*a = AccountStateVerified
	case "unverified":
		*a = AccountStateUnverified
	case "verification_requested":
		*a = AccountStateVerificationRequested
	}
	return nil
}

// MarshalJSON ...
func (a *AccountState) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"state": a.String(),
	})
}

/*
 * PermissionLevel
 */

// PermissionLevel  ...
type PermissionLevel int

// Definition of PermissionLevel
const (
	PermissionLevelUnspecified PermissionLevel = iota
	PermissionLevelOwnerLevel
	PermissionLevelMemberLevel
)

func (p PermissionLevel) String() string {
	return [...]string{
		"PERMISSION_LEVEL_UNSPECIFIED",
		"OWNER_LEVEL",
		"MEMBER_LEVEL",
	}[p]
}

// UnmarshalJSON ...
func (p *PermissionLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*p = PermissionLevelUnspecified
	case "owner_level":
		*p = PermissionLevelOwnerLevel
	case "member_level":
		*p = PermissionLevelMemberLevel
	}
	return nil
}

// MarshalJSON ...
func (p *PermissionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}
