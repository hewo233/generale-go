package controller

import (
	"fmt"
	"hash/fnv"
)

func GenerateColor(userID string) string {
	hasher := fnv.New32a()
	hasher.Write([]byte(userID))
	hash := hasher.Sum32()
	red := hash & 0xFF
	green := (hash >> 8) & 0xFF
	blue := (hash >> 16) & 0xFF
	return fmt.Sprintf("#%02x%02x%02x", red, green, blue)
}
