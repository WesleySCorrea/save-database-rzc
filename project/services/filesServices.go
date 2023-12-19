package services

import (
	"fmt"
	"saveDatabase/project/models"
	"saveDatabase/project/repositories"
)

func FilesFindAll() ([]*models.Geral, error) {

	fmt.Println("Chegou na Service")
	files, err := repositories.FilesFindAll()
	if err != nil {
		return nil, err
	}
	fmt.Println("Saiu da Repository")

	return files, nil
}

func FileFindByPhone(phone string) ([]models.Geral, error) {

	files, err := repositories.FileFindByPhone(phone)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func SaveFiles(file *models.Geral) (*models.Geral, error) {

	fmt.Println("Chegou Service")
	fmt.Println(file)

	persistedFiles, err := repositories.SaveFiles(file)
	if err != nil {
		return nil, err
	}

	fmt.Println("Depois da Repository")
	fmt.Println(persistedFiles)

	return persistedFiles, nil
}
