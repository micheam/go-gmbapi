package gmbapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/micheam/gmbapi/util"
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
func (a *AccountAccess) List(params url.Values) (*AccountsList, error) {
	_url := "https://mybusiness.googleapis.com/v4/accounts/"
	// FIXME(micheam): hundle NextPageToken
	b, err := a.client.doRequest(http.MethodGet, _url, nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest accounts.list: %w", err)
	}
	var res = new(AccountsList)
	if err := json.Unmarshal(b, res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return res, nil
}

// AccountsList : accounts.list
type AccountsList struct {
	Accounts      []*Account `json:"accounts"`
	NextPageToken string     `json:"nextPageToken"`
}

// Account is a data for Account Resource of Google My Business API.
type Account struct {
	Name             *string                `json:"name"` //  the resource name, in the format 'accounts/{accountId}'.
	AccountName      *string                `json:"accountName"`
	Type             *AccountType           `json:"type"`
	Role             *AccountRole           `json:"role"`
	State            *AccountState          `json:"state"`
	AccountNumber    *string                `json:"accountNumber"`
	PermissionLevel  *PermissionLevel       `json:"permissionLevel"`
	OrganizationInfo *util.OrganizationInfo `json:"organizationInfo"`
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
