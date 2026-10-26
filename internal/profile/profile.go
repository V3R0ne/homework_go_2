package profile

import (
	"fmt"
	"strings"
)

// Блок profile — объединяющая задача.
// Здесь нужно закрепить:
// string normalization, int, bool, if,
// форматирование строки и простую бизнес-логику.

// BuildUserCard собирает карточку пользователя.
//
// TODO: верните строку "name=<name> age=<age> group=<group> status=<status>". Имя очистите по краям, приведите к нижнему регистру и сделайте первый символ заглавным; пустое имя замените на "Unknown". Возраст от 18 — adult, иначе minor; active=true — active, иначе inactive.
func BuildUserCard(name string, age int, active bool) string {
	clean := strings.ToLower(strings.TrimSpace(name))
	if clean == "" {
		clean = "Unknown"
	} else {
		runes := []rune(clean)
		clean = strings.ToUpper(string(runes[0])) + string(runes[1:])
	}

	group := "minor"
	if age >= 18 {
		group = "adult"
	}

	status := "inactive"
	if active {
		status = "active"
	}

	return fmt.Sprintf("name=%s age=%d group=%s status=%s", clean, age, group, status)
}
