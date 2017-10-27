## ローカル開発環境

### 概要

ローカル環境で動作確認する方法


### 手順

1. ソースビルド(jarファイル生成)
2. Docker起動
3. Docker環境のTomcat起動
4. 動作確認


### 依存モジュール

|ツール名|バージョン|補足
|-----|----|---|
|ant|1.10.*||
|Docker| 17.09.0-ce|※動作確認バージョン|
|Docker Machine|0.12.2|※動作確認バージョン|


### 作業ディレクトリ

```
$ cd {絶対パス}/ebica2-java-api
```


### ビルド

```
$ ant
```

### Docker起動

```
$ docker build .
$ docker run -it -v {絶対パス}/WebContent:/var/lib/tomcat6/webapps/ebica -p {ホスト側のポート番号}:8080 {image id}

```
※ {ホスト側のポート番号}:ebica2.0開発用API呼び出し画面から動作確認する時は必ず80番ポートを指定する。

### Tomcat起動
Docker環境で以下のコマンドを実施

```
$ cd /etc/init.d/
$ tomcat8 start
```


### Tomcat停止
Docker環境で以下のコマンドを実施

```
$ cd /etc/init.d/
$ tomcat8 stop
```

###動作確認

- ブラウザ確認

http://192.168.99.100/ebica/\_projects/ebisol/ebica/dev/api_post.jsp
 
- コマンド確認

```
$ curl http://192.168.99.100/ebica/api -F rid=login -F userid=dev -F userpwd=password
```
※Docker起動時に80番ポートを指定しなかった場合は以下のようを実行する。

```
$ curl http://192.168.99.100:{ホスト側のポート番号}/ebica/api -F rid=login -F userid=dev -F userpwd=password
```



   
