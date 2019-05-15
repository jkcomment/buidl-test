# 動作手順

### 概要
任意の文字列を暗号化/復号化する。
特に指定がなかったため共通鍵暗号方式(AES)で実装(暗号化時のデータ形式はBase64)。

### 必要なもの
- Goインストール済(Ver 1.11以上)
- curl

## 1. 依存パッケージのダウンロード

```Terminal
$ export GO111MODULE=on
$ go mod download 
```

## 2. プログラム起動
```Terminal
$ go run main.go crypto.go
```

## 3. 動作確認
ターミナルからcurlコマンドで動作確認を行う

- encrypt
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"message":"jkcomment"}' localhost:8080/encrypt
"EpN8rVZuEJ9_n1VwPE9VW4JIQvMPTWQESQ=="

```
- decrypt
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"message":"EpN8rVZuEJ9_n1VwPE9VW4JIQvMPTWQESQ=="}' localhost:8080/decrypt
"jkcomment"
```