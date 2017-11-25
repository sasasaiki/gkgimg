package imageStorage

import (
	"fmt"
	"mime/multipart"
	"os"
	"testing"
)

func Test_DirImageStorage_Add(t *testing.T) {
	const jpegPath = "testImages/test.jpg"
	fileJpeg, _ := os.Open(jpegPath)
	jpegFileHeader := multipart.FileHeader{Filename: fileJpeg.Name()}

	const pngPath = "testImages/test.png"
	filePng, _ := os.Open(pngPath)
	pngFileHeader := multipart.FileHeader{Filename: filePng.Name()}

	const textPath = "testImages/test.text"
	fileText, _ := os.Open(textPath)
	textFileHeader := multipart.FileHeader{Filename: fileText.Name()}

	type args struct {
		file       multipart.File
		fileHeader multipart.FileHeader
		fileName   string
		directory  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pngをpngとして保存できる",
			args: args{
				file:       openFile(pngPath),
				fileHeader: pngFileHeader,
				fileName:   "test",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "jpegをjpegとして保存できる",
			args: args{
				file:       openFile(jpegPath),
				fileHeader: jpegFileHeader,
				fileName:   "test",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "textも保存できるらしい",
			args: args{
				file:       openFile(textPath),
				fileHeader: textFileHeader,
				fileName:   "test",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "jpegをpngとして保存してみる.エラーは出ない",
			args: args{
				file:       openFile(jpegPath),
				fileHeader: pngFileHeader,
				fileName:   "test_jpg_to",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "pngをjpegとして保存してみる.エラーは出ない",
			args: args{
				file:       openFile(pngPath),
				fileHeader: jpegFileHeader,
				fileName:   "test_png_to",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "textをjpegとして保存してみる.エラーは出ないがひらけないファイルができる",
			args: args{
				file:       openFile(textPath),
				fileHeader: jpegFileHeader,
				fileName:   "test_text_to",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "閉じられたfileはエラーを返す",
			args: args{
				file:       closedFile(jpegPath),
				fileHeader: jpegFileHeader,
				fileName:   "test_text_to",
				directory:  "testImageStorage",
			},
			wantErr: true,
		},
	}

	im := new(DirImgStorage)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := im.SaveWithFileHeader(tt.args.file, &tt.args.fileHeader, tt.args.fileName, tt.args.directory); (err != nil) != tt.wantErr {
				t.Errorf("imageManager.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func openFile(path string) multipart.File {
	if !existFile(path) {
		return nil
	}
	file, _ := os.Open(path)
	return file
}

func closedFile(path string) multipart.File {
	file, _ := os.Open(path)
	file.Close()
	return file
}

func existFile(path string) bool {
	_, e := os.Stat(path)
	if e != nil {
		fmt.Println("=========== error!!", path+" が存在しません！！============")
		return false
	}
	return true
}

func Test_printError(t *testing.T) {
	type args struct {
		message string
		e       error
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printError(tt.args.message, tt.args.e)
		})
	}
}
