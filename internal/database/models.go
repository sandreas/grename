package database

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// DEFAULTS
const (
	TagGroupDefault = 0
	TagTypeDefault  = 0
)

// EXIF
const (
	TagGroupExif     = 1
	TagTypeExifMake  = 1
	TagTypeExifModel = 2
	TagTypeExifDate  = 3
)

type TagType uint
type TagGroup uint

//type Model struct {
//	ID         string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//}
// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }

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

/*
type Tag struct {
	ID         string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Address      string  `gorm:"index:addr"` // create index with name `addr` for address
	IgnoreMe     int     `gorm:"-"` // ignore this field

}



// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name string `gorm:"default:'galeone'"`

}

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `User`
type User struct {
	gorm.Model
	Name string
}

*/
