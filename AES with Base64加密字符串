# golang 用AES和Base64加密字符串
------

## golang 用AES和Base64加密字符串

```
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "io"
    "log"
)

func main() {
    key := []byte("a very very very very secret key") // 32 bytes
    plaintext := []byte("some really really really long plaintext")
    fmt.Printf("%s\n", plaintext)
    ciphertext, err := encrypt(key, plaintext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%0x\n", ciphertext)
    result, err := decrypt(key, ciphertext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", result)
}

// See alternate IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    b := base64.StdEncoding.EncodeToString(text)
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
    return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    if len(text) < aes.BlockSize {
        return nil, errors.New("ciphertext too short")
    }
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    data, err := base64.StdEncoding.DecodeString(string(text))
    if err != nil {
        return nil, err
    }
    return data, nil
}
```

输出结果：
```
some really really really long plaintext
15e8bd5984ed9ab70efcf1b182f45818bb297982845b64e86b97469cab58fca34851eb65df672aa74f34d57954562331a6b76aab63e5c69898b41da2d494fd5fa271699d8b0ffab4
some really really really long plaintext
```

## PKCS5Padding与PKCS5Unpadding
```
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5Unpadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

## ECB加密模式

```
block, err := des.NewCipher(key)
if err != nil {
    ...
}
bs := block.BlockSize()
src = PKCS5Padding(src, bs)
if len(src)%bs != 0 {
    ....
}
out := make([]byte, len(src))
dst := out
for len(src) > 0 {
    block.Encrypt(dst, src[:bs])
    src = src[bs:]
    dst = dst[bs:]
}
...

## ECB下的解密
```
block, err := des.NewCipher(key)
if err != nil {
    ...
}

out := make([]byte, len(src))
dst := out
bs := block.BlockSize()
if len(src)%bs != 0 {
    ...
}

for len(src) > 0 {
    block.Decrypt(dst, src[:bs])
    src = src[bs:]
    dst = dst[bs:]
}
out = PKCS5UnPadding(out)
```


```


```
