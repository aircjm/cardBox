package cellDao

import (
	"github.com/aircjm/gocard/dao"
	"github.com/aircjm/gocard/dto"
)

func CreateOrUpdateCell(cell *dto.Cell) (ID int64, err error) {
	newCell := new(dto.Cell)
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
	var list []dto.Cell
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
	var parentList []*dto.Cell
	// todo 查询 获取父节点
	dao.DB.Model(&parentList).Find(nil)
	if err != nil {
		return nil, err
	}
	for _, v := range parentList {
		parent := dto.CellTree{v.ID, v.CellName, v.Pid, v.CellType, v.CellStatus, v.Url, []*dto.CellTree{}}
		var childrenList []*dto.Cell
		// todo 查询所有子节点
		if err != nil {
			return nil, err
		}
		for _, c := range childrenList {
			child := dto.CellTree{c.ID, c.CellName, c.Pid, c.CellType, v.CellStatus, c.Url, []*dto.CellTree{}}
			parent.Children = append(parent.Children, &child)
		}
		dataList = append(dataList, parent)
	}
	return dataList, nil
}
