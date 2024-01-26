package repositories

import (
	"errors"
	"fmt"
	"saveDatabase/project/configs/dataBase"
	"saveDatabase/project/models"

	"gorm.io/gorm"
)

func ClientFindAll() ([]*models.Client, error) {
	fmt.Println("Chegou na Repository")
	db := dataBase.Conn()

	var clients []*models.Client
	result := db.Table("clients").Preload("Groups").Find(&clients)
	if result.Error != nil {
		return nil, fmt.Errorf("Clientes não encontrados")
	}

	dataBase.CloseConnection(db)

	fmt.Println(clients)
	return clients, nil
}

func ClientFindByID(id string) (*models.Client, error) {
	db := dataBase.Conn()

	var client models.Client
	result := db.Preload("Groups").First(&client, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Cliente não encontrado")
		}
		return nil, result.Error
	}

	dataBase.CloseConnection(db)

	return &client, nil
}

func ClientFindByIDAndGroup(id, group string) ([]models.Files, error) {
	db := dataBase.Conn()

	var files []models.Files
	result := db.Table("files").Where("client_id = ? AND phone LIKE ?", id, "%"+group+"%").Find(&files)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Arquivo não encontrado")
		}
		return nil, result.Error
	}

	dataBase.CloseConnection(db)

	return files, nil
}

func ClientFindByIDAndGroupWithDate(id, group, date string) ([]models.Files, error) {
	db := dataBase.Conn()

	var files []models.Files
	result := db.Table("files").Where("client_id = ? AND phone LIKE ? AND DATE(data_send_message) = ?", id, "%"+group+"%", date).Find(&files)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Arquivo não encontrado")
		}
		return nil, result.Error
	}

	dataBase.CloseConnection(db)

	return files, nil
}

func SaveClient(client *models.Client) (*models.Client, error) {
	db := dataBase.Conn()

	result := db.Table("clients").Create(client)
	if result.Error != nil {
		return nil, fmt.Errorf("Repository: Erro ao salvar um cliente")
	}

	dataBase.CloseConnection(db)

	return client, nil
}

func UpdateClient(id string, client *models.Client) (*models.Client, error) {
	db := dataBase.Conn()

	result := db.Model(&models.Client{}).Where("id = ?", id).Updates(client)
	if result.Error != nil {
		return nil, fmt.Errorf("Cliente não encontrado")
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Cliente não encontrado")
	}

	updateclient := &models.Client{}
	if err := db.First(updateclient, id).Error; err != nil {
		return nil, fmt.Errorf("erro ao atualizar o cliente")
	}
	dataBase.CloseConnection(db)

	return updateclient, nil
}
