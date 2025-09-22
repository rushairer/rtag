package main

// Messages contains all translatable strings
type Messages struct {
	// Command descriptions
	RootShort string
	RootLong  string

	InitShort string
	InitLong  string

	AddShort string
	AddLong  string

	PushShort string
	PushLong  string

	ListShort string
	ListLong  string

	RmShort string
	RmLong  string

	LangShort string
	LangLong  string

	// Flag descriptions
	PushAllFlag string

	// User messages
	ErrorReadingRtagFile    string
	RtagFileEmptyOrNotExist string
	CurrentTags             string
	AddTagFailed            string
	AddTagSuccess           string
	ReadTagsFailed          string
	NoTagsFound             string
	TagNotExistInFile       string
	SpecifyTagOrUseAll      string
	AllTags                 string
	RemoveTagFailed         string
	RemoveTagSuccess        string
	EnterTag                string
	Exit                    string
	TagCannotBeEmpty        string
	ContinueAdding          string
	ReadInputFailed         string
	StartPushingTags        string
	CreateGitTag            string
	CreateTagFailed         string
	PushingTagsToRemote     string
	PushTagsFailed          string
	PushTagsSuccess         string
	TagAlreadyExists        string
	TagNotExist             string

	// Language command messages
	CurrentLanguage         string
	AvailableLanguages      string
	LanguageUsage           string
	EnvironmentVariable     string
	ConfigFile              string
	LanguageSetTo           string
	LanguagePreferenceSaved string
	LanguageChangeNote      string
	InvalidLanguage         string
}

// GetAllMessages returns messages for all supported languages
func GetAllMessages() map[Language]Messages {
	return map[Language]Messages{
		LangEN: getEnglishMessages(),
		LangZH: getChineseMessages(),
		LangFR: getFrenchMessages(),
		LangRU: getRussianMessages(),
	}
}

// getEnglishMessages returns English translations
func getEnglishMessages() Messages {
	return Messages{
		RootShort: "A CLI tool for managing release tags",
		RootLong:  "rtag is a command line tool for managing release tags with git integration.",

		InitShort: "Initialize .rtag file",
		InitLong:  "Initialize .rtag file for storing published tags.",

		AddShort: "Add a new tag",
		AddLong:  "Add a new tag to the .rtag file. If no tag is provided, interactive mode will be used.",

		PushShort: "Push tags to remote repository",
		PushLong:  "Push tags to remote repository. Use --all flag to push all tags.",

		ListShort: "List all tags",
		ListLong:  "List all tags from the .rtag file.",

		RmShort: "Remove a tag",
		RmLong:  "Remove a tag from the .rtag file.",

		LangShort: "Set or display current language",
		LangLong:  "Set the interface language or display current language settings.",

		PushAllFlag: "Push all tags",

		ErrorReadingRtagFile:    "Error reading .rtag file: %v",
		RtagFileEmptyOrNotExist: ".rtag file is empty or does not exist",
		CurrentTags:             "Current tags:",
		AddTagFailed:            "Failed to add tag: %v",
		AddTagSuccess:           "Successfully added tag: %s",
		ReadTagsFailed:          "Failed to read tags: %v",
		NoTagsFound:             "No tags found",
		TagNotExistInFile:       "Tag '%s' does not exist in .rtag file",
		SpecifyTagOrUseAll:      "Please specify a tag to push or use --all flag",
		AllTags:                 "All tags:",
		RemoveTagFailed:         "Failed to remove tag: %v",
		RemoveTagSuccess:        "Successfully removed tag: %s",
		EnterTag:                "Enter a tag (or 'exit' to quit): ",
		Exit:                    "Exit",
		TagCannotBeEmpty:        "Tag cannot be empty, please try again",
		ContinueAdding:          "Continue adding? (y/n): ",
		ReadInputFailed:         "Failed to read input: %v",
		StartPushingTags:        "Starting to push tags (timestamp: %s)...",
		CreateGitTag:            "Creating git tag: %s",
		CreateTagFailed:         "Failed to create tag %s: %v",
		PushingTagsToRemote:     "Pushing tags to remote repository...",
		PushTagsFailed:          "Failed to push tags: %v",
		PushTagsSuccess:         "Successfully pushed all tags",
		TagAlreadyExists:        "tag '%s' already exists",
		TagNotExist:             "tag '%s' does not exist",

		CurrentLanguage:         "Current language",
		AvailableLanguages:      "Available languages",
		LanguageUsage:           "Usage: rtag lang [%s]",
		EnvironmentVariable:     "Environment variable: RTAG_LANG=[%s]",
		ConfigFile:              "Config file: ~/.config/rtag/config",
		LanguageSetTo:           "Language set to %s",
		LanguagePreferenceSaved: "Language preference saved successfully.",
		LanguageChangeNote:      "Note: Language change will take effect on next command execution.",
		InvalidLanguage:         "Invalid language '%s'. Supported languages: %s",
	}
}

