package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// ValidateRule
type ValidateRule struct {
	ValueType  string `json:"value_type"`
	RangeType  string `json:"value_range_type"` //between、enum、range
	RangeValue string `json:"value_range"`
}

func (rule *ValidateRule) Validate(data interface{}) (bool, error) {

}

const rangeTypeBetween = "between"
const rangeTypeEnum = "enum"
const rangeTypeExpression = "expression"

func ValidateInt(data interface{}, rangeType string, rangeValue string) (bool, error) {
	value := getInt64(data)
	switch rangeType {
	case rangeTypeBetween:
		rangeInt := []int64{}
		if err := json.Unmarshal([]byte(rangeValue), &rangeInt); err != nil {
			return false, errors.New("range value error: " + err.Error())
		}
		if len(rangeInt) < 2 {
			return false, errors.New("range value length < 2 ")
		}

		if rangeInt[1] > rangeInt[0] {
			return false, errors.New("range value set error")
		}

		if value < rangeInt[0] {
			return false, fmt.Errorf("value %d  < %d", value, rangeInt[0])
		}
		if value > rangeInt[1] {
			return false, fmt.Errorf("value %d  > %d", value, rangeInt[1])
		}
		return true, nil
	case rangeTypeEnum:
		rangeInt := []int64{}
		if err := json.Unmarshal([]byte(rangeValue), &rangeInt); err != nil {
			return false, errors.New("range value error: " + err.Error())
		}
		for _, item := range rangeInt {
			if item == value {
				return true, nil
			}
		}
		return false, fmt.Errorf("value %d  not in any of %s", value, rangeValue)
	}

}

func getInt64(data interface{}) int64 {
	return 0
}

func main() {

	args := getMap()
	fmt.Println(GetFloat64ValFromMap(args, "-AmbientOcclusionIntensity"))
}

func getMap() map[string]interface{} {
	return map[string]interface{}{
		"-AmbientOcclusionIntensity":             "0.26",
		"-AntiAliasingMethod":                    "DLSS",
		"-BackGroundIndex":                       "G1",
		"-CollectIrradiance":                     "",
		"-DenoiserType":                          "intel",
		"-Exposure":                              "0.2",
		"-GlobalIllumination":                    "",
		"-HighResScreenshotDelay":                "8",
		"-Interval":                              "10",
		"-RTRender":                              "0",
		"-RTXGI":                                 "",
		"-RayTracingAOSamplesPerPixel":           "32",
		"-RayTracingGIMaxBounces":                "1",
		"-RayTracingReflectionsMaxBounces":       "2",
		"-RayTracingReflectionsSamplesPerPixel":  "1",
		"-RayTracingReflectionsScreenPercentage": "100",
		"-RayTracingReflectionsShadowType":       "1",
		"-RectLightSamplesPerPixel":              "8",
		"-ResX":                                  "512",
		"-ResY":                                  "512",
		"-ResolutionX":                           "2048",
		"-ResolutionY":                           "2048",
		"-SampleCount":                           "64",
		"-StylizedGI":                            "",
		"-SubfrustumTileSize":                    "2048",
	}
}

func getJSON() string {
	data := map[string]string{
		"-AmbientOcclusionIntensity":             "0.26",
		"-AntiAliasingMethod":                    "DLSS",
		"-BackGroundIndex":                       "G1",
		"-CollectIrradiance":                     "",
		"-DenoiserType":                          "intel",
		"-Exposure":                              "0.2",
		"-GlobalIllumination":                    "",
		"-HighResScreenshotDelay":                "8",
		"-Interval":                              "10",
		"-RTRender":                              "0",
		"-RTXGI":                                 "",
		"-RayTracingAOSamplesPerPixel":           "32",
		"-RayTracingGIMaxBounces":                "1",
		"-RayTracingReflectionsMaxBounces":       "2",
		"-RayTracingReflectionsSamplesPerPixel":  "1",
		"-RayTracingReflectionsScreenPercentage": "100",
		"-RayTracingReflectionsShadowType":       "1",
		"-RectLightSamplesPerPixel":              "8",
		"-ResX":                                  "512",
		"-ResY":                                  "512",
		"-ResolutionX":                           "2048",
		"-ResolutionY":                           "2048",
		"-SampleCount":                           "64",
		"-StylizedGI":                            "",
		"-SubfrustumTileSize":                    "2048",
	}
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func GetFloat64ValFromMap(container map[string]interface{}, key string) float64 {

	val, ok := container[key]
	if !ok || val == nil {
		return 0
	}

	if v, ok := val.(float64); ok {
		return v
	}

	if v, ok := val.(float32); ok {
		return float64(v)
	}

	if v, ok := val.(string); ok {
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			return val
		}
	}

	return 0
}
