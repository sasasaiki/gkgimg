package imageStorage

import (
	"fmt"
	"mime/multipart"
	"os"
	"reflect"
	"testing"
)

const (
	TestDir  = "testImages/"
	JpegFile = "test.jpg"
	JpegPath = TestDir + JpegFile
	PngFile  = "test.jpg"
	PngPath  = TestDir + PngFile
	TextFile = "test.text"
	TextPath = TestDir + TextFile
)

func TestDirImgStorageSaveWithFileHeader(t *testing.T) {
	jpegFileHeader := multipart.FileHeader{Filename: JpegFile}

	pngFileHeader := multipart.FileHeader{Filename: PngFile}

	textFileHeader := multipart.FileHeader{Filename: TextFile}

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
				file:       openFile(PngPath),
				fileHeader: pngFileHeader,
				fileName:   "test",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "jpegをjpegとして保存できる",
			args: args{
				file:       openFile(JpegPath),
				fileHeader: jpegFileHeader,
				fileName:   "test",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "textも保存できるらしい",
			args: args{
				file:       openFile(TextPath),
				fileHeader: textFileHeader,
				fileName:   "test",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "jpegをpngとして保存してみる.エラーは出ない",
			args: args{
				file:       openFile(JpegPath),
				fileHeader: pngFileHeader,
				fileName:   "test_jpg_to",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "pngをjpegとして保存してみる.エラーは出ない",
			args: args{
				file:       openFile(PngPath),
				fileHeader: jpegFileHeader,
				fileName:   "test_png_to",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "textをjpegとして保存してみる.エラーは出ないがひらけないファイルができる",
			args: args{
				file:       openFile(TextPath),
				fileHeader: jpegFileHeader,
				fileName:   "test_text_to",
				directory:  "testImageStorage",
			},
			wantErr: false,
		},
		{
			name: "閉じられたfileはエラーを返す",
			args: args{
				file:       closedFile(JpegPath),
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

func TestPrintError(t *testing.T) {
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
