package repositories

import (
	"errors"
	"fmt"
	"saveDatabase/project/configs/dataBase"
	"saveDatabase/project/models"

	"gorm.io/gorm"
)

func FilesFindAll() ([]*models.Files, error) {
	fmt.Println("Chegou na Repository")
	db := dataBase.Conn()

	var files []*models.Files
	result := db.Table("files").Find(&files)
	if result.Error != nil {
		return nil, fmt.Errorf("Arquivos não encontrados")
	}

	dataBase.CloseConnection(db)

	fmt.Println(files)
	return files, nil
}

func FileFindByPhone(phone string) ([]models.Files, error) {
	db := dataBase.Conn()

	var files []models.Files
	result := db.Table("files").Where("phone LIKE ?", "%"+phone+"%").Find(&files)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Arquivo não encontrado")
		}
		return nil, result.Error
	}

	dataBase.CloseConnection(db)

	return files, nil
}

func FileFindByDate(date string) ([]models.Files, error) {
	db := dataBase.Conn()

	var files []models.Files
	result := db.Table("files").Where("DATE(data_send_message) = ?", date).Find(&files)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Arquivo não encontrado")
		}
		return nil, result.Error
	}

	dataBase.CloseConnection(db)

	return files, nil
}

func FileFindByGroupAndDate(group, date string) ([]models.Files, error) {
	db := dataBase.Conn()

	var files []models.Files
	result := db.Table("files").Where("phone LIKE ? AND DATE(data_send_message) = ?", "%"+group+"%", date).
		Find(&files)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Arquivo não encontrado")
		}
		return nil, result.Error
	}

	dataBase.CloseConnection(db)

	return files, nil
}

func SaveFiles(file *models.Files) (*models.Files, error) {
	db := dataBase.Conn()

	result := db.Table("files").Create(file)
	if result.Error != nil {
		return nil, fmt.Errorf("Repository: Erro ao salvar um arquivo")
	}

	dataBase.CloseConnection(db)

	return file, nil
}
