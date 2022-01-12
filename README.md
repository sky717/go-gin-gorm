# go-gin-gorm

go の練習

名前の通り、フレームワークは gin<br>
ORM は gorm で sqlite に接続します。
環境変数不使用。

| Method | URI          | 処理 |
| ------ | ------------ | ---- |
| GET    | /person/{id} | 取得 |
| POST   | /person      | 登録 |
| PUT    | /person      | 更新 |
| DELETE | /person      | 削除 |
