package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	sp "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("invalid data format")
	}

	val, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("error parsing value: %w", err)
	}
	if val <= 0 {
		return fmt.Errorf("steps cannot be zero")
	}

	ds.Steps = val

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("duration cannot be zero")
	}

	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration cannot be zero or negative")
	}
	dist := sp.Distance(ds.Steps, ds.Height)
	calories, err := sp.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories), nil
}
