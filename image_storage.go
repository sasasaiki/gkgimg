package imageStorage

import (
	"fmt"
	"image"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

)

//ImgStorageI イメージ保管に必要なメソッドを持つ
type ImgStorageI interface {
	SaveWithFileHeader(multipart.File, *multipart.FileHeader, string, string) error
	SaveWithOriginFileName(multipart.File, string, string, string) error
}

//DirImgStorage ディレクトリに保存する
type DirImgStorage struct {
}

//SaveWithFileHeader 拡張子を指定せず、渡したファイルの拡張子を使用して保存する
func (im *DirImgStorage) SaveWithFileHeader(file multipart.File, fileHeader *multipart.FileHeader, newFileName string, directory string) error {

	e := im.SaveWithOriginFileName(file, filepath.Ext(fileHeader.Filename), newFileName, directory)
	if e != nil {
		printError("AddWithAutoExtension()", e)
		return e
	}

	return nil
}

//SaveWithOriginFileName 新しく保存するファイル名には拡張子を指定せず、渡したファイルの拡張子を使用して保存する
// example:
// originFileName = hoge.png
// newFileName = fuga
func (im *DirImgStorage) SaveWithOriginFileName(file multipart.File, originFileExtension string, newFileName string, directory string) error {
	defer file.Close()
	storageFilePath := filepath.Join(directory, newFileName+originFileExtension)

	data, e := ioutil.ReadAll(file)
	if e != nil {
		printError("Add()でReadALL(file)に失敗", e)
		return e
	}

	e = ioutil.WriteFile(storageFilePath, data, 0600)
	if e != nil {
		printError("Add()でfileの保存に失敗", e)
		return e
	}

	return nil
}


}

func (im *DirImgStorage) Update() {
}

func (im *DirImgStorage) Delete() {

}

func (im *DirImgStorage) Get() {

}

func printError(message string, e error) {
	fmt.Println("in image-storage ", message, " error occurred", e)
}
