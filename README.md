# image-storage
画像をリサイズして保存する機能を持ったパッケージ

This package resize and store image file

## How to use

1.install
```
go get github.com/sasasaiki/image-storage

#If you use glide
#glide get github.com/sasasaiki/image-storage

```

2.use

example
```
import "github.com/sasasaiki/image-storage"

func main(){

	const originalFile = "originalFile.jpg"
	file, _ := os.Open("originalDir/"+originalFile)

	im := imageStorage.DirImgStorage{}
	e := im.SaveResizedImage(file, originalFile, "newFileName", "storeDir", 400, 0, 90)
	if e != nil {
		fmt.Println("error")
		return
	}
}
```

pngとjpgに対応しています。

サーバーで使う場合のサンプルプロジェクトがこちらにありますので参考にどうぞ。

It corresponds to png and jpg.

Here is a sample project for using on a server, so please refer to it.

https://github.com/sasasaiki/image-storage-server


## SaveAsItIs
渡されたファイルをそのまま保存する
SaveAsItIs(file multipart.File, originFileName, newFileName, directory string) error
もあります。
こっちはリサイズ等の処理を行わないのでimage以外もいけます。


