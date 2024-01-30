package services

import (
	"errors"
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
		return nil, fmt.Errorf("Erro ao fazer parsing da data: " + date + "  Formato Invalido.")
	}

	formattedDate := newDate.Format("2006-01-02")
	if formattedDate == "" {
		return nil, errors.New("erro ao formatar a data para o formato 'yyyy-MM-dd'")
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
		return nil, errors.New("erro ao formatar a data para o formato 'yyyy-MM-dd'")
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

	brasilia, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o fuso horário:", err)
		return nil, err
	}

	file.DataSendMessage = time.Now().In(brasilia)

	file.ClientID, err = FindClientIDByGroup(file.Phone)
	if err != nil {
		return nil, err
	}
	if file.ClientID == 0 {
		return nil, errors.New("grupo enviado não pertence a nenhum cliente - Client_id = 0")
	}

	persistedFiles, err := repositories.SaveFiles(file)
	if err != nil {
		return nil, err
	}

	fmt.Println("Depois da Repository")
	fmt.Println(persistedFiles)

	return persistedFiles, nil
}
