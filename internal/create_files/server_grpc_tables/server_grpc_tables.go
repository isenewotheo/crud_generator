package server_grpc_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc_server
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.IsGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы grpc_server
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы grpc_server
		if config.Settings.NEED_CREATE_GRPC_SERVER_TEST == true {
			err = CreateFilesTest(Table1)
			if err != nil {
				log.Error("CreateFilesTest() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//UPDATE_EVERY_COLUMN
		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			//файлы grpc_server update
			err = CreateFilesUpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}

			//тестовые файлы grpc_server update
			if config.Settings.NEED_CREATE_GRPC_SERVER_TEST == true {
				err = CreateFilesUpdateEveryColumnTest(Table1)
				if err != nil {
					log.Error("CreateFilesTest() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

		}

		//NEED_CREATE_CACHE_API
		if config.Settings.NEED_CREATE_CACHE_API == true {
			//файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_FILES == true {
				err = CreateFilesCache(Table1)
				if err != nil {
					log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

			//тестовые файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_TEST_FILES == true {
				err = CreateFilesCacheTest(Table1)
				if err != nil {
					log.Error("CreateFilesTest() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

		}
	}
	return err
}

// CreateFiles - создаёт 1 файл в папке grpc_server
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "server_grpc.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + ".go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudTableURL := create_files.FindCrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

		if Table1.PrimaryKeyColumnsCount == 1 {
			TextGRPCServer = ReplaceIDRequestID_1PK(TextGRPCServer, Table1)
		}

		//замена "m.ID = AliasFromInt(ID)"
		TextGRPCServer = create_files.ReplacePrimaryKeyM_ID(TextGRPCServer, Table1)

		//замена "ID := Request.ID"
		TextGRPCServer = create_files.ReplacePrimaryKeyOtvetID(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

		//замена int64(ID) на ID
		TextGRPCServer = create_files.ReplaceIDtoID(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

		//удалим лишние функции
		TextGRPCServer = DeleteFuncDelete(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncRestore(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncFind_byExtID(TextGRPCServer, Table1)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.ConvertRequestIdToAlias(TextGRPCServer, Table1)
	}

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesTest - создаёт 1 файл в папке grpc_server
func CreateFilesTest(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "server_grpc_test.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//удалим лишние функции
		TextGRPCServer = DeleteFuncTestDelete(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncTestRestore(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncTestFind_byExtID(TextGRPCServer, Table1)

		//добавим импорты
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudStarterURL := create_files.FindCrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ConstantsURL := create_files.FindConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

		TextGRPCServer = create_files.CheckAndAddImport(TextGRPCServer, "encoding/json")

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextGRPCServer = create_files.ReplaceConnect_WithApplicationName(TextGRPCServer)

		if Table1.PrimaryKeyColumnsCount > 1 {
		}

		//Postgres_ID_Test = ID Minimum
		TextGRPCServer = create_files.Replace_Model_ID_Test(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

		//замена Otvet.ID = -1
		TextGRPCServer = create_files.ReplaceModelIDEqual1(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

		//
		TextGRPCServer = create_files.ReplaceOtvetIDEqual0(TextGRPCServer, Table1)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//Postgres_ID_Test = ID Minimum
	//if Table1.IDMinimum != "" {
	//	TextFind := "const " + ModelName + "_ID_Test = "
	//	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, TextFind+"0", TextFind+Table1.IDMinimum)
	//}

	// замена ID на PrimaryKey
	TextGRPCServer = create_files.ReplacePrimaryKeyOtvetID(TextGRPCServer, Table1)
	TextGRPCServer = create_files.ReplacePrimaryKeyM_ID(TextGRPCServer, Table1)

	//SkipNow()
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//замена импортов на новые URL
	//TextGRPCServer = create_files.ReplaceServiceURLImports(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// "+ModelName+"_Delete ")

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore ")

	return Otvet
}

//// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
//func DeleteFuncDeleteCtx(Text, ModelName string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Delete_ctx ")
//
//	return Otvet
//}
//
//// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
//func DeleteFuncRestoreCtx(Text, ModelName string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore_ctx ")
//
//	return Otvet
//}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_FindByExtID ")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_Delete")

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"Restore")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncTestFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_FindByExtID")

	return Otvet
}

// CreateFilesUpdateEveryColumn - создаёт 1 файл в папке grpc_server
func CreateFilesUpdateEveryColumn(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServerFunc := DirTemplatesGRPCServer + constants.SERVER_GRPC_TABLE_UPDATE_FUNC_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServerUpdateFunc := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_update.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateGRPCServerFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunc, " error: ", err)
	}
	TextGRPCServerUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateGRPCServerUpdate := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_TABLE_UPDATE_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerUpdate)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerUpdate, " error: ", err)
	}
	TextGRPCServer := string(bytes)
	TextGRPCServer = TextGRPCServer + "\n"

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		DBConstantsURL := create_files.FindDBConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, DBConstantsURL)

		CrudTableURL := create_files.FindCrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumn(TextGRPCServerUpdateFunc, Table1)
	// пустой файл не нужен
	if TextUpdateEveryColumn == "" {
		return err
	}

	TextGRPCServer = TextGRPCServer + TextUpdateEveryColumn

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.ConvertRequestIdToAlias(TextGRPCServer, Table1)
		TextGRPCServer = create_files.CheckAndAddImportAlias(TextGRPCServer)

		//замена RequestId{}
		TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

	}

	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.DeleteEmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServerUpdateFunc, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumn - возвращает текст для всех таблиц
func FindTextUpdateEveryColumn(TextGRPCServerUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumn1(TextGRPCServerUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumn1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumn1(TextGRPCServerUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPCServerUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	_, _, TextRequestFieldGolang, TextGolangLine := create_files.FindTextProtobufRequest_ID_Type(Table1, Column1, "Request.")
	//if Table1.PrimaryKeyColumnsCount > 1 {
	TextRequest := create_files.FindTextProtobufRequest_Column_ManyPK(Table1, Column1)
	//ColumnPK := create_files.FindPrimaryKeyColumn(Table1)
	IsPrimaryKey := create_files.IsPrimaryKeyColumn(Table1, Column1)

	//замена ID на PrimaryKey
	Otvet = create_files.ReplacePrimaryKeyOtvetID(Otvet, Table1)
	Otvet = create_files.ReplacePrimaryKeyM_ID(Otvet, Table1)

	//ColumnNameGolang := create_files.ConvertGolangTypeToProtobufType(Table1, Column1, "m")

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Update", ModelName+"_"+FuncName)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	if IsPrimaryKey == true {
		Otvet = strings.ReplaceAll(Otvet, "\tColumnName := Request.FieldName\n", "")
	} else if TextGolangLine != "" {
		Otvet = strings.ReplaceAll(Otvet, "ColumnName := Request.FieldName", TextGolangLine)
	}
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "Request.FieldName", TextRequestFieldGolang)
	Otvet = strings.ReplaceAll(Otvet, "Model.ColumnName", "Model."+ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Model.Update()", "Model."+FuncName+"()")

	Otvet = ReplaceIDRequestID_1PK(Otvet, Table1)

	return Otvet
}

// ReplaceIDRequestID_1PK - замена "ID := Request.ID"
func ReplaceIDRequestID_1PK(Text string, Table1 *types.Table) string {
	Otvet := Text

	//замена ID := Request.ID
	ColumnPK := create_files.FindPrimaryKeyColumn(Table1)
	_, _, RequestFieldPK, GolangLinePK := create_files.FindTextProtobufRequest_ID_Type(Table1, ColumnPK, "Request.")

	if GolangLinePK != "" {
		Otvet = strings.ReplaceAll(Otvet, "ID := Request.ID", GolangLinePK)
	} else {
		Otvet = strings.ReplaceAll(Otvet, "ID := Request.ID", ColumnPK.NameGo+" := "+RequestFieldPK)
	}

	return Otvet
}

// CreateFilesUpdateEveryColumnTest - создаёт 1 файл в папке grpc_server
func CreateFilesUpdateEveryColumnTest(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServerFunc := DirTemplatesGRPCServer + constants.SERVER_GRPC_TABLE_UPDATE_FUNC_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServerUpdate := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_update_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateGRPCServerFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunc, " error: ", err)
	}
	TextGRPCServerUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateGRPCServerUpdate := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_TABLE_UPDATE_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerUpdate)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerUpdate, " error: ", err)
	}
	TextGRPCServer := string(bytes)
	TextGRPCServer = TextGRPCServer + "\n"

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		//ModelURL := create_files.FindModelURL()
		//TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelURL)

		CrudStarterURL := create_files.FindCrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ConstantsURL := create_files.FindConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextGRPCServer = create_files.ReplaceConnect_WithApplicationName(TextGRPCServer)

		//TextGRPCServer = create_files.ConvertRequestIdToAlias(TextGRPCServer, Table1)
	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumnTest(TextGRPCServerUpdateFunc, Table1)

	//Postgres_ID_Test = ID Minimum
	TextGRPCServer = create_files.Replace_Model_ID_Test(TextGRPCServer, Table1)

	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	// пустой файл не нужен
	if TextUpdateEveryColumn == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextUpdateEveryColumn

	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена RequestId{}
	TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

	//добавим импорт uuid
	TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

	//Import Timestamp
	TextGRPCServer = create_files.CheckAndAddImportTimestamp_FromText(TextGRPCServer)

	//SkipNow() если нет строк в БД
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.DeleteEmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServerUpdate, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumnTest - возвращает текст для всех таблиц
func FindTextUpdateEveryColumnTest(TextGRPCServerUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumnTest1(TextGRPCServerUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumnTest1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumnTest1(TextGRPCServerUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPCServerUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	_, TextRequestField, TextRequestFieldGolang, _ := create_files.FindTextProtobufRequest_ID_Type(Table1, Column1, "Request2.")
	TextModelColumnName := create_files.ConvertGolangTypeToProtobufType(Table1, Column1, "m")
	TextRequestID := create_files.FindTextProtobufRequest_ManyPK(Table1)

	//Postgres_ID_Test = ID Minimum
	Otvet = create_files.Replace_Model_ID_Test(Otvet, Table1)

	//if TextGolangLine != "" {
	//	Otvet = strings.ReplaceAll(Otvet, "value := Request.FieldName", TextGolangLine)
	//}

	//if Table1.PrimaryKeyColumnsCount == 1 {
	//} else {
	TextRequestString := create_files.FindTextProtobufRequest_Column_ManyPK(Table1, Column1)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId{}", "grpc_proto."+TextRequestID+"{}")
	//}

	Otvet = strings.ReplaceAll(Otvet, "Request.ColumnName", TextRequestFieldGolang)
	Otvet = strings.ReplaceAll(Otvet, "Request2.ColumnName", "Request2."+TextRequestField)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestString", "grpc_proto."+TextRequestString)
	Otvet = strings.ReplaceAll(Otvet, "m.ColumnName", TextModelColumnName)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Update(", ModelName+"_"+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	return Otvet
}

// CreateFilesCache - создаёт 1 файл в папке grpc_server
func CreateFilesCache(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesGRPCServer + constants.SERVER_GRPC_TABLE_CACHE_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyCache := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_cache.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateCache)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCache, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		DBConstantsURL := create_files.FindDBConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, DBConstantsURL)

		CrudTableURL := create_files.FindCrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

	}

	//TextGRPCServer = create_files.ReplaceIntFromProtoRequest(TextGRPCServer, Table1)

	//замена RequestId{}
	TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

	TextGRPCServer = create_files.ReplacePrimaryKeyOtvetID(TextGRPCServer, Table1)

	TextGRPCServer = create_files.ReplacePrimaryKeyM_ID(TextGRPCServer, Table1)

	if Table1.PrimaryKeyColumnsCount == 1 {
		ColumnPK := create_files.FindPrimaryKeyColumn(Table1)
		TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "ReplaceManyID(ID)", ColumnPK.NameGo)
		//ColumnPK := create_files.FindPrimaryKeyColumn(Table1)
	} else {
		TextIDMany := "ReplaceManyID(ID)"
		TextIDMany = create_files.ReplaceIDtoID_Many(TextIDMany, Table1)
		TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "ReplaceManyID(ID)", TextIDMany)
	}
	TextGRPCServer = create_files.ReplaceIDtoID_Many(TextGRPCServer, Table1)

	//добавим импорт uuid
	TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

	//добавим импорт alias
	TextGRPCServer = create_files.CheckAndAddImportAlias(TextGRPCServer)

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//	TextGRPCServer = create_files.ConvertRequestIdToAlias(TextGRPCServer, Table1)
	//}

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesCacheTest - создаёт 1 файл в папке grpc_server
func CreateFilesCacheTest(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesGRPCServer + constants.SERVER_GRPC_TABLE_CACHE_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyCache := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_cache_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateCache)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCache, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudStarterURL := create_files.FindCrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ConstantsURL := create_files.FindConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextGRPCServer = create_files.ReplaceConnect_WithApplicationName(TextGRPCServer)

		//Postgres_ID_Test = ID Minimum
		TextGRPCServer = create_files.Replace_Model_ID_Test(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//	TextGRPCServer = create_files.ConvertRequestIdToAlias(TextGRPCServer, Table1)
	//}

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//SkipNow()
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}
