package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"gopkg.in/gorp.v2"
)

type PgsqlRepo struct {
	DbMap *gorp.DbMap
}

func NewPgsqlRepo(dbMap *gorp.DbMap) *PgsqlRepo {
	return &PgsqlRepo{
		DbMap: dbMap,
	}
}

func (repo *PgsqlRepo) Save(group *models.Group) (groupId uuid.UUID, err error) {
	trans, err := repo.DbMap.Begin()
	if err != nil {
		return groupId, utils.GetAppError(err, "error while saving group", http.StatusInternalServerError)
	}
	err = trans.Insert(group)
	if err != nil {
		return groupId, utils.GetAppError(err, "error while saving group", http.StatusInternalServerError)
	}
	temp := make([]interface{}, len(group.Members))
	for i, val := range group.Members {
		temp[i] = val
	}
	err = trans.Insert(temp...)
	if err != nil {
		return groupId, utils.GetAppError(err, "error while saving group", http.StatusInternalServerError)
	}
	err = trans.Commit()
	if err != nil {
		return groupId, utils.GetAppError(err, "error while saving group", http.StatusInternalServerError)
	}
	return group.Id, nil
}

func (repo *PgsqlRepo) AddMembers(groupId uuid.UUID, members []*models.Member) error {
	temp := make([]interface{}, len(members))
	for i, val := range members {
		temp[i] = val
	}
	err := repo.DbMap.Insert(temp...)
	if err != nil {
		return utils.GetAppError(err, "error while adding members", http.StatusInternalServerError)
	}
	return nil
}

func (repo *PgsqlRepo) RemMembers(groupId uuid.UUID, memberIds []uuid.UUID) error {
	var memIds string
	for i, id := range memberIds {
		if i > 0 {
			memIds += ", "
		}
		memIds += "'" + id.String() + "'"
	}
	query := fmt.Sprintf("delete from members where group_id = '%s' and member_id in (%s)", groupId.String(), memIds)
	_, err := repo.DbMap.Exec(query)
	if err != nil {
		return utils.GetAppError(err, "error while removing group members", http.StatusInternalServerError)
	}
	return nil
}

func (repo *PgsqlRepo) GetMembersCount(groupId uuid.UUID, memberIds []uuid.UUID) (int, error) {
	var memIds string
	for i, id := range memberIds {
		if i > 0 {
			memIds += ", "
		}
		memIds += "'" + id.String() + "'"
	}
	query := fmt.Sprintf("select count(*) from members where group_id = '%s' and member_id in (%s)", groupId.String(), memIds)
	count, err := repo.DbMap.SelectInt(query)
	if err != nil {
		return 0, utils.GetAppError(err, "error while getting member count", http.StatusInternalServerError)
	}
	return int(count), nil
}

func (repo *PgsqlRepo) RenameGroup(groupId uuid.UUID, name string) error {
	panic("not implemented")
}

func (repo *PgsqlRepo) Delete(groupId uuid.UUID) error {
	_, err := repo.DbMap.Delete(&models.Group{Id: groupId})
	if err != nil {
		return utils.GetAppError(err, "error while deleting group", http.StatusInternalServerError)
	}
	return nil
}

func (repo *PgsqlRepo) FindByName(name string) (*models.Group, error) {
	result := models.Group{}
	err := repo.DbMap.SelectOne(&result, "select * from groups where name= $1", name)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.GetAppError(err, "error while finding group by name", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, nil
}

func (repo *PgsqlRepo) FindById(groupId uuid.UUID) (*models.Group, error) {
	result := models.Group{}
	err := repo.DbMap.SelectOne(&result, "select * from groups where id= $1", groupId)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.GetAppError(err, "error while finding group by Id", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	members := []*models.Member{}
	_, err = repo.DbMap.Select(&members, "select * from members where group_id= $1", groupId)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.GetAppError(err, "error while finding group by Id", http.StatusInternalServerError)
	}
	result.Members = members
	return &result, nil
}
