# SecretMerger

本專案為多媒體安全課程期末報告，實作灰階 BMP 影像的位元反轉、合併與 PSNR 計算。

## 安裝與執行

1. 安裝 Go<br>
    如果你的系統為`Debian/Ubuntu`
    ```bash
    sudo apt install golang
    ``` 

    如果你的系統為`Fedora`
    ```bash
    sudo dnf install go
    ``` 

    如果你的系統為`openSUSE`
    ```bash
    sudo zypper install go
    ``` 

    如果你的系統為`Arch Linux`
    ```bash
    sudo pacman -S go
    ``` 
    
    如果你的系統為`macOS`或是原生套件庫太舊
    ```bash
    brew install go
    ``` 
    > 安裝 [homebrew](https://brew.sh/)
2. 下載和執行本專案：
   ```bash
   git clone <本專案網址>
   cd secretmerger

   go run ./cmd/main/main.go
   ```
3. 更改使用到的圖片
    > 編輯 cmd/main/main.go
    ```go
    func main() {
	    inputY := "testdata/peppers_gray.bmp"
	    inputX := "testdata/baboon_gray.bmp"
	    outputZ := "output/Z_combined.bmp"
	    outputReversed := "output/Z_reversed.bmp"
	    outputDoubleReversed := "output/Z_double_reversed.bmp"
    }
    ```
    其中 `inputX` 為載體圖片之路徑，`inputY` 為隱藏圖片之路徑