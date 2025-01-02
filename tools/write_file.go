package tools

import (
	"fmt"
	"os"
)

func UpdateTenantID2File(tenantID, filePath string) error {
	fmt.Println("写入标识符至文件：" + filePath)
	writeFile(filePath, []byte(tenantID))
	return nil
}

// func writeFile(filename string, data []byte) error {
// 	fmt.Println("write file:", filename)
// 	return ioutil.WriteFile(filename, data, 0644)
// }

func writeFile(filename string, data []byte) {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(filename)
			if err != nil {
				fmt.Println("create file error:", filename)
			}
		}
	}

	f.Truncate(0)
	f.Write(data)
	f.Close()
}
