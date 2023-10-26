package model

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, table1 := range MapAll {
		err = CreateFiles(table1)
		if err != nil {
			log.Error("CreateFiles() table: ", table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + constants.FolderReady + micro.SeparatorFile()
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + "pkg" + micro.SeparatorFile() + "model" + micro.SeparatorFile()

	FilenameTemplateModel := DirTemplatesModel + "model.go_"
	FilenameReadyModel := DirReadyModel + strings.ToLower(Table1.Name) + ".go"

	bytes, err := os.ReadFile(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//создание текста
	TextModelStruct, ModelName, err := FindTextModelStruct(Table1)
	TextModel = ReplaceModelStruct(TextModel, TextModelStruct)

	////(l LawsuitStatusType) = (b Branch)
	//TextTemplateVarModel := "(" + strings.ToLower(config.Settings.TEXT_TEMPLATE_MODEL[:1]) + " " + config.Settings.TEXT_TEMPLATE_MODEL
	//TextVarModel := "(" + strings.ToLower(ModelName[:1]) + " " + ModelName
	//TextModel = strings.ReplaceAll(TextModel, TextTemplateVarModel, TextVarModel)
	//
	////(l *LawsuitStatusType) = (b *Branch)
	//TextTemplateVarModel = "(" + strings.ToLower(config.Settings.TEXT_TEMPLATE_MODEL[:1]) + " *" + config.Settings.TEXT_TEMPLATE_MODEL
	//TextVarModel = "(" + strings.ToLower(ModelName[:1]) + " *" + ModelName
	//TextModel = strings.ReplaceAll(TextModel, TextTemplateVarModel, TextVarModel)

	//
	TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextModel = constants.TEXT_GENERATED + TextModel

	if config.Settings.HAS_IS_DELETED == true {
		TextModel = DeleteFuncDelete(TextModel, ModelName, Table1)
		TextModel = DeleteFuncRestore(TextModel, ModelName, Table1)
	}
	TextModel = DeleteFuncFind_byExtID(TextModel, ModelName, Table1)

	TextModel = AddImportTime(TextModel, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

func FindTextModelStruct(Table1 *types.Table) (string, string, error) {
	var Otvet string
	var ModelName string
	var err error

	TableName := Table1.Name
	ModelName = create_files.FindSingularName(TableName)
	ModelName = create_files.FormatName(ModelName)
	Table1.NameGo = ModelName

	Otvet = `// ` + ModelName + ` - model from table ` + TableName + `
type ` + ModelName + ` struct {
`

	//сортировка
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//цикл по всем колонкам
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		TextColumn := FindTextColumn(&Column1)
		Otvet = Otvet + TextColumn + "\n"
		Table1.MapColumns[key1] = Column1
	}

	Otvet = Otvet + "\n}"
	return Otvet, ModelName, err
}

func FindTextColumn(Column1 *types.Column) string {
	Otvet := ""
	//	Code string `json:"code" gorm:"column:code;default:0"`

	ColumnName := Column1.Name
	ColumnModelName := create_files.FormatName(Column1.Name)
	Column1.NameGo = ColumnModelName
	SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
	if ok == false {
		log.Panic("GetMappings() ", Column1.Type, " error: not found")
	}
	Type_go := SQLMapping1.GoType
	Column1.TypeGo = Type_go
	TextDefaultValue := FindTextDefaultValue(Type_go)
	TextPrimaryKey := FindTextPrimaryKey(Column1.IsIdentity)
	Description := Column1.Description
	Description = strconv.Quote(Description) //экранирование символов

	Otvet = Otvet + "\t" + ColumnModelName + " " + Type_go + "\t"
	Otvet = Otvet + "`json:\"" + ColumnName + "\""
	Otvet = Otvet + "\tgorm:\"column:" + ColumnName + TextPrimaryKey + TextDefaultValue + "\""
	Otvet = Otvet + "\tdb:\"" + ColumnName + "\""
	Otvet = Otvet + "`"
	Otvet = Otvet + "\t//" + Description

	return Otvet
}

func FindTextDefaultValue(Type_go string) string {
	var Otvet string

	sValue := ""
	switch Type_go {
	case "string":
		sValue = "\\\"\\\""
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		sValue = "0"
	case "time.Time":
		sValue = "null"
	}

	if sValue != "" {
		Otvet = ";default:" + sValue
	}

	return Otvet
}

func FindTextPrimaryKey(Is_identity bool) string {
	Otvet := ""

	if Is_identity == true {
		Otvet = ";primaryKey;autoIncrement:true"
	}

	return Otvet
}

func ReplaceModelStruct(TextTemplateModel, TextModelStruct string) string {
	Otvet := ""

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL

	//найдём начало и конец
	TextFind1 := "// " + ModelName
	pos1 := strings.Index(TextTemplateModel, TextFind1)
	if pos1 < 0 {
		TextFind1 := "type " + ModelName + " struct {"
		pos1 = strings.Index(TextTemplateModel, TextFind1)
	}

	if pos1 < 0 {
		log.Panic("ReplaceModelStruct() error: in model.go_ not found text: ", TextFind1)
	}

	s2 := TextTemplateModel[pos1:]
	TextFind1 = "}\n"
	posEnd := strings.Index(s2, TextFind1)
	if posEnd < 0 {
		log.Panic("ReplaceModelStruct() error: in model.go_ not found text: ", TextFind1)
	}

	//
	Otvet = TextTemplateModel[:pos1] + TextModelStruct + TextTemplateModel[pos1+posEnd+1:]

	return Otvet
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(TextModel, ModelName string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	//FirstSymbol := strings.ToLower(ModelName)[:1]
	TextFind := "Delete(*" + ModelName + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	TextFind = "\n// Delete "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	//FirstSymbol := strings.ToLower(Modelname)[:1]
	TextFind := "Restore(*" + Modelname + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	TextFind = "\n// Restore "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

func AddImportTime(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//если уже есть импорт
	pos1 := strings.Index(Otvet, `"time"`)
	if pos1 >= 0 {
		return Otvet
	}

	HasTimeColumn := FindHasTimeColumn(Table1)
	if HasTimeColumn == false {
		return Otvet
	}

	//
	pos1 = strings.Index(Otvet, "import (")
	if pos1 < 0 {
		log.Error("not found word: import (")
		return TextModel
	}

	Otvet = Otvet[:pos1+8] + "\n\t" + `"time"` + Otvet[pos1+8:]

	return Otvet
}

func FindHasTimeColumn(Table1 *types.Table) bool {
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

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

	//
	_, ok1 := Table1.MapColumns["ext_id"]
	//if ok == false {
	//	return Otvet
	//}

	//
	_, ok2 := Table1.MapColumns["connection_id"]
	if ok1 == true && ok2 == true {
		return Otvet
	}

	//FirstSymbol := strings.ToLower(Modelname)[:1]
	TextFind := "Find_ByExtID(*" + Modelname + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	TextFind = "\n// Find_ByExtID "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}
