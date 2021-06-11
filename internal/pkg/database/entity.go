package database

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

var snowflakeNode *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	snowflakeNode = n
}

type UUIDPrimaryKey struct {
	ID uuid.UUID `gorm:"comment:UUID主键;primarykey;type:uuid;default:uuid_generate_v4()" `
}

func NewUUIDPrimaryKey() *UUIDPrimaryKey {
	return &UUIDPrimaryKey{ID: uuid.New()}
}

type AutoIncrementPrimaryKey struct {
	ID int `gorm:"comment:自增主键;primarykey;autoIncrement"`
}

type SnowflakePrimaryKey struct {
	ID *snowflake.ID `gorm:"comment:雪花主键;primarykey"`
}

func NewSnowFlakePrimaryKey() SnowflakePrimaryKey {
	id := snowflakeNode.Generate()
	return SnowflakePrimaryKey{ID: &id}
}

type Entity struct {
	CreatedTime time.Time `gorm:"comment:创建时间;index;autoCreateTime" `
	UpdatedTime time.Time `gorm:"comment:更新时间;autoUpdateTime" `
	IsDeleted   bool      `gorm:"comment:删除标记" `
}
