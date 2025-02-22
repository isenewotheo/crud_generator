package constants_file

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc proto
func CreateAllFiles() error {
	var err error

	err = CreateFileConstants()
	if err != nil {
		log.Error("CreateFileConstants() error: ", err)
		return err
	}

	return err
}

// CreateFileConstants - создаёт 1 файл в папке constants
func CreateFileConstants() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesConstants := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CONSTANTS + micro.SeparatorFile()
	DirReadyConstants := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CONSTANTS + micro.SeparatorFile()
	FilenameReadyConstants := DirReadyConstants + config.Settings.TEMPLATES_CONSTANTS_FILENAME
	FilenameReadyConstants = create_files.DeleteLastUnderline(FilenameReadyConstants)

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyConstants)

	FilenameTemplateConstants := DirTemplatesConstants + config.Settings.TEMPLATES_CONSTANTS_FILENAME
	bytes, err := os.ReadFile(FilenameTemplateConstants)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateConstants, " error: ", err)
	}
	TextConstants := string(bytes)

	//заменим название сервиса
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceNameNew := config.Settings.SERVICE_NAME
	TextConstants = strings.ReplaceAll(TextConstants, ServiceNameTemplate, ServiceNameNew)

	//запись файла
	err = os.WriteFile(FilenameReadyConstants, []byte(TextConstants), constants.FILE_PERMISSIONS)

	return err
}
