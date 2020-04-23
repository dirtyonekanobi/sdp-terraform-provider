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

// ApplianceAllOfIotConnector IoT Connector settings.
type ApplianceAllOfIotConnector struct {
	// Whether the Iot Connector is enabled on this appliance or not.
	Enabled *bool                              `json:"enabled,omitempty"`
	Clients *ApplianceAllOfIotConnectorClients `json:"clients,omitempty"`
}

// NewApplianceAllOfIotConnector instantiates a new ApplianceAllOfIotConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfIotConnector() *ApplianceAllOfIotConnector {
	this := ApplianceAllOfIotConnector{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewApplianceAllOfIotConnectorWithDefaults instantiates a new ApplianceAllOfIotConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfIotConnectorWithDefaults() *ApplianceAllOfIotConnector {
	this := ApplianceAllOfIotConnector{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfIotConnector) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfIotConnector) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfIotConnector) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfIotConnector) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ApplianceAllOfIotConnector) GetClients() ApplianceAllOfIotConnectorClients {
	if o == nil || o.Clients == nil {
		var ret ApplianceAllOfIotConnectorClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfIotConnector) GetClientsOk() (*ApplianceAllOfIotConnectorClients, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ApplianceAllOfIotConnector) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given ApplianceAllOfIotConnectorClients and assigns it to the Clients field.
func (o *ApplianceAllOfIotConnector) SetClients(v ApplianceAllOfIotConnectorClients) {
	o.Clients = &v
}

func (o ApplianceAllOfIotConnector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfIotConnector struct {
	value *ApplianceAllOfIotConnector
	isSet bool
}

func (v NullableApplianceAllOfIotConnector) Get() *ApplianceAllOfIotConnector {
	return v.value
}

func (v *NullableApplianceAllOfIotConnector) Set(val *ApplianceAllOfIotConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfIotConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfIotConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfIotConnector(val *ApplianceAllOfIotConnector) *NullableApplianceAllOfIotConnector {
	return &NullableApplianceAllOfIotConnector{value: val, isSet: true}
}

func (v NullableApplianceAllOfIotConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfIotConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}