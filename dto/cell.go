package dto

import (
	"github.com/adlio/trello"
	"time"
)

// 细胞类型元素
type CellType int8

const (
	ZeroType      CellType = iota // 0
	FlashCardType                 // 1
	TagType                       // 2
	TaskType                      // 3
)

// 细胞元素 最基本的单元
type Cell struct {
	ID           int64   `gorm:"primary_key"`  // id
	CellName     string  `gorm:"unique_index"` // 细胞名称
	CellDesc     string  // 细胞内容
	Pid          int64   // 父id
	CellType     int8    // 细胞类型 1flashCard 2tag
	Source       string  // 数据来源内容
	SourceUrl    string  // 数据来源url
	LinkCellList []*Cell // 关联卡片集合
	TagList      []*Cell // tag集合
	CellStatus   int8    // 细胞状态 0无状态 1待完成
	TaskFlag     int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UpdateName   string
	DeletedAt    *time.Time `sql:"index"`
}

func (cell Cell) ConvertToCell(card *trello.Card) Cell {
	cell.CellName = card.Name
	cell.CellDesc = card.Desc
	cell.CellType = int8(FlashCardType)
	cell.Pid = -1 // -1表示无父节点 0表示root节点
	cell.Source = "TrelloCard_" + card.ID
	cell.SourceUrl = card.URL
	cell.CellStatus = int8(ZeroType)
	// 转换trello的label到tag
	var tagList []*Cell
	// todo label 设置
	for _, label := range card.Labels {
		var tag Cell
		tag.CellName = label.Name
		tag.CellType = int8(TagType)
		tag.CellStatus = int8(ZeroType)
		tagList = append(tagList, &tag)
	}
	cell.TagList = tagList
	cell.CreatedAt = time.Now()
	cell.UpdatedAt = time.Now()
	cell.UpdateName = "trello_import"
	return cell
}

// 树结构
type CellTree struct {
	ID         int64
	CellName   string
	Pid        int64
	CellType   int8
	CellStatus int8
	Url        string
	Children   []*CellTree
}
