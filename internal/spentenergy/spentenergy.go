package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0.0, errors.New("steps, weight and height must be positive")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := ((weight * averageSpeed * minutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0.0, errors.New("steps, weight and height must be positive")
	}
	if duration <= 0 {
		return 0.0, errors.New("duration must be positive")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * averageSpeed * minutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}

	averageSpeed := Distance(steps, height) / duration.Hours()
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	strideLenght := height * stepLengthCoefficient
	return (float64(steps) * strideLenght) / mInKm
}
