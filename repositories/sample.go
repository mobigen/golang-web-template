package repositories

import (
	"github.com/mobigen/golang-web-template/common"
	"github.com/mobigen/golang-web-template/infrastructures/datastore"
	"github.com/mobigen/golang-web-template/models"
	"github.com/mobigen/golang-web-template/tools/util"
)

// Sample is struct of todo.
type Sample struct {
	*datastore.DataStore
}

// New is constructor that creates SampleRepository
func (Sample) New(handler *datastore.DataStore) *Sample {
	return &Sample{handler}
}

// GetAll get all sample from database(store)
func (repo *Sample) GetAll() (*[]models.Sample, error) {
	dst := new([]models.Sample)
	result := repo.Orm.Find(dst)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected <= 0 {
		return dst, common.ErrNoHaveResult
	}
	return dst, nil
}

// GetByID get sample whoes id match
func (repo *Sample) GetByID(id int) (*models.Sample, error) {
	dst := new(models.Sample)
	result := repo.Orm.Find(dst).Where(&models.Sample{ID: id})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected <= 0 {
		return dst, common.ErrNoHaveResult
	}
	return dst, nil
}

// Create create sample
func (repo *Sample) Create(input *models.Sample) (*models.Sample, error) {
	input.CreateAt = util.GetMillis()
	result := repo.Orm.Create(input)
	if result.Error != nil {
		return nil, result.Error
	}
	return input, nil
}

// Update update sample
func (repo *Sample) Update(input *models.Sample) (*models.Sample, error) {
	// Save/Update All Fields
	// repo.Orm.Save(input)

	// ID       int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	// Name     string `json:"name" gorm:"column:name;not null"`
	// Desc     string `json:"desc" gorm:"column:desc;size:256"`
	// CreateAt int64  `json:"createAt" gorm:"column:createAt;autoCreateTime:milli"`
	result := repo.Orm.Model(input).
		Where(&models.Sample{ID: input.ID}).
		Updates(
			map[string]interface{}{
				"name": input.Name,
				"desc": input.Desc,
			})
	if result.Error != nil {
		return nil, result.Error
	}
	return input, nil

}

// Delete delete sample from id(primaryKey)
func (repo *Sample) Delete(id int) (*models.Sample, error) {
	dst := new(models.Sample)
	result := repo.Orm.Find(dst).Where(&models.Sample{ID: id})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected <= 0 {
		return nil, common.ErrNoHaveResult
	}
	// Delete with additional conditions
	result = repo.Orm.Delete(&models.Sample{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return dst, nil
}
