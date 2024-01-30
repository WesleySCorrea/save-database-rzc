package repositories

import (
	"errors"
	"fmt"
	"saveDatabase/project/configs/dataBase"
	"saveDatabase/project/models"

	"gorm.io/gorm"
)

func GroupFindByPhone(phone string) (*models.Groups, error) {
	db := dataBase.Conn()

	var groups models.Groups
	result := db.Table("groups").Where("name_group LIKE ?", "%"+phone+"%").Find(&groups)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Grupo não encontrado")
		}
		return nil, result.Error
	}
	if groups.NameGroup == "" {
		return nil, fmt.Errorf("Grupo não encontrado")
	}

	dataBase.CloseConnection(db)

	return &groups, nil
}

func GroupFindByID(id string) (*models.Groups, error) {
	db := dataBase.Conn()

	var group models.Groups
	result := db.First(&group, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Grupo não encontrado")
		}
		return nil, result.Error
	}

	dataBase.CloseConnection(db)

	return &group, nil
}

func FindClientIDByGroup(group string) (int, error) {
	db := dataBase.Conn()

	var client_id int
	result := db.Table("groups").
		Select("client_id").Where("name_group LIKE ?", "%"+group+"%").Find(&client_id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("ClientID não encontrado")
		}
		return 0, result.Error
	}
	// if groups.NameGroup == "" {
	// 	return nil, fmt.Errorf("Grupo não encontrado")
	// }

	dataBase.CloseConnection(db)

	return client_id, nil

}

func SaveGroup(group *models.Groups) (*models.Groups, error) {
	db := dataBase.Conn()

	result := db.Table("groups").Create(group)
	if result.Error != nil {
		return nil, fmt.Errorf("Repository: Erro ao salvar um grupo")
	}

	dataBase.CloseConnection(db)

	return group, nil
}

func DeleteGroup(id string) error {
	db := dataBase.Conn()

	result := db.Table("groups").Where("id = ?", id).Delete(&models.Groups{})
	if result.Error != nil {
		return fmt.Errorf("Repository: Erro ao excluir o grupo")
	}

	dataBase.CloseConnection(db)

	return nil
}