// getChineseMessages returns Chinese translations
func getChineseMessages() Messages {
	return Messages{
		RootShort: "发布标签管理工具",
		RootLong:  "rtag 是一个用于管理发布标签并集成 git 的命令行工具。",

		InitShort: "初始化 .rtag 文件",
		InitLong:  ".rtag 文件用于存储已发布的 tag。",

		AddShort: "添加新标签",
		AddLong:  "向 .rtag 文件添加新标签。如果未提供标签，将使用交互模式。",

		PushShort: "推送标签到远程仓库",
		PushLong:  "推送标签到远程仓库。使用 --all 标志推送所有标签。",

		ListShort: "列出所有标签",
		ListLong:  "列出 .rtag 文件中的所有标签。",

		RmShort: "删除标签",
		RmLong:  "从 .rtag 文件中删除标签。",

		LangShort: "设置或显示当前语言",
		LangLong:  "设置界面语言或显示当前语言设置。",

		PushAllFlag: "推送所有标签",

		ErrorReadingRtagFile:    "错误读取 .rtag 文件: %v",
		RtagFileEmptyOrNotExist: ".rtag 文件为空或不存在",
		CurrentTags:             "当前 tags:",
		AddTagFailed:            "添加 tag 失败: %v",
		AddTagSuccess:           "成功添加 tag: %s",
		ReadTagsFailed:          "读取 tags 失败: %v",
		NoTagsFound:             "没有找到任何 tags",
		TagNotExistInFile:       "Tag '%s' 不存在于 .rtag 文件中",
		SpecifyTagOrUseAll:      "请指定要推送的 tag 或使用 --all 标志",
		AllTags:                 "所有 tags:",
		RemoveTagFailed:         "删除 tag 失败: %v",
		RemoveTagSuccess:        "成功删除 tag: %s",
		EnterTag:                "请输入一个 tag (或输入 'exit' 退出): ",
		Exit:                    "退出",
		TagCannotBeEmpty:        "Tag 不能为空，请重新输入",
		ContinueAdding:          "是否继续添加? (y/n): ",
		ReadInputFailed:         "读取输入失败: %v",
		StartPushingTags:        "开始推送 tags (时间戳: %s)...",
		CreateGitTag:            "创建 git tag: %s",
		CreateTagFailed:         "创建 tag %s 失败: %v",
		PushingTagsToRemote:     "推送 tags 到远程仓库...",
		PushTagsFailed:          "推送 tags 失败: %v",
		PushTagsSuccess:         "成功推送所有 tags",
		TagAlreadyExists:        "tag '%s' 已存在",
		TagNotExist:             "tag '%s' 不存在",

		CurrentLanguage:         "当前语言",
		AvailableLanguages:      "可用语言",
		LanguageUsage:           "用法: rtag lang [%s]",
		EnvironmentVariable:     "环境变量: RTAG_LANG=[%s]",
		ConfigFile:              "配置文件: ~/.config/rtag/config",
		LanguageSetTo:           "语言已设置为%s",
		LanguagePreferenceSaved: "语言偏好保存成功。",
		LanguageChangeNote:      "注意: 语言更改将在下次命令执行时生效。",
		InvalidLanguage:         "无效语言 '%s'。支持的语言: %s",
	}
}

