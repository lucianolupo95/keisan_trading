package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CallKaton ejecuta el módulo Katon (Python) desde Fuuton (Go)
func CallKaton(pythonPath string) (string, error) {
	cmd := exec.Command("python", pythonPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// ParseKatonResponse extrae la respuesta "Katon OK" del output
func ParseKatonResponse(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Ping response:") {
			return strings.TrimSpace(strings.ReplaceAll(line, "Ping response:", ""))
		}
	}
	return "NO RESPONSE"
}

func main() {
	fmt.Println("Fuuton activo")
	fmt.Println("Calling Katon (Python)...")
	fmt.Println()

	// Ruta a ping.py
	katonPath := "../katon-python/ping.py"

	// Verificar que el archivo existe
	if _, err := os.Stat(katonPath); os.IsNotExist(err) {
		fmt.Printf("Error: Archivo %s no encontrado\n", katonPath)
		return
	}

	// Llamar a Katon
	output, err := CallKaton(katonPath)
	if err != nil {
		fmt.Printf("Error ejecutando Katon: %v\n", err)
		return
	}

	// Parsear respuesta
	response := ParseKatonResponse(output)
	fmt.Printf("Katon response: %s\n", response)
	fmt.Println()
	fmt.Println("Katon output (full):")
	fmt.Println(output)
}
