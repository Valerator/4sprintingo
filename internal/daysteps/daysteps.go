package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("len < 2")
	}
	counter, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	if counter <= 0 {
		return 0, 0, fmt.Errorf("count of steps must be > 0")
	}
	t, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}
	if t <= 0 {
		return 0, 0, fmt.Errorf("count of time must be > 0")
	}
	return counter, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	counter, t, err := parsePackage(data)
	if err != nil || counter == 0 {
		return ""
	}
	if counter <= 0 {
		return ""
	}
	distationMeters := float64(counter) * stepLength
	distationKilometers := distationMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(counter, weight, height, t)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", counter, distationKilometers, calories)
}
