/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var userConfigFile = "./user.ini"

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "manager user",
	Long: `user add
user login`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("user called")
		if len(args) > 0 {
			command := args[0]
			if strings.Compare(command, "add") == 0 {
				// add username password
				if len(args) != 3 {
					fmt.Println("参数错误 需要 username password")
					return
				}
				username := args[1]
				// 添加用户
				user := readConfigFile()
				if _, ok := user[username]; ok {
					// 已存在
					fmt.Println("用户" + username + "已存在")
					return
				}
				password := genPassword(username, args[2])
				appendFile(userConfigFile, username+":"+password)

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readConfigFile() map[string]string {
	// 文件不存在 创建
	exists, _ := PathExists(userConfigFile)
	if !exists {
		os.Create(userConfigFile)
	}
	//打开文件
	file, err := os.Open(userConfigFile)
	//判断文件打开是否错误
	if err != nil {
		fmt.Println("open file err:", err)
	}
	//要及时关闭文件，防止内存泄漏，defer延时关闭，文件会在函数使用完后关闭
	defer file.Close()
	//默认缓冲4096
	reader := bufio.NewReader(file)
	//循环读取文件
	user := make(map[string]string)
	for {
		str, err := reader.ReadString('\n') //每读取到一个换行就结束
		if err == io.EOF {                  //io.EOF表示文件末尾
			break
		}
		lines := strings.Split(str, ":")
		user[lines[0]] = lines[1]
	}

	return user
}

func appendFile(fileName, line string) error {
	fp, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm) // 读写方式打开

	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。

	_, err := fp.WriteString(line + "\n")
	return err

}

//判断文件是否存在
//path：要判断的文件路径
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	//当为空文件或文件夹存在
	if err == nil {
		return true, nil
	}
	//os.IsNotExist(err)为true，文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	//其它类型，不确定是否存在
	return false, err
}

func genPassword(username, password string) string {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte("go-todo" + username))
	st := m5.Sum(nil)
	result := hex.EncodeToString(st)
	return result
}
