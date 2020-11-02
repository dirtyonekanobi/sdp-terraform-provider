/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v13+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 13
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// LdapProvider struct for LdapProvider
type LdapProvider struct {
	IdentityProvider
	// Hostnames/IP addresses to connect.
	Hostnames []string `json:"hostnames"`
	// Port to connect.
	Port int32 `json:"port"`
	// Whether to use LDAPS protocol or not.
	SslEnabled *bool `json:"sslEnabled,omitempty"`
	// The Distinguished Name to login to LDAP and query users with.
	AdminDistinguishedName string `json:"adminDistinguishedName"`
	// The password to login to LDAP and query users with. Required on creation.
	AdminPassword *string `json:"adminPassword,omitempty"`
	// The subset of the LDAP server to search users from. If not set, root of the server is used.
	BaseDn *string `json:"baseDn,omitempty"`
	// The object class of the users to be authenticated and queried.
	ObjectClass *string `json:"objectClass,omitempty"`
	// The name of the attribute to get the exact username from the LDAP server.
	UsernameAttribute *string `json:"usernameAttribute,omitempty"`
	// The filter to use while querying users' nested groups.
	MembershipFilter *string `json:"membershipFilter,omitempty"`
	// The subset of the LDAP server to search groups from. If not set, \"baseDn\" is used.
	MembershipBaseDn *string                           `json:"membershipBaseDn,omitempty"`
	PasswordWarning  *LdapProviderAllOfPasswordWarning `json:"passwordWarning,omitempty"`
}

// NewLdapProvider instantiates a new LdapProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProvider(hostnames []string, port int32, adminDistinguishedName string) *LdapProvider {
	this := LdapProvider{}
	this.Hostnames = hostnames
	this.Port = port
	var sslEnabled bool = false
	this.SslEnabled = &sslEnabled
	this.AdminDistinguishedName = adminDistinguishedName
	var objectClass string = "user"
	this.ObjectClass = &objectClass
	var usernameAttribute string = "sAMAccountName"
	this.UsernameAttribute = &usernameAttribute
	var membershipFilter string = "(objectCategory=group)"
	this.MembershipFilter = &membershipFilter
	return &this
}

// NewLdapProviderWithDefaults instantiates a new LdapProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderWithDefaults() *LdapProvider {
	this := LdapProvider{}
	var sslEnabled bool = false
	this.SslEnabled = &sslEnabled
	var objectClass string = "user"
	this.ObjectClass = &objectClass
	var usernameAttribute string = "sAMAccountName"
	this.UsernameAttribute = &usernameAttribute
	var membershipFilter string = "(objectCategory=group)"
	this.MembershipFilter = &membershipFilter
	return &this
}

// GetHostnames returns the Hostnames field value
func (o *LdapProvider) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetHostnamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostnames, true
}

// SetHostnames sets field value
func (o *LdapProvider) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetPort returns the Port field value
func (o *LdapProvider) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *LdapProvider) SetPort(v int32) {
	o.Port = v
}

// GetSslEnabled returns the SslEnabled field value if set, zero value otherwise.
func (o *LdapProvider) GetSslEnabled() bool {
	if o == nil || o.SslEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SslEnabled
}

// GetSslEnabledOk returns a tuple with the SslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetSslEnabledOk() (*bool, bool) {
	if o == nil || o.SslEnabled == nil {
		return nil, false
	}
	return o.SslEnabled, true
}

// HasSslEnabled returns a boolean if a field has been set.
func (o *LdapProvider) HasSslEnabled() bool {
	if o != nil && o.SslEnabled != nil {
		return true
	}

	return false
}

// SetSslEnabled gets a reference to the given bool and assigns it to the SslEnabled field.
func (o *LdapProvider) SetSslEnabled(v bool) {
	o.SslEnabled = &v
}

// GetAdminDistinguishedName returns the AdminDistinguishedName field value
func (o *LdapProvider) GetAdminDistinguishedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdminDistinguishedName
}

// GetAdminDistinguishedNameOk returns a tuple with the AdminDistinguishedName field value
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetAdminDistinguishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminDistinguishedName, true
}

// SetAdminDistinguishedName sets field value
func (o *LdapProvider) SetAdminDistinguishedName(v string) {
	o.AdminDistinguishedName = v
}

