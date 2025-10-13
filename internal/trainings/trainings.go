package trainings
import (
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"time"
	"strings"
	"fmt"
	"strconv"
	sp "github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"log"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("error parsing value: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("error parsing value")
	}

	t.Steps = steps
	t.TrainingType = slice[1]

	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("error parsing duration")
	}

	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := sp.Distance(t.Steps, t.Height)
	averageSpeed := sp.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error

	switch(t.TrainingType){
		case "Ходьба":
			calories, err = sp.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
			if err != nil {
            	log.Println(err)
            	return "", err
        	}
		case "Бег":
			calories, err = sp.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
			if err != nil {
            	log.Println(err)
            	return "", err
        	}
		default:
			return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	result := fmt.Sprintf(
        "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
        t.TrainingType,
        t.Duration.Hours(),
        dist,
        averageSpeed,
        calories,
    )
	return result, nil
}
