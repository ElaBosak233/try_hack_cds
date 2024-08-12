package repository

import (
	"github.com/elabosak233/cloudsdale/internal/model"
	"github.com/elabosak233/cloudsdale/internal/model/request"
	"gorm.io/gorm"
)

type INoticeRepository interface {
	Find(req request.NoticeFindRequest) ([]model.Notice, int64, error)
	Create(notice model.Notice) (model.Notice, error)
	Update(notice model.Notice) (model.Notice, error)
	Delete(notice model.Notice) error
}

type NoticeRepository struct {
	db *gorm.DB
}

func NewNoticeRepository(db *gorm.DB) INoticeRepository {
	return &NoticeRepository{db: db}
}

func (t *NoticeRepository) Find(req request.NoticeFindRequest) ([]model.Notice, int64, error) {
	var notices []model.Notice
	applyFilters := func(q *gorm.DB) *gorm.DB {
		if req.ID != 0 {
			q = q.Where("id = ?", req.ID)
		}
		if req.GameID != 0 {
			q = q.Where("game_id = ?", req.GameID)
		}
		if req.Type != "" {
			q = q.Where("type = ?", req.Type)
		}
		return q
	}
	db := applyFilters(t.db.Table("notices"))
	var total int64 = 0
	result := db.Model(&model.Notice{}).Count(&total)
	db = db.Order("notices.id DESC")
	result = db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "username", "nickname", "email"})
		}).
		Preload("Team", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name", "email"})
		}).
		Preload("Challenge", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "title"})
		}).
		Find(&notices)
	return notices, total, result.Error
}

func (t *NoticeRepository) Create(notice model.Notice) (model.Notice, error) {
	result := t.db.Table("notices").Create(&notice)
	return notice, result.Error
}

func (t *NoticeRepository) Update(notice model.Notice) (model.Notice, error) {
	result := t.db.Table("notices").Model(&notice).Updates(&notice)
	return notice, result.Error
}

func (t *NoticeRepository) Delete(notice model.Notice) error {
	result := t.db.Table("notices").Delete(&notice)
	return result.Error
}
