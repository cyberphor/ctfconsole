
type Config struct {
	UIProtocol  string
	UIIP        string
	UIPort      int
	APIProtocol string
	APIIP       string
	APIPort     int
	APILogPath  string
	DBProtocol  string
	DBIP        string
	DBPort      int
	DBUsername  string
	DBPassword  string
	DBName      string
}

func (c Config) GetUIEndpoint() string {
	return c.UIProtocol + "://" + c.UIIP + ":" + strconv.Itoa(c.UIPort)
}

func (c Config) GetAPIEndpoint() string {
	return c.APIIP + ":" + strconv.Itoa(c.APIPort)
}

func (c Config) GetDBEndpoint() string {
	return c.DBProtocol + "://" + c.DBUsername + ":" + c.DBPassword + "@" + c.DBIP + "/" + c.DBName + "?sslmode=disable"
}

func Logger(logFilePath string) (*slog.Logger, error) {
	var file *os.File
	var err error
	var writer io.Writer
	var handler slog.Handler

	file, err = os.Create(logFilePath)
	writer = io.Writer(file)
	handler = slog.NewJSONHandler(writer, nil)
	return slog.New(handler), err
}

func getEnvStr(key string) (string, error) {
	var valueStr string
	var defined bool

	valueStr, defined = os.LookupEnv(key)
	if !defined {
		return valueStr, fmt.Errorf("%s is not defined", key)
	}

	return valueStr, nil
}

func getEnvInt(key string) (int, error) {
	var valueStr string
	var valueInt int
	var defined bool
	var err error

	valueStr, defined = os.LookupEnv(key)
	if !defined {
		return 0, fmt.Errorf("%s is not defined", key)
	}

	valueInt, err = strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("%s is not an integer", key)
	}

	return valueInt, nil
}

func GetConfig() (Config, error) {
	var config Config
	var err error

	// get ui parameters
	config.UIProtocol, err = getEnvStr("CTFCONSOLE_UI_PROTOCOL")
	config.UIIP, err = getEnvStr("CTFCONSOLE_UI_IP_ADDRESS")
	config.UIPort, err = getEnvInt("CTFCONSOLE_UI_PORT")

	// get api parameters
	config.APILogPath, err = getEnvStr("CTFCONSOLE_API_LOG_PATH")
	config.APIProtocol, err = getEnvStr("CTFCONSOLE_API_PROTOCOL")
	config.APIIP, err = getEnvStr("CTFCONSOLE_API_IP_ADDRESS")
	config.APIPort, err = getEnvInt("CTFCONSOLE_API_PORT")

	// get db parameters
	config.DBName, err = getEnvStr("CTFCONSOLE_DB_NAME")
	config.DBUsername, err = getEnvStr("CTFCONSOLE_DB_USER")
	config.DBPassword, err = getEnvStr("CTFCONSOLE_DB_PASSWORD")
	config.DBProtocol, err = getEnvStr("CTFCONSOLE_DB_PROTOCOL")
	config.DBIP, err = getEnvStr("CTFCONSOLE_DB_IP_ADDRESS")
	config.DBPort, err = getEnvInt("CTFCONSOLE_DB_PORT")

	return config, err
}
