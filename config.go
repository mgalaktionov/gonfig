package gonfig

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

// GetIntFlagFirst - возвращает целочисленный указатель
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение флага имеет приориет.
func GetIntFlagFirst(flagName, envKey, usage string, def int) *int {
	v, _ := LookupEnvOrInt(envKey, def)
	return flag.Int(flagName, v, usage)
}

// GetFloatFlagFirst - возвращает указатель типа float64.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение флага имеет приориет.
func GetFloatFlagFirst(flagName, envKey, usage string, def float64) *float64 {
	v, _ := LookupEnvOrFloat(envKey, def)
	return flag.Float64(flagName, v, usage)
}

// GetStringFlagFirst - возвращает строковый указатель.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение флага имеет приориет.
func GetStringFlagFirst(flagName, envKey, usage string, def string) *string {
	v, _ := LookupEnvOrString(envKey, def)
	return flag.String(flagName, v, usage)
}

// GetDurationFlagFirst - возвращает указатель типа time.Duration.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение флага имеет приориет.
func GetDurationFlagFirst(flagName, envKey, usage string, def time.Duration) *time.Duration {
	v, _ := LookupEnvOrDuration(envKey, def)
	return flag.Duration(flagName, v, usage)
}

// GetBoolFlagFirst - возвращает булевй указатель.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение флага имеет приориет.
func GetBoolFlagFirst(flagName, envKey, usage string, def bool) *bool {
	v, _ := LookupEnvOrBool(envKey, def)
	return flag.Bool(flagName, v, usage)
}

// GetIntEnvFirst - возвращает целочисленный указатель
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение переменной среды имеет приориет.
func GetIntEnvFirst(flagName, envKey, usage string, def int) *int {
	v, isDef := LookupEnvOrInt(envKey, def)
	f := flag.Int(flagName, v, usage)
	if isDef {
		return f
	}
	return &v
}

// GetFloatEnvFirst - возвращает указатель типа float64.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение переменной среды имеет приориет.
func GetFloatEnvFirst(flagName, envKey, usage string, def float64) *float64 {
	v, isDef := LookupEnvOrFloat(envKey, def)
	f := flag.Float64(flagName, v, usage)
	if isDef {
		return f
	}
	return &v
}

// GetStringEnvFirst - возвращает строковый указатель.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение переменной среды имеет приориет.
func GetStringEnvFirst(flagName, envKey, usage string, def string) *string {
	v, isDef := LookupEnvOrString(envKey, def)
	f := flag.String(flagName, v, usage)
	if isDef {
		return f
	}
	return &v
}

// GetDurationEnvFirst - возвращает указатель типа time.Duration.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение переменной среды имеет приориет.
func GetDurationEnvFirst(flagName, envKey, usage string, def time.Duration) *time.Duration {
	v, isDef := LookupEnvOrDuration(envKey, def)
	f := flag.Duration(flagName, v, usage)
	if isDef {
		return f
	}
	return &v
}

// GetBoolEnvFirst - возвращает булевй указатель.
// Значение читается из флага flagName и переменной среды envKey,
// если не задано ни одно из значений - будет возвращено значение по умолчанию (def),
// значение переменной среды имеет приориет.
func GetBoolEnvFirst(flagName, envKey, usage string, def bool) *bool {
	v, isDef := LookupEnvOrBool(envKey, def)
	f := flag.Bool(flagName, v, usage)
	if isDef {
		return f
	}
	return &v
}

// LookupEnvOrString - читает строковую переменную среды,
// если переменная среды задана - возвращает ее значение и false,
// если нет, возвращает значение по умолчанию (defaultVal) и true.
func LookupEnvOrString(key string, defaultVal string) (string, bool) {
	if val, ok := os.LookupEnv(key); ok {
		return val, false
	}
	return defaultVal, true
}

// LookupEnvOrInt - чиатает целочисленную переменную среды,
// если переменная среды задана - возвращает ее значение и false,
// если нет, возвращает значение по умолчанию (defaultVal) и true.
// Если при чтении значения произойдет ошибка, она будет залогирована.
func LookupEnvOrInt(key string, defaultVal int) (int, bool) {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		return v, false
	}
	return defaultVal, true
}

// LookupEnvOrFloat - чиатает переменную среды с плавающей точкой,
// если переменная среды задана - возвращает ее значение и false,
// если нет, возвращает значение по умолчанию (defaultVal) и true.
// Если при чтении значения произойдет ошибка, она будет залогирована.
func LookupEnvOrFloat(key string, defaultVal float64) (float64, bool) {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal(err)
		}
		return v, false
	}
	return defaultVal, true
}

// LookupEnvOrDuration - чиатает переменную среды типа time.Duration,
// если переменная среды задана - возвращает ее значение и false,
// если нет, возвращает значение по умолчанию (defaultVal) и true.
// Если при чтении значения произойдет ошибка, она будет залогирована.
func LookupEnvOrDuration(key string, defaultVal time.Duration) (time.Duration, bool) {
	if val, ok := os.LookupEnv(key); ok {
		v, err := time.ParseDuration(val)
		if err != nil {
			log.Fatal(err)
		}
		return v, false
	}
	return defaultVal, true
}

// LookupEnvOrBool - чиатает булевую переменную среды,
// если переменная среды задана - возвращает ее значение и false,
// если нет, возвращает значение по умолчанию (defaultVal) и true.
// Если при чтении значения произойдет ошибка, она будет залогирована.
func LookupEnvOrBool(key string, defaultVal bool) (bool, bool) {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			log.Fatal(err)
		}
		return v, false
	}
	return defaultVal, true
}
