/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v12+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information abot the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 12
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// Entitlement struct for Entitlement
type Entitlement struct {
	BaseEntity
	// This field is deprecated as of 5.1 in favor of 'appShortcut.name'. For backwards compatibility, it will set 'appShortcut.name' if it does not exist.
	DisplayName string `json:"displayName"`
	// If true, the Entitlement will be disregarded during authorization.
	Disabled *bool `json:"disabled,omitempty"`
	// ID of the site for this Entitlement.
	Site string `json:"site"`
	// Whether all the Conditions must succeed to have access to this Entitlement or just one.
	ConditionLogic *string `json:"conditionLogic,omitempty"`
	// List of Condition IDs applies to this Entitlement.
	Conditions []string `json:"conditions"`
	// List of all IP Access actions in this Entitlement.
	Actions     []EntitlementAllOfActions    `json:"actions"`
	AppShortcut *EntitlementAllOfAppShortcut `json:"appShortcut,omitempty"`
}

// NewEntitlement instantiates a new Entitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlement(displayName string, site string, conditions []string, actions []EntitlementAllOfActions) *Entitlement {
	this := Entitlement{}
	this.DisplayName = displayName
	var disabled bool = false
	this.Disabled = &disabled
	this.Site = site
	var conditionLogic string = "and"
	this.ConditionLogic = &conditionLogic
	this.Conditions = conditions
	this.Actions = actions
	return &this
}

// NewEntitlementWithDefaults instantiates a new Entitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWithDefaults() *Entitlement {
	this := Entitlement{}
	var disabled bool = false
	this.Disabled = &disabled
	var conditionLogic string = "and"
	this.ConditionLogic = &conditionLogic
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *Entitlement) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Entitlement) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Entitlement) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *Entitlement) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Entitlement) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetSite returns the Site field value
func (o *Entitlement) GetSite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *Entitlement) SetSite(v string) {
	o.Site = v
}

// GetConditionLogic returns the ConditionLogic field value if set, zero value otherwise.
func (o *Entitlement) GetConditionLogic() string {
	if o == nil || o.ConditionLogic == nil {
		var ret string
		return ret
	}
	return *o.ConditionLogic
}

// GetConditionLogicOk returns a tuple with the ConditionLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetConditionLogicOk() (*string, bool) {
	if o == nil || o.ConditionLogic == nil {
		return nil, false
	}
	return o.ConditionLogic, true
}

// HasConditionLogic returns a boolean if a field has been set.
func (o *Entitlement) HasConditionLogic() bool {
	if o != nil && o.ConditionLogic != nil {
		return true
	}

	return false
}

// SetConditionLogic gets a reference to the given string and assigns it to the ConditionLogic field.
func (o *Entitlement) SetConditionLogic(v string) {
	o.ConditionLogic = &v
}

// GetConditions returns the Conditions field value
func (o *Entitlement) GetConditions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetConditionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value
func (o *Entitlement) SetConditions(v []string) {
	o.Conditions = v
}

// GetActions returns the Actions field value
func (o *Entitlement) GetActions() []EntitlementAllOfActions {
	if o == nil {
		var ret []EntitlementAllOfActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetActionsOk() (*[]EntitlementAllOfActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *Entitlement) SetActions(v []EntitlementAllOfActions) {
	o.Actions = v
}

// GetAppShortcut returns the AppShortcut field value if set, zero value otherwise.
func (o *Entitlement) GetAppShortcut() EntitlementAllOfAppShortcut {
	if o == nil || o.AppShortcut == nil {
		var ret EntitlementAllOfAppShortcut
		return ret
	}
	return *o.AppShortcut
}

// GetAppShortcutOk returns a tuple with the AppShortcut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetAppShortcutOk() (*EntitlementAllOfAppShortcut, bool) {
	if o == nil || o.AppShortcut == nil {
		return nil, false
	}
	return o.AppShortcut, true
}

// HasAppShortcut returns a boolean if a field has been set.
func (o *Entitlement) HasAppShortcut() bool {
	if o != nil && o.AppShortcut != nil {
		return true
	}

	return false
}

// SetAppShortcut gets a reference to the given EntitlementAllOfAppShortcut and assigns it to the AppShortcut field.
func (o *Entitlement) SetAppShortcut(v EntitlementAllOfAppShortcut) {
	o.AppShortcut = &v
}

func (o Entitlement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEntity, errBaseEntity := json.Marshal(o.BaseEntity)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	errBaseEntity = json.Unmarshal([]byte(serializedBaseEntity), &toSerialize)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if true {
		toSerialize["site"] = o.Site
	}
	if o.ConditionLogic != nil {
		toSerialize["conditionLogic"] = o.ConditionLogic
	}
	if true {
		toSerialize["conditions"] = o.Conditions
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if o.AppShortcut != nil {
		toSerialize["appShortcut"] = o.AppShortcut
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlement struct {
	value *Entitlement
	isSet bool
}

func (v NullableEntitlement) Get() *Entitlement {
	return v.value
}

func (v *NullableEntitlement) Set(val *Entitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlement(val *Entitlement) *NullableEntitlement {
	return &NullableEntitlement{value: val, isSet: true}
}

func (v NullableEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
