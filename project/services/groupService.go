package services

import (
	"fmt"
	"saveDatabase/project/models"
	"saveDatabase/project/repositories"
)

func GroupFindByPhone(phone string) (*models.Groups, error) {

	group, err := repositories.GroupFindByPhone(phone)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func GroupFindByID(id string) (*models.Groups, error) {

	group, err := repositories.GroupFindByID(id)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func FindClientIDByGroup(group string) (int, error) {

	client_id, err := repositories.FindClientIDByGroup(group)
	if err != nil {
		return 0, err
	}

	return client_id, nil
}

func SaveGroup(group *models.Groups) (*models.Groups, error) {

	fmt.Println("Chegou Service")
	fmt.Println(group)

	persistedGroup, err := repositories.SaveGroup(group)
	if err != nil {
		return nil, err
	}

	fmt.Println("Depois da Repository")
	fmt.Println(persistedGroup)

	return persistedGroup, nil
}

func DeleteGroup(id string) error {

	fmt.Println("Chegou na Service")

	err := repositories.DeleteGroup(id)
	if err != nil {
		return err
	}

	fmt.Println("Saiu da Repository")

	return nil
}
