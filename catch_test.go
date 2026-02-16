package main

import (
	"testing"
)

func TestCatch(t *testing.T) {
	if Catch(255, 1.0) > Catch(255, 2.0) {
		t.Errorf("Pokeball outperformed Ultraball")
	}
	//tests that Catch returns appropriate catch probabilities for known catch rates
	result := Catch(3, 1.0)
	if result < 0.015 || result > 0.017 {
		t.Errorf("Catch returned %f, expected approximately 0.016", result)
	}
	result = Catch(5, 1.0)
	if result < 0.022 || result > 0.024 {
		t.Errorf("Catch returned %f, expected approximately 0.023", result)
	}
	result = Catch(6, 1.0)
	if result < 0.025 || result > 0.027 {
		t.Errorf("Catch returned %f, expected approximately 0.026", result)
	}
	result = Catch(10, 1.0)
	if result < 0.038 || result > 0.040 {
		t.Errorf("Catch returned %f, expected approximately 0.039", result)
	}
	result = Catch(15, 1.0)
	if result < 0.051 || result > 0.053 {
		t.Errorf("Catch returned %f, expected approximately 0.052", result)
	}
	result = Catch(20, 1.0)
	if result < 0.064 || result > 0.066 {
		t.Errorf("Catch returned %f, expected approximately 0.065", result)
	}
	result = Catch(25, 1.0)
	if result < 0.076 || result > 0.078 {
		t.Errorf("Catch returned %f, expected approximately 0.077", result)
	}
	result = Catch(30, 1.0)
	if result < 0.087 || result > 0.089 {
		t.Errorf("Catch returned %f, expected approximately 0.088", result)
	}
	result = Catch(35, 1.0)
	if result < 0.098 || result > 0.100 {
		t.Errorf("Catch returned %f, expected approximately 0.099", result)
	}
	result = Catch(45, 1.0)
	if result < 0.118 || result > 0.120 {
		t.Errorf("Catch returned %f, expected approximately 0.119", result)
	}
	result = Catch(50, 1.0)
	if result < 0.128 || result > 0.130 {
		t.Errorf("Catch returned %f, expected approximately 0.129", result)
	}
	result = Catch(55, 1.0)
	if result < 0.138 || result > 0.140 {
		t.Errorf("Catch returned %f, expected approximately 0.139", result)
	}
	result = Catch(60, 1.0)
	if result < 0.147 || result > 0.149 {
		t.Errorf("Catch returned %f, expected approximately 0.148", result)
	}
	result = Catch(65, 1.0)
	if result < 0.156 || result > 0.158 {
		t.Errorf("Catch returned %f, expected approximately 0.157", result)
	}
	result = Catch(70, 1.0)
	if result < 0.165 || result > 0.167 {
		t.Errorf("Catch returned %f, expected approximately 0.166", result)
	}
	result = Catch(75, 1.0)
	if result < 0.174 || result > 0.176 {
		t.Errorf("Catch returned %f, expected approximately 0.175", result)
	}
	result = Catch(80, 1.0)
	if result < 0.183 || result > 0.185 {
		t.Errorf("Catch returned %f, expected approximately 0.184", result)
	}
	result = Catch(90, 1.0)
	if result < 0.200 || result > 0.202 {
		t.Errorf("Catch returned %f, expected approximately 0.201", result)
	}
	result = Catch(100, 1.0)
	if result < 0.216 || result > 0.218 {
		t.Errorf("Catch returned %f, expected approximately 0.217", result)
	}
	result = Catch(120, 1.0)
	if result < 0.248 || result > 0.250 {
		t.Errorf("Catch returned %f, expected approximately 0.249", result)
	}
	result = Catch(125, 1.0)
	if result < 0.256 || result > 0.258 {
		t.Errorf("Catch returned %f, expected approximately 0.257", result)
	}
	result = Catch(127, 1.0)
	if result < 0.259 || result > 0.261 {
		t.Errorf("Catch returned %f, expected approximately 0.260", result)
	}
	result = Catch(130, 1.0)
	if result < 0.264 || result > 0.266 {
		t.Errorf("Catch returned %f, expected approximately 0.265", result)
	}
	result = Catch(140, 1.0)
	if result < 0.279 || result > 0.281 {
		t.Errorf("Catch returned %f, expected approximately 0.280", result)
	}
	result = Catch(145, 1.0)
	if result < 0.286 || result > 0.288 {
		t.Errorf("Catch returned %f, expected approximately 0.287", result)
	}
	result = Catch(150, 1.0)
	if result < 0.294 || result > 0.296 {
		t.Errorf("Catch returned %f, expected approximately 0.295", result)
	}
	result = Catch(155, 1.0)
	if result < 0.301 || result > 0.303 {
		t.Errorf("Catch returned %f, expected approximately 0.302", result)
	}
	result = Catch(160, 1.0)
	if result < 0.308 || result > 0.310 {
		t.Errorf("Catch returned %f, expected approximately 0.309", result)
	}
	result = Catch(170, 1.0)
	if result < 0.323 || result > 0.325 {
		t.Errorf("Catch returned %f, expected approximately 0.324", result)
	}
	result = Catch(180, 1.0)
	if result < 0.337 || result > 0.339 {
		t.Errorf("Catch returned %f, expected approximately 0.338", result)
	}
	result = Catch(190, 1.0)
	if result < 0.351 || result > 0.353 {
		t.Errorf("Catch returned %f, expected approximately 0.352", result)
	}
	result = Catch(200, 1.0)
	if result < 0.365 || result > 0.367 {
		t.Errorf("Catch returned %f, expected approximately 0.366", result)
	}
	result = Catch(205, 1.0)
	if result < 0.371 || result > 0.373 {
		t.Errorf("Catch returned %f, expected approximately 0.372", result)
	}
	result = Catch(220, 1.0)
	if result < 0.392 || result > 0.394 {
		t.Errorf("Catch returned %f, expected approximately 0.393", result)
	}
	result = Catch(225, 1.0)
	if result < 0.398 || result > 0.400 {
		t.Errorf("Catch returned %f, expected approximately 0.399", result)
	}
	result = Catch(235, 1.0)
	if result < 0.412 || result > 0.414 {
		t.Errorf("Catch returned %f, expected approximately 0.413", result)
	}
	result = Catch(255, 1.0)
	if result < 0.438 || result > 0.440 {
		t.Errorf("Catch returned %f, expected approximately 0.439", result)
	}
	result = Catch(0, 255)
	if result != 0 {
		t.Errorf("Catch returned %f, expected 0", result)
	}
	result = Catch(255000, 0)
	if result != 0 {
		t.Errorf("Catch returned %f, expected 0", result)
	}
	result = Catch(3, 255.0)
	if result != 1.0 {
		t.Errorf("Catch returned %f, expected 1", result)
	}
}
