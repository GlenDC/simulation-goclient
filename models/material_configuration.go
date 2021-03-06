package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MaterialConfiguration material configuration
// swagger:model MaterialConfiguration
type MaterialConfiguration struct {

	// anisotropic strain coefficient parallel
	// Required: true
	AnisotropicStrainCoefficientParallel *float64 `json:"anisotropicStrainCoefficientParallel"`

	// anisotropic strain coefficient perpendicular
	// Required: true
	AnisotropicStrainCoefficientPerpendicular *float64 `json:"anisotropicStrainCoefficientPerpendicular"`

	// anisotropic strain coefficient z
	// Required: true
	AnisotropicStrainCoefficientZ *float64 `json:"anisotropicStrainCoefficientZ"`

	// assumed strain
	// Required: true
	AssumedStrain *float64 `json:"assumedStrain"`

	// atomic weight
	// Required: true
	AtomicWeight *float64 `json:"atomicWeight"`

	// created time stamp, set server-side, read only field
	// Required: true
	Created *strfmt.DateTime `json:"created"`

	// creating user, set server-side, read only field
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// elastic modulus
	// Required: true
	ElasticModulus *float64 `json:"elasticModulus"`

	// elastic modulus of base
	// Required: true
	ElasticModulusOfBase *float64 `json:"elasticModulusOfBase"`

	// energy absorbing rate by powder
	// Required: true
	EnergyAbsorbingRateByPowder *float64 `json:"energyAbsorbingRateByPowder"`

	// energy absorbing rate by solid
	// Required: true
	EnergyAbsorbingRateBySolid *float64 `json:"energyAbsorbingRateBySolid"`

	// fusion latent heat
	// Required: true
	FusionLatentHeat *float64 `json:"fusionLatentHeat"`

	// material configuration identifier
	ID int32 `json:"id,omitempty"`

	// liquidus temperature
	// Required: true
	LiquidusTemperature *float64 `json:"liquidusTemperature"`

	// Location where the lookup file is stored in S3.  Set server side.  Only used internally within 3DSIM.
	LookupFileLocation string `json:"lookupFileLocation,omitempty"`

	// material identifier for this material configuration
	// Required: true
	MaterialID *int32 `json:"materialId"`

	// material strain sensitivity
	// Required: true
	MaterialStrainSensitivity *float64 `json:"materialStrainSensitivity"`

	// material yield strength
	// Required: true
	MaterialYieldStrength *float64 `json:"materialYieldStrength"`

	// mean free path of laser in bulk
	// Required: true
	MeanFreePathOfLaserInBulk *float64 `json:"meanFreePathOfLaserInBulk"`

	// mean free path of laser in powder
	// Required: true
	MeanFreePathOfLaserInPowder *float64 `json:"meanFreePathOfLaserInPowder"`

	// poisson ratio
	// Required: true
	PoissonRatio *float64 `json:"poissonRatio"`

	// powder to solid density ratio
	// Required: true
	PowderToSolidDensityRatio *float64 `json:"powderToSolidDensityRatio"`

	// powder to solid specific heat ratio
	// Required: true
	PowderToSolidSpecificHeatRatio *float64 `json:"powderToSolidSpecificHeatRatio"`

	// powder to solid thermal conductivity ratio
	// Required: true
	PowderToSolidThermalConductivityRatio *float64 `json:"powderToSolidThermalConductivityRatio"`

	// purging gas convection coefficient
	// Required: true
	PurgingGasConvectionCoefficient *float64 `json:"purgingGasConvectionCoefficient"`

	// solid density at room temperature
	// Required: true
	SolidDensityAtRoomTemperature *float64 `json:"solidDensityAtRoomTemperature"`

	// solid specific heat at room temperature
	// Required: true
	SolidSpecificHeatAtRoomTemperature *float64 `json:"solidSpecificHeatAtRoomTemperature"`

	// solid state transition temperature
	// Required: true
	SolidStateTransitionTemperature *float64 `json:"solidStateTransitionTemperature"`

	// solid thermal conductivity at room temperature
	// Required: true
	SolidThermalConductivityAtRoomTemperature *float64 `json:"solidThermalConductivityAtRoomTemperature"`

	// solidus temperature
	// Required: true
	SolidusTemperature *float64 `json:"solidusTemperature"`

	// strain plasticity relaxation factor
	// Required: true
	StrainPlasticityRelaxationFactor *float64 `json:"strainPlasticityRelaxationFactor"`

	// support yield strength ratio
	// Required: true
	SupportYieldStrengthRatio *float64 `json:"supportYieldStrengthRatio"`

	// surface tension
	// Required: true
	SurfaceTension *float64 `json:"surfaceTension"`

	// thermal expansion coefficient
	// Required: true
	ThermalExpansionCoefficient *float64 `json:"thermalExpansionCoefficient"`

	// vaporization latent heat
	// Required: true
	VaporizationLatentHeat *float64 `json:"vaporizationLatentHeat"`

	// vaporization start temperature
	// Required: true
	VaporizationStartTemperature *float64 `json:"vaporizationStartTemperature"`

	// vaporization temperature
	// Required: true
	VaporizationTemperature *float64 `json:"vaporizationTemperature"`

	// version label for the material configuration
	// Required: true
	// Max Length: 16
	Version *string `json:"version"`
}

