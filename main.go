package main

import (
	"fmt"
)

func main() {
	sysInfo, err := c0din_client.GetSystemInfo()
	if err != nil {
		fmt.Println("Erro ao obter informações do sistema:", err)
		return
	}

	fmt.Println("Informações do sistema:")
	fmt.Println("Disco(s):")
	for _, disk := range sysInfo.Disks {
		fmt.Printf("\tNome: %s, Tamanho: %d GB, Espaço livre: %d GB\n", disk.Name, disk.Size/1024/1024/1024, disk.FreeSpace/1024/1024/1024)
	}
	fmt.Println("Sistema operacional:", sysInfo.OS)
	fmt.Println("Modelo da CPU:", sysInfo.CPUModel)
	fmt.Println("Modelo da RAM:", sysInfo.RAMModel)
	fmt.Printf("Tamanho da RAM: %d GB\n", sysInfo.RAMSize/1024/1024/1024)
	fmt.Printf("Temperatura da CPU: %.2f °C\n", sysInfo.CPUTemperature)
}
