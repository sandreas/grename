package database

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type TagType uint
type TagGroup uint

// DEFAULTS
const (
	TagGroupDefault uint8 = 0
	TagTypeDefault  uint8 = 0
)

// EXIF
const (
	TagGroupExif     uint8 = 1
	TagTypeExifMake  uint8 = 1
	TagTypeExifModel uint8 = 2
	TagTypeExifDate  uint8 = 3
)

type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Model) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

type Log struct {
	*Model
	Message string     `gorm:"type:varchar(1000)"`
	Context LogContext `json:"context,omitempty"`
}

type LogContext struct {
	Action string
	// Data map

}

type File struct {
	*Model
	MimeMediaType string `gorm:"type:varchar(30)"`
	MimeSubType   string `gorm:"type:varchar(50)"`
	Hash          string `gorm:"type:varchar(100)"`
	Location      string `gorm:"type:varchar(4096)"`
	Tags          []*FileTag
}

type Tag struct {
	*Model
	Value string `gorm:"type:varchar(255);unique_index"`
}

type FileTag struct {
	File   *File
	FileID uuid.UUID `gorm:"type:uuid"`
	Tag    *Tag
	TagID  uuid.UUID `gorm:"type:uuid"`
	Group  uint8     `gorm:"default:0"`
	Type   uint8     `gorm:"default:0"`
}
