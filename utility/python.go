package utility

import (
	"log"
	"os"
	"os/exec"
)

type Python struct{}

func (p *Python) Handle(filepath string) string {

	cmdArgs := append([]string{"./py3/get_files_content.py"}, filepath)
	cmd := exec.Command("python", cmdArgs...)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")

	// 获取命令输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing Python script: %v", err)
	}

	return string(output)
}