// Validate validates this material configuration
func (m *MaterialConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnisotropicStrainCoefficientParallel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAnisotropicStrainCoefficientPerpendicular(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAnisotropicStrainCoefficientZ(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAssumedStrain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAtomicWeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateElasticModulus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateElasticModulusOfBase(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnergyAbsorbingRateByPowder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnergyAbsorbingRateBySolid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFusionLatentHeat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLiquidusTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaterialID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaterialStrainSensitivity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaterialYieldStrength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeanFreePathOfLaserInBulk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeanFreePathOfLaserInPowder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePoissonRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowderToSolidDensityRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowderToSolidSpecificHeatRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowderToSolidThermalConductivityRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePurgingGasConvectionCoefficient(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolidDensityAtRoomTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolidSpecificHeatAtRoomTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolidStateTransitionTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolidThermalConductivityAtRoomTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolidusTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStrainPlasticityRelaxationFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportYieldStrengthRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSurfaceTension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThermalExpansionCoefficient(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVaporizationLatentHeat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVaporizationStartTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVaporizationTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaterialConfiguration) validateAnisotropicStrainCoefficientParallel(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientParallel", "body", m.AnisotropicStrainCoefficientParallel); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateAnisotropicStrainCoefficientPerpendicular(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientPerpendicular", "body", m.AnisotropicStrainCoefficientPerpendicular); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateAnisotropicStrainCoefficientZ(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientZ", "body", m.AnisotropicStrainCoefficientZ); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateAssumedStrain(formats strfmt.Registry) error {

	if err := validate.Required("assumedStrain", "body", m.AssumedStrain); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateAtomicWeight(formats strfmt.Registry) error {

	if err := validate.Required("atomicWeight", "body", m.AtomicWeight); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateElasticModulus(formats strfmt.Registry) error {

	if err := validate.Required("elasticModulus", "body", m.ElasticModulus); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateElasticModulusOfBase(formats strfmt.Registry) error {

	if err := validate.Required("elasticModulusOfBase", "body", m.ElasticModulusOfBase); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateEnergyAbsorbingRateByPowder(formats strfmt.Registry) error {

	if err := validate.Required("energyAbsorbingRateByPowder", "body", m.EnergyAbsorbingRateByPowder); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateEnergyAbsorbingRateBySolid(formats strfmt.Registry) error {

	if err := validate.Required("energyAbsorbingRateBySolid", "body", m.EnergyAbsorbingRateBySolid); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateFusionLatentHeat(formats strfmt.Registry) error {

	if err := validate.Required("fusionLatentHeat", "body", m.FusionLatentHeat); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateLiquidusTemperature(formats strfmt.Registry) error {

	if err := validate.Required("liquidusTemperature", "body", m.LiquidusTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateMaterialID(formats strfmt.Registry) error {

	if err := validate.Required("materialId", "body", m.MaterialID); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateMaterialStrainSensitivity(formats strfmt.Registry) error {

	if err := validate.Required("materialStrainSensitivity", "body", m.MaterialStrainSensitivity); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateMaterialYieldStrength(formats strfmt.Registry) error {

	if err := validate.Required("materialYieldStrength", "body", m.MaterialYieldStrength); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateMeanFreePathOfLaserInBulk(formats strfmt.Registry) error {

	if err := validate.Required("meanFreePathOfLaserInBulk", "body", m.MeanFreePathOfLaserInBulk); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateMeanFreePathOfLaserInPowder(formats strfmt.Registry) error {

	if err := validate.Required("meanFreePathOfLaserInPowder", "body", m.MeanFreePathOfLaserInPowder); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validatePoissonRatio(formats strfmt.Registry) error {

	if err := validate.Required("poissonRatio", "body", m.PoissonRatio); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validatePowderToSolidDensityRatio(formats strfmt.Registry) error {

	if err := validate.Required("powderToSolidDensityRatio", "body", m.PowderToSolidDensityRatio); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validatePowderToSolidSpecificHeatRatio(formats strfmt.Registry) error {

	if err := validate.Required("powderToSolidSpecificHeatRatio", "body", m.PowderToSolidSpecificHeatRatio); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validatePowderToSolidThermalConductivityRatio(formats strfmt.Registry) error {

	if err := validate.Required("powderToSolidThermalConductivityRatio", "body", m.PowderToSolidThermalConductivityRatio); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validatePurgingGasConvectionCoefficient(formats strfmt.Registry) error {

	if err := validate.Required("purgingGasConvectionCoefficient", "body", m.PurgingGasConvectionCoefficient); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateSolidDensityAtRoomTemperature(formats strfmt.Registry) error {

	if err := validate.Required("solidDensityAtRoomTemperature", "body", m.SolidDensityAtRoomTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateSolidSpecificHeatAtRoomTemperature(formats strfmt.Registry) error {

	if err := validate.Required("solidSpecificHeatAtRoomTemperature", "body", m.SolidSpecificHeatAtRoomTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateSolidStateTransitionTemperature(formats strfmt.Registry) error {

	if err := validate.Required("solidStateTransitionTemperature", "body", m.SolidStateTransitionTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateSolidThermalConductivityAtRoomTemperature(formats strfmt.Registry) error {

	if err := validate.Required("solidThermalConductivityAtRoomTemperature", "body", m.SolidThermalConductivityAtRoomTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateSolidusTemperature(formats strfmt.Registry) error {

	if err := validate.Required("solidusTemperature", "body", m.SolidusTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateStrainPlasticityRelaxationFactor(formats strfmt.Registry) error {

	if err := validate.Required("strainPlasticityRelaxationFactor", "body", m.StrainPlasticityRelaxationFactor); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateSupportYieldStrengthRatio(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrengthRatio", "body", m.SupportYieldStrengthRatio); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateSurfaceTension(formats strfmt.Registry) error {

	if err := validate.Required("surfaceTension", "body", m.SurfaceTension); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateThermalExpansionCoefficient(formats strfmt.Registry) error {

	if err := validate.Required("thermalExpansionCoefficient", "body", m.ThermalExpansionCoefficient); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateVaporizationLatentHeat(formats strfmt.Registry) error {

	if err := validate.Required("vaporizationLatentHeat", "body", m.VaporizationLatentHeat); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateVaporizationStartTemperature(formats strfmt.Registry) error {

	if err := validate.Required("vaporizationStartTemperature", "body", m.VaporizationStartTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateVaporizationTemperature(formats strfmt.Registry) error {

	if err := validate.Required("vaporizationTemperature", "body", m.VaporizationTemperature); err != nil {
		return err
	}

	return nil
}

func (m *MaterialConfiguration) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MaxLength("version", "body", string(*m.Version), 16); err != nil {
		return err
	}

	return nil
}
