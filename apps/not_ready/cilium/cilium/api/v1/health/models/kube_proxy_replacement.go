// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubeProxyReplacement Status of kube-proxy replacement
// swagger:model kubeProxyReplacement
type KubeProxyReplacement struct {

	// features
	Features *KubeProxyReplacementFeatures `json:"features,omitempty"`

	// mode
	// Enum: [Disabled Strict Probe Partial]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this kube proxy replacement
func (m *KubeProxyReplacement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubeProxyReplacement) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

var kubeProxyReplacementTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Disabled","Strict","Probe","Partial"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubeProxyReplacementTypeModePropEnum = append(kubeProxyReplacementTypeModePropEnum, v)
	}
}

const (

	// KubeProxyReplacementModeDisabled captures enum value "Disabled"
	KubeProxyReplacementModeDisabled string = "Disabled"

	// KubeProxyReplacementModeStrict captures enum value "Strict"
	KubeProxyReplacementModeStrict string = "Strict"

	// KubeProxyReplacementModeProbe captures enum value "Probe"
	KubeProxyReplacementModeProbe string = "Probe"

	// KubeProxyReplacementModePartial captures enum value "Partial"
	KubeProxyReplacementModePartial string = "Partial"
)

// prop value enum
func (m *KubeProxyReplacement) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, kubeProxyReplacementTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KubeProxyReplacement) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacement) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeatures kube proxy replacement features
// swagger:model KubeProxyReplacementFeatures
type KubeProxyReplacementFeatures struct {

	// external i ps
	ExternalIPs *KubeProxyReplacementFeaturesExternalIPs `json:"externalIPs,omitempty"`

	// host reachable services
	HostReachableServices *KubeProxyReplacementFeaturesHostReachableServices `json:"hostReachableServices,omitempty"`

	// node port
	NodePort *KubeProxyReplacementFeaturesNodePort `json:"nodePort,omitempty"`
}

// Validate validates this kube proxy replacement features
func (m *KubeProxyReplacementFeatures) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalIPs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostReachableServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubeProxyReplacementFeatures) validateExternalIPs(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalIPs) { // not required
		return nil
	}

	if m.ExternalIPs != nil {
		if err := m.ExternalIPs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "externalIPs")
			}
			return err
		}
	}

	return nil
}

func (m *KubeProxyReplacementFeatures) validateHostReachableServices(formats strfmt.Registry) error {

	if swag.IsZero(m.HostReachableServices) { // not required
		return nil
	}

	if m.HostReachableServices != nil {
		if err := m.HostReachableServices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "hostReachableServices")
			}
			return err
		}
	}

	return nil
}

func (m *KubeProxyReplacementFeatures) validateNodePort(formats strfmt.Registry) error {

	if swag.IsZero(m.NodePort) { // not required
		return nil
	}

	if m.NodePort != nil {
		if err := m.NodePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "nodePort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeatures) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesExternalIPs kube proxy replacement features external i ps
// swagger:model KubeProxyReplacementFeaturesExternalIPs
type KubeProxyReplacementFeaturesExternalIPs struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this kube proxy replacement features external i ps
func (m *KubeProxyReplacementFeaturesExternalIPs) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesExternalIPs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesExternalIPs) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesExternalIPs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesHostReachableServices kube proxy replacement features host reachable services
// swagger:model KubeProxyReplacementFeaturesHostReachableServices
type KubeProxyReplacementFeaturesHostReachableServices struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// protocols
	Protocols []string `json:"protocols"`
}

// Validate validates this kube proxy replacement features host reachable services
func (m *KubeProxyReplacementFeaturesHostReachableServices) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesHostReachableServices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesHostReachableServices) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesHostReachableServices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesNodePort kube proxy replacement features node port
// swagger:model KubeProxyReplacementFeaturesNodePort
type KubeProxyReplacementFeaturesNodePort struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// mode
	// Enum: [SNAT DSR HYBRID]
	Mode string `json:"mode,omitempty"`

	// port max
	PortMax int64 `json:"portMax,omitempty"`

	// port min
	PortMin int64 `json:"portMin,omitempty"`
}

// Validate validates this kube proxy replacement features node port
func (m *KubeProxyReplacementFeaturesNodePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kubeProxyReplacementFeaturesNodePortTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SNAT","DSR","HYBRID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubeProxyReplacementFeaturesNodePortTypeModePropEnum = append(kubeProxyReplacementFeaturesNodePortTypeModePropEnum, v)
	}
}

const (

	// KubeProxyReplacementFeaturesNodePortModeSNAT captures enum value "SNAT"
	KubeProxyReplacementFeaturesNodePortModeSNAT string = "SNAT"

	// KubeProxyReplacementFeaturesNodePortModeDSR captures enum value "DSR"
	KubeProxyReplacementFeaturesNodePortModeDSR string = "DSR"

	// KubeProxyReplacementFeaturesNodePortModeHYBRID captures enum value "HYBRID"
	KubeProxyReplacementFeaturesNodePortModeHYBRID string = "HYBRID"
)

// prop value enum
func (m *KubeProxyReplacementFeaturesNodePort) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, kubeProxyReplacementFeaturesNodePortTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KubeProxyReplacementFeaturesNodePort) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("features"+"."+"nodePort"+"."+"mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesNodePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesNodePort) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesNodePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
