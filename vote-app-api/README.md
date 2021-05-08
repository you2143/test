# datsumo-api

## プロジェクトの実行準備

### 環境変数の定義

プロジェクトのルートディレクトリに`.env.sample`というファイルがあるので、コピーして`.env`ファイルを作成してください

`ELASTIC_HOST`と`MYSQL_HOST`以外はご自身の環境毎に独自で設定していただいて構いません

## プロジェクトの実行

Goを開発環境で実行するためにはDockerを起動後、`docker-compose up --build`を実行します

APIの定義はswagとgin-swaggerというプラグインを使用して定義しています

起動後に、`http://localhost:10001/swagger/index.html`のページにアクセスするとswaggerで定義されたAPI一覧が確認できます

APIの引数、返り値を変更した場合は、以下のコマンドをGolangを実行しているDocker内部で実行してください。

`swag init --parseDependency --parseInternal`

## フォルダ構成について

- docker(DockerfileやDocker起動時の設定ファイルなど)

- docs(Swagによって自動で生成されるドキュメント関連)

- domain(ドメインモデルに関連するファイル群、クリーンアーキテクチャのEnterprise Business Rulesに相当)

- infrastructure(外部サービスとの連携部分、クリーンアーキテクチャのFrameworks & Driversに相当)

 - datastore(DBとのやりとりを行う。gormとの接続)

- interfaces(外部の技術要素をまとめたもの、クリーンアーキテクチャのinterface Adapterに相当)

 - controllers(ルーティングから呼び出される関数を定義したもの)

 - database(DBと連携するためのRepositoryを定義したもの)

- usecase(ユースケースに応じた関数を用意したもの)

 - presenter(outputの整形)

 - repository(controllerからdomainにアクセスするためのリポジトリ)