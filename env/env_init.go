package env

import (
	"os"
	"strings"
)

// 环境变量处理, key转大写, 同时把`-./`转换为`_`
// a-b=>a_b, a.b=>a_b, a/b=>a_b, HelloWorld=>hello_world
func init() {
	replacer := strings.NewReplacer("-", "_", ".", "_", "/", "_")
	for _, env := range os.Environ() {
		envs := strings.SplitN(env, "=", 2)
		var key = trim(envs[0])
		if len(envs) != 2 || key == "" || strings.HasPrefix(key, "_") {
			continue
		}

		_ = os.Unsetenv(key)

		key = replacer.Replace(key)
		key = strings.ToUpper(strings.ReplaceAll(key, "__", "_"))
		_ = os.Setenv(key, envs[1])
	}
}