// getFrenchMessages returns French translations
func getFrenchMessages() Messages {
	return Messages{
		RootShort: "Outil CLI pour gérer les tags de version",
		RootLong:  "rtag est un outil en ligne de commande pour gérer les tags de version avec intégration git.",

		InitShort: "Initialiser le fichier .rtag",
		InitLong:  "Initialiser le fichier .rtag pour stocker les tags publiés.",

		AddShort: "Ajouter un nouveau tag",
		AddLong:  "Ajouter un nouveau tag au fichier .rtag. Si aucun tag n'est fourni, le mode interactif sera utilisé.",

		PushShort: "Pousser les tags vers le dépôt distant",
		PushLong:  "Pousser les tags vers le dépôt distant. Utilisez le flag --all pour pousser tous les tags.",

		ListShort: "Lister tous les tags",
		ListLong:  "Lister tous les tags du fichier .rtag.",

		RmShort: "Supprimer un tag",
		RmLong:  "Supprimer un tag du fichier .rtag.",

		LangShort: "Définir ou afficher la langue actuelle",
		LangLong:  "Définir la langue de l'interface ou afficher les paramètres de langue actuels.",

		PushAllFlag: "Pousser tous les tags",

		ErrorReadingRtagFile:    "Erreur lors de la lecture du fichier .rtag: %v",
		RtagFileEmptyOrNotExist: "Le fichier .rtag est vide ou n'existe pas",
		CurrentTags:             "Tags actuels:",
		AddTagFailed:            "Échec de l'ajout du tag: %v",
		AddTagSuccess:           "Tag ajouté avec succès: %s",
		ReadTagsFailed:          "Échec de la lecture des tags: %v",
		NoTagsFound:             "Aucun tag trouvé",
		TagNotExistInFile:       "Le tag '%s' n'existe pas dans le fichier .rtag",
		SpecifyTagOrUseAll:      "Veuillez spécifier un tag à pousser ou utiliser le flag --all",
		AllTags:                 "Tous les tags:",
		RemoveTagFailed:         "Échec de la suppression du tag: %v",
		RemoveTagSuccess:        "Tag supprimé avec succès: %s",
		EnterTag:                "Entrez un tag (ou 'exit' pour quitter): ",
		Exit:                    "Quitter",
		TagCannotBeEmpty:        "Le tag ne peut pas être vide, veuillez réessayer",
		ContinueAdding:          "Continuer à ajouter? (y/n): ",
		ReadInputFailed:         "Échec de la lecture de l'entrée: %v",
		StartPushingTags:        "Début de la poussée des tags (horodatage: %s)...",
		CreateGitTag:            "Création du tag git: %s",
		CreateTagFailed:         "Échec de la création du tag %s: %v",
		PushingTagsToRemote:     "Poussée des tags vers le dépôt distant...",
		PushTagsFailed:          "Échec de la poussée des tags: %v",
		PushTagsSuccess:         "Tous les tags ont été poussés avec succès",
		TagAlreadyExists:        "le tag '%s' existe déjà",
		TagNotExist:             "le tag '%s' n'existe pas",

		CurrentLanguage:         "Langue actuelle",
		AvailableLanguages:      "Langues disponibles",
		LanguageUsage:           "Usage: rtag lang [%s]",
		EnvironmentVariable:     "Variable d'environnement: RTAG_LANG=[%s]",
		ConfigFile:              "Fichier de config: ~/.config/rtag/config",
		LanguageSetTo:           "Langue définie sur %s",
		LanguagePreferenceSaved: "Préférence de langue sauvegardée avec succès.",
		LanguageChangeNote:      "Note: Le changement de langue prendra effet lors de la prochaine exécution de commande.",
		InvalidLanguage:         "Langue invalide '%s'. Langues supportées: %s",
	}
}

