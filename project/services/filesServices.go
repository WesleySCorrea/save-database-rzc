package services

import (
	"fmt"
	"saveDatabase/project/models"
	"saveDatabase/project/repositories"
	"time"
)

func FilesFindAll() ([]*models.Files, error) {

	fmt.Println("Chegou na Service")
	files, err := repositories.FilesFindAll()
	if err != nil {
		return nil, err
	}
	fmt.Println("Saiu da Repository")

	return files, nil
}

func FileFindByPhone(phone string) ([]models.Files, error) {

	files, err := repositories.FileFindByPhone(phone)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func FileFindByDate(date string) ([]models.Files, error) {

	newDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer parsing da data: " + date + " \n Formato Invalido.")
	}

	formattedDate := newDate.Format("2006-01-02")
	if formattedDate == "" {
		return nil, fmt.Errorf("Erro ao formatar a data para '2006-01-02'")
	}

	files, err := repositories.FileFindByDate(formattedDate)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func FileFindByGroupAndDate(group, date string) ([]models.Files, error) {

	newDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer parsing da data: " + date + " \n Formato Invalido.")
	}

	formattedDate := newDate.Format("2006-01-02")
	if formattedDate == "" {
		return nil, fmt.Errorf("Erro ao formatar a data para '2006-01-02'")
	}

	files, err := repositories.FileFindByGroupAndDate(group, formattedDate)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func SaveFiles(file *models.Files) (*models.Files, error) {

	fmt.Println("Chegou Service")
	fmt.Println(file)
	file.DataSendMessage = time.Now()

	persistedFiles, err := repositories.SaveFiles(file)
	if err != nil {
		return nil, err
	}

	fmt.Println("Depois da Repository")
	fmt.Println(persistedFiles)

	return persistedFiles, nil
}
