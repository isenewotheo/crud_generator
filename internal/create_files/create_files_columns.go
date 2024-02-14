package create_files

import (
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
)

// Has_Column_ExtID_ConnectionID - возвращает true если есть поля ExtId и ConnectionID
func Has_Column_ExtID_ConnectionID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return Otvet
	}

	//
	_, ok = Table1.MapColumns["connection_id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_IsDeleted - возвращает true если есть поле is_deleted
func Has_Column_IsDeleted(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["is_deleted"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_ColumnType_Time - возвращает true если есть колонка с типом время
func Has_ColumnType_Time(Table1 *types.Table) bool {
	Otvet := false

	for _, Column1 := range Table1.MapColumns {
		SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
		if ok == false {
			log.Panic("GetMappings() ", Column1.Type, " error: not found")
		}
		if SQLMapping1.GoType == "time.Time" {
			Otvet = true
			break
		}
	}

	return Otvet
}

func Has_Column_ID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_ExtID - возвращает true если есть поле ext_id
func Has_Column_ExtID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_CreatedAt - возвращает true если есть поле created_at
func Has_Column_CreatedAt(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["created_at"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_ModifiedAt - возвращает true если есть поле modified_at
func Has_Column_ModifiedAt(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["modified_at"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_DeletedAt - возвращает true если есть поле deleted_at
func Has_Column_DeletedAt(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["deleted_at"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_TableNameID - возвращает true если есть поле table_name_id
func Has_Column_TableNameID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["table_name_id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_TableRowID - возвращает true если есть поле table_row_id
func Has_Column_TableRowID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["table_row_id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_IsGroup - возвращает true если есть поле is_group
func Has_Column_IsGroup(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["is_group"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_ParentID - возвращает true если есть поле parent_id
func Has_Column_ParentID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["parent_id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_Name - возвращает true если есть поле name
func Has_Column_Name(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["name"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_Description - возвращает true если есть поле description
func Has_Column_Description(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["description"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// ----

// Has_Columns_CommonStruct - возвращает true если есть все общие структуры
func Has_Columns_CommonStruct(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_ExtID(Table1) && Has_Column_CreatedAt(Table1) && Has_Column_ModifiedAt(Table1) && Has_Column_DeletedAt(Table1) && Has_Column_IsDeleted(Table1) && Has_Column_ID(Table1)

	return Otvet
}

// Has_Columns_NameStruct - возвращает true если есть колонки name + description
func Has_Columns_NameStruct(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_Name(Table1) && Has_Column_Description(Table1)

	return Otvet
}

// Has_Columns_Groups - возвращает true если есть колонки is_group + parent_id
func Has_Columns_Groups(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_IsGroup(Table1) && Has_Column_ParentID(Table1)

	return Otvet
}

// Has_Columns_ExtLink - возвращает true если есть колонки table_name_id + table_row_id
func Has_Columns_ExtLink(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_TableNameID(Table1) && Has_Column_TableRowID(Table1)

	return Otvet
}

// ----

// Is_Column_CommonStruct - возвращает true если это колонка ext_id, created_at, modified_at, deleted_at, id
func Is_Column_CommonStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "ext_id", "created_at", "modified_at", "deleted_at", "is_deleted", "id":
		Otvet = true
	}

	return Otvet
}

// Is_Column_NameStruct - возвращает true если это колонка name или description
func Is_Column_NameStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "name", "description":
		Otvet = true
	}

	return Otvet
}

// Is_Column_GroupsStruct - возвращает true если это колонка is_group, parent_id
func Is_Column_GroupsStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "is_group", "parent_id":
		Otvet = true
	}

	return Otvet
}

// Is_Column_ExtLinksStruct - возвращает true если это колонка table_name_id, table_row_id
func Is_Column_ExtLinksStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "table_name_id", "table_row_id":
		Otvet = true
	}

	return Otvet
}

// Is_Common_Сolumn - возвращает true если это общая колонка: table_name_id, table_row_id, is_group, parent_id, name или description, ext_id, created_at, modified_at, deleted_at, id
func Is_Common_Сolumn(Column1 *types.Column) bool {
	Otvet := false

	Otvet = Is_Column_CommonStruct(Column1) || Is_Column_NameStruct(Column1) || Is_Column_GroupsStruct(Column1) || Is_Column_ExtLinksStruct(Column1)

	return Otvet
}

// Is_NotNeedUpdate_Сolumn - возвращает true если это общая колонка: table_name_id, table_row_id, is_group, parent_id, ext_id, created_at, modified_at, deleted_at, id
func Is_NotNeedUpdate_Сolumn(Column1 *types.Column) bool {
	Otvet := false

	Otvet = Is_Column_CommonStruct(Column1) || Is_Column_GroupsStruct(Column1) || Is_Column_ExtLinksStruct(Column1)

	return Otvet
}
