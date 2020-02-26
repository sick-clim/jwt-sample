# jwt-sample

JWT(Json Web Token)のサンプル
ジョットと読む
* 発行者が鍵を使ってJSONを署名し、トークンとして使う
* 任意の情報を含める事が可能
* 発行者は鍵を使ってトークンの検証を行うため改竄を検知できる

## 構成

* ヘッダ、ペイロード、署名の3つで構成
* それぞれはBase64でエンコードされる
* .(dot)で結合される

## 作り方

アルゴリズムとTokenタイプをヘッダに設定する

```
# ヘッダ
echo -n '{"alg": "HS256","typ":"JWT"}' | base64
# ペイロード
echo -n '{"sub":"1234567890","iat":1516239022}' | base64
```

## Goで動作確認

```
curl "http://localhost:8070/public"
curl "http://localhost:8070/private"
# JWT発行
curl "http://localhost:8070/auth"
curl "http://localhost:8070/private" -H "Authorization:Bearer <JWT>"
```
JWT発行の際にはBasic認証を使用したりする

下記サイトでJWTのエンコード・デコードできる
https://jwt.io/