// GetAdminPassword returns the AdminPassword field value if set, zero value otherwise.
func (o *LdapProvider) GetAdminPassword() string {
	if o == nil || o.AdminPassword == nil {
		var ret string
		return ret
	}
	return *o.AdminPassword
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetAdminPasswordOk() (*string, bool) {
	if o == nil || o.AdminPassword == nil {
		return nil, false
	}
	return o.AdminPassword, true
}

// HasAdminPassword returns a boolean if a field has been set.
func (o *LdapProvider) HasAdminPassword() bool {
	if o != nil && o.AdminPassword != nil {
		return true
	}

	return false
}

// SetAdminPassword gets a reference to the given string and assigns it to the AdminPassword field.
func (o *LdapProvider) SetAdminPassword(v string) {
	o.AdminPassword = &v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *LdapProvider) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *LdapProvider) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *LdapProvider) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *LdapProvider) GetObjectClass() string {
	if o == nil || o.ObjectClass == nil {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetObjectClassOk() (*string, bool) {
	if o == nil || o.ObjectClass == nil {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *LdapProvider) HasObjectClass() bool {
	if o != nil && o.ObjectClass != nil {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *LdapProvider) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetUsernameAttribute returns the UsernameAttribute field value if set, zero value otherwise.
func (o *LdapProvider) GetUsernameAttribute() string {
	if o == nil || o.UsernameAttribute == nil {
		var ret string
		return ret
	}
	return *o.UsernameAttribute
}

// GetUsernameAttributeOk returns a tuple with the UsernameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetUsernameAttributeOk() (*string, bool) {
	if o == nil || o.UsernameAttribute == nil {
		return nil, false
	}
	return o.UsernameAttribute, true
}

// HasUsernameAttribute returns a boolean if a field has been set.
func (o *LdapProvider) HasUsernameAttribute() bool {
	if o != nil && o.UsernameAttribute != nil {
		return true
	}

	return false
}

// SetUsernameAttribute gets a reference to the given string and assigns it to the UsernameAttribute field.
func (o *LdapProvider) SetUsernameAttribute(v string) {
	o.UsernameAttribute = &v
}

// GetMembershipFilter returns the MembershipFilter field value if set, zero value otherwise.
func (o *LdapProvider) GetMembershipFilter() string {
	if o == nil || o.MembershipFilter == nil {
		var ret string
		return ret
	}
	return *o.MembershipFilter
}

// GetMembershipFilterOk returns a tuple with the MembershipFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetMembershipFilterOk() (*string, bool) {
	if o == nil || o.MembershipFilter == nil {
		return nil, false
	}
	return o.MembershipFilter, true
}

// HasMembershipFilter returns a boolean if a field has been set.
func (o *LdapProvider) HasMembershipFilter() bool {
	if o != nil && o.MembershipFilter != nil {
		return true
	}

	return false
}

// SetMembershipFilter gets a reference to the given string and assigns it to the MembershipFilter field.
func (o *LdapProvider) SetMembershipFilter(v string) {
	o.MembershipFilter = &v
}

// GetMembershipBaseDn returns the MembershipBaseDn field value if set, zero value otherwise.
func (o *LdapProvider) GetMembershipBaseDn() string {
	if o == nil || o.MembershipBaseDn == nil {
		var ret string
		return ret
	}
	return *o.MembershipBaseDn
}

// GetMembershipBaseDnOk returns a tuple with the MembershipBaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetMembershipBaseDnOk() (*string, bool) {
	if o == nil || o.MembershipBaseDn == nil {
		return nil, false
	}
	return o.MembershipBaseDn, true
}

// HasMembershipBaseDn returns a boolean if a field has been set.
func (o *LdapProvider) HasMembershipBaseDn() bool {
	if o != nil && o.MembershipBaseDn != nil {
		return true
	}

	return false
}

// SetMembershipBaseDn gets a reference to the given string and assigns it to the MembershipBaseDn field.
func (o *LdapProvider) SetMembershipBaseDn(v string) {
	o.MembershipBaseDn = &v
}

// GetPasswordWarning returns the PasswordWarning field value if set, zero value otherwise.
func (o *LdapProvider) GetPasswordWarning() LdapProviderAllOfPasswordWarning {
	if o == nil || o.PasswordWarning == nil {
		var ret LdapProviderAllOfPasswordWarning
		return ret
	}
	return *o.PasswordWarning
}

// GetPasswordWarningOk returns a tuple with the PasswordWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProvider) GetPasswordWarningOk() (*LdapProviderAllOfPasswordWarning, bool) {
	if o == nil || o.PasswordWarning == nil {
		return nil, false
	}
	return o.PasswordWarning, true
}

// HasPasswordWarning returns a boolean if a field has been set.
func (o *LdapProvider) HasPasswordWarning() bool {
	if o != nil && o.PasswordWarning != nil {
		return true
	}

	return false
}

// SetPasswordWarning gets a reference to the given LdapProviderAllOfPasswordWarning and assigns it to the PasswordWarning field.
func (o *LdapProvider) SetPasswordWarning(v LdapProviderAllOfPasswordWarning) {
	o.PasswordWarning = &v
}

func (o LdapProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedIdentityProvider, errIdentityProvider := json.Marshal(o.IdentityProvider)
	if errIdentityProvider != nil {
		return []byte{}, errIdentityProvider
	}
	errIdentityProvider = json.Unmarshal([]byte(serializedIdentityProvider), &toSerialize)
	if errIdentityProvider != nil {
		return []byte{}, errIdentityProvider
	}
	if true {
		toSerialize["hostnames"] = o.Hostnames
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.SslEnabled != nil {
		toSerialize["sslEnabled"] = o.SslEnabled
	}
	if true {
		toSerialize["adminDistinguishedName"] = o.AdminDistinguishedName
	}
	if o.AdminPassword != nil {
		toSerialize["adminPassword"] = o.AdminPassword
	}
	if o.BaseDn != nil {
		toSerialize["baseDn"] = o.BaseDn
	}
	if o.ObjectClass != nil {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if o.UsernameAttribute != nil {
		toSerialize["usernameAttribute"] = o.UsernameAttribute
	}
	if o.MembershipFilter != nil {
		toSerialize["membershipFilter"] = o.MembershipFilter
	}
	if o.MembershipBaseDn != nil {
		toSerialize["membershipBaseDn"] = o.MembershipBaseDn
	}
	if o.PasswordWarning != nil {
		toSerialize["passwordWarning"] = o.PasswordWarning
	}
	return json.Marshal(toSerialize)
}

type NullableLdapProvider struct {
	value *LdapProvider
	isSet bool
}

func (v NullableLdapProvider) Get() *LdapProvider {
	return v.value
}

func (v *NullableLdapProvider) Set(val *LdapProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProvider(val *LdapProvider) *NullableLdapProvider {
	return &NullableLdapProvider{value: val, isSet: true}
}

func (v NullableLdapProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}