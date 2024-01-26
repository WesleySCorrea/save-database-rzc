package services

import (
	"fmt"
	"saveDatabase/project/models"
	"saveDatabase/project/repositories"
	"time"
)

func ClientFindAll() ([]*models.Client, error) {

	fmt.Println("Chegou na Service")
	clients, err := repositories.ClientFindAll()
	if err != nil {
		return nil, err
	}
	fmt.Println("Saiu da Repository")

	return clients, nil
}

func ClientFindByID(id string) (*models.Client, error) {

	client, err := repositories.ClientFindByID(id)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ClientFindByIDAndGroup(id, group string) ([]models.Files, error) {

	files, err := repositories.ClientFindByIDAndGroup(id, group)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func ClientFindByIDAndGroupWithDate(id, group, date string) ([]models.Files, error) {

	newDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer parsing da data: " + date + ". Formato Invalido.")
	}

	formattedDate := newDate.Format("2006-01-02")
	if formattedDate == "" {
		return nil, fmt.Errorf("Erro ao formatar a data para 'yyyy-MM-dd'")
	}

	files, err := repositories.ClientFindByIDAndGroupWithDate(id, group, formattedDate)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func SaveClient(client *models.Client) (*models.Client, error) {

	fmt.Println("Chegou Service")
	fmt.Println(client)

	persistedClient, err := repositories.SaveClient(client)
	if err != nil {
		return nil, err
	}

	fmt.Println("Depois da Repository")
	fmt.Println(persistedClient)

	return persistedClient, nil
}

func UpdateClient(id string, client *models.Client) (*models.Client, error) {
	_, err := repositories.ClientFindByID(id)
	if err != nil {
		return nil, err
	}

	clientUpdate, err := repositories.UpdateClient(id, client)
	if err != nil {
		return nil, err
	}

	return clientUpdate, nil
}