// getRussianMessages returns Russian translations
func getRussianMessages() Messages {
	return Messages{
		RootShort: "CLI инструмент для управления тегами релизов",
		RootLong:  "rtag - это инструмент командной строки для управления тегами релизов с интеграцией git.",

		InitShort: "Инициализировать файл .rtag",
		InitLong:  "Инициализировать файл .rtag для хранения опубликованных тегов.",

		AddShort: "Добавить новый тег",
		AddLong:  "Добавить новый тег в файл .rtag. Если тег не указан, будет использован интерактивный режим.",

		PushShort: "Отправить теги в удаленный репозиторий",
		PushLong:  "Отправить теги в удаленный репозиторий. Используйте флаг --all для отправки всех тегов.",

		ListShort: "Показать все теги",
		ListLong:  "Показать все теги из файла .rtag.",

		RmShort: "Удалить тег",
		RmLong:  "Удалить тег из файла .rtag.",

		LangShort: "Установить или показать текущий язык",
		LangLong:  "Установить язык интерфейса или показать текущие настройки языка.",

		PushAllFlag: "Отправить все теги",

		ErrorReadingRtagFile:    "Ошибка чтения файла .rtag: %v",
		RtagFileEmptyOrNotExist: "Файл .rtag пуст или не существует",
		CurrentTags:             "Текущие теги:",
		AddTagFailed:            "Не удалось добавить тег: %v",
		AddTagSuccess:           "Тег успешно добавлен: %s",
		ReadTagsFailed:          "Не удалось прочитать теги: %v",
		NoTagsFound:             "Теги не найдены",
		TagNotExistInFile:       "Тег '%s' не существует в файле .rtag",
		SpecifyTagOrUseAll:      "Пожалуйста, укажите тег для отправки или используйте флаг --all",
		AllTags:                 "Все теги:",
		RemoveTagFailed:         "Не удалось удалить тег: %v",
		RemoveTagSuccess:        "Тег успешно удален: %s",
		EnterTag:                "Введите тег (или 'exit' для выхода): ",
		Exit:                    "Выход",
		TagCannotBeEmpty:        "Тег не может быть пустым, попробуйте снова",
		ContinueAdding:          "Продолжить добавление? (y/n): ",
		ReadInputFailed:         "Не удалось прочитать ввод: %v",
		StartPushingTags:        "Начинаем отправку тегов (временная метка: %s)...",
		CreateGitTag:            "Создание git тега: %s",
		CreateTagFailed:         "Не удалось создать тег %s: %v",
		PushingTagsToRemote:     "Отправка тегов в удаленный репозиторий...",
		PushTagsFailed:          "Не удалось отправить теги: %v",
		PushTagsSuccess:         "Все теги успешно отправлены",
		TagAlreadyExists:        "тег '%s' уже существует",
		TagNotExist:             "тег '%s' не существует",

		CurrentLanguage:         "Текущий язык",
		AvailableLanguages:      "Доступные языки",
		LanguageUsage:           "Использование: rtag lang [%s]",
		EnvironmentVariable:     "Переменная окружения: RTAG_LANG=[%s]",
		ConfigFile:              "Файл конфигурации: ~/.config/rtag/config",
		LanguageSetTo:           "Язык установлен на %s",
		LanguagePreferenceSaved: "Языковые предпочтения успешно сохранены.",
		LanguageChangeNote:      "Примечание: Изменение языка вступит в силу при следующем выполнении команды.",
		InvalidLanguage:         "Недопустимый язык '%s'. Поддерживаемые языки: %s",
	}
}
