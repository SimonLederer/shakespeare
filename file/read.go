package file

import (
	"os"
)
// ReadFile reads in a file from a given directory and returns as a string and a nil error
// If there are any errors, it will return a non-nil error
func ReadFile(fileName string) ([]byte, error){
	file, err:= os.Open(fileName)
	defer file.Close()
	if err == nil {
		fileInfo, err := os.Stat(fileName)
		if err == nil{
			fileLen:=fileInfo.Size()
			data:=make([]byte, fileLen)
			_,err:=file.Read(data)
			if err == nil{
				return data, nil
			}else{
				return nil, err
			}
		}else{
			return nil, err
		}
	}else{
		return nil, err
	}
}