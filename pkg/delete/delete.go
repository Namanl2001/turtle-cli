package delete

import (
	"os"
)

func Delete(name string) error {
	home := os.Getenv("HOME")
	filepath := home + "/turtle-secrets/" + name

	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}
