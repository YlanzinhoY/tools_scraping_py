package services

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CommandRunner(name, commandPip string) {

	var cmd *exec.Cmd

	if runtime.GOOS == "darwin" {
		cmd = exec.Command("zsh", "-c", commandPip)
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("bash", "-c", commandPip)
	} else {
		fmt.Println("Sistema operacional não suportado")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Printf("Erro ao executar o comando: %s\n", err)
		return
	}

	fmt.Printf("%s instalado com sucesso!\n", name)
}
