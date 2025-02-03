package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func PreprocessToJSON(input string) string {
	// Elimina los corchetes iniciales y finales
	trimmed := strings.Trim(input, "{}")

	// Separa por comas, pero mantiene arreglos y objetos anidados
	pairs := splitKeepingArraysAndObjects(trimmed)

	// Regex para identificar clave:valor y aplicar tratamiento adecuado
	re := regexp.MustCompile(`(\w+):(.*)`)
	var jsonPairs []string
	for _, pair := range pairs {
		matches := re.FindStringSubmatch(pair)
		if len(matches) == 3 {
			key := matches[1]
			value := strings.TrimSpace(matches[2])

			// Procesar el valor según su tipo
			if isBoolean(value) || isArray(value) {
				// Dejar valores numéricos, booleanos y arreglos como están
				jsonPairs = append(jsonPairs, fmt.Sprintf(`"%s":%s`, key, value))
			} else {
				// Envolver cadenas con comillas
				jsonPairs = append(jsonPairs, fmt.Sprintf(`"%s":"%s"`, key, value))
			}
		}
	}

	// Une las partes y encapsula en llaves
	return "{" + strings.Join(jsonPairs, ",") + "}"
}

// Divide la cadena pero mantiene arreglos y objetos intactos
func splitKeepingArraysAndObjects(input string) []string {
	var parts []string
	var buffer strings.Builder
	openBrackets := 0

	for _, char := range input {
		if char == ',' && openBrackets == 0 {
			parts = append(parts, buffer.String())
			buffer.Reset()
		} else {
			if char == '[' || char == '{' {
				openBrackets++
			} else if char == ']' || char == '}' {
				openBrackets--
			}
			buffer.WriteRune(char)
		}
	}

	// Agrega la última parte
	if buffer.Len() > 0 {
		parts = append(parts, buffer.String())
	}

	return parts
}

// Verifica si el valor es un número
func isNumeric(value string) bool {
	re := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	return re.MatchString(value)
}

// Verifica si el valor es un booleano
func isBoolean(value string) bool {
	return value == "true" || value == "false"
}

// Verifica si el valor es un arreglo
func isArray(value string) bool {
	return strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]")
}