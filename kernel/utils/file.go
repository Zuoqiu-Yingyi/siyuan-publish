package utils

import "os"

// 判断所给路径文件/文件夹是否存在
// REF [golang 判断文件/路径是否存在、是否为文件夹/文件_Lucky小黄人的博客-CSDN博客](https://blog.csdn.net/qq_41767116/article/details/124792196)
func IsPathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if os.IsNotExist(err) { // IsNotExist 来判断，是不是不存在的错误
		// 如果返回的错误类型使用 os.isNotExist() 判断为true，说明文件或者文件夹不存在
		return false, nil
	} else {
		return false, err // 如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
	}
}
