package models

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/dao"
	"time"
)

// 细胞类型元素
type CellType int8

const (
	Zero      CellType = iota // 0
	FlashCard                 // 1
	Tag                       // 2
	Task                      // 3
)

// 细胞元素 最基本的单元
type Cell struct {
	ID          int64   `gorm:"primary_key"`  // id
	CellName    string  `gorm:"unique_index"` // 细胞名称
	CellDesc    string  // 细胞内容
	Pid         int64   // 父id
	CellType    int8    // 细胞类型 1flashCard 2tag
	Url         string  // 唯一url
	Source      string  // 数据来源内容
	SourceUrl   string  // 数据来源url
	LinkellList []*Cell // 关联卡片集合
	TagList     []*Cell // tag集合
	CellStatus  int8    // 细胞状态 0无状态 1待完成
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UpdatedUser string
	DeletedAt   *time.Time `sql:"index"`
}

func (cell Cell) ConvertToCell(card *trello.Card) Cell {
	cell.CellName = card.Name
	cell.CellDesc = card.Desc
	cell.CellType = int8(FlashCard)
	cell.Pid = -1 // 无父节点 0 表示root节点
	cell.Source = "TrelloCard_" + card.ID
	cell.SourceUrl = card.URL
	cell.CellStatus = int8(Zero)

	var tagList []*Cell
	// todo label 设置
	for _, label := range card.Labels {
		var tag *Cell
		tag.CellName = label.Name
		tag.CellType = int8(Tag)
		tag.CellStatus = int8(Zero)
		tagList = append(tagList, tag)
	}
	cell.TagList = tagList
	cell.CreatedAt = time.Now()
	cell.UpdatedAt = time.Now()
	cell.UpdatedUser = "trello_import"
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

func CreateOrUpdateCell(cell *Cell) (ID int64, err error) {
	newCell := new(Cell)
	db := dao.DB.Model(&newCell)
	db.Where("cell_name = ?", cell.CellName).Find(&newCell)
	if newCell.ID > 0 {
		cell.ID = newCell.ID
		db.Update(&cell)
	} else {
		db.Create(&cell)
	}

	return newCell.ID, nil
}

func GetCellList() (dataList []interface{}, err error) {
	var list []Cell
	// todo 获取集合 list
	if err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func GetCellTree() (dataList []interface{}, err error) {
	var parentList []*Cell
	// todo 查询 获取父节点
	dao.DB.Model(&parentList).Find(nil)
	if err != nil {
		return nil, err
	}
	for _, v := range parentList {
		parent := CellTree{v.ID, v.CellName, v.Pid, v.CellType, v.CellStatus, v.Url, []*CellTree{}}
		var childrenList []*Cell
		// todo 查询所有子节点
		if err != nil {
			return nil, err
		}
		for _, c := range childrenList {
			child := CellTree{c.ID, c.CellName, c.Pid, c.CellType, v.CellStatus, c.Url, []*CellTree{}}
			parent.Children = append(parent.Children, &child)
		}
		dataList = append(dataList, parent)
	}
	return dataList, nil
}
