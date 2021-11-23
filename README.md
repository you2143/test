# vote-app

## プロジェクトの構成
バックエンドはGo  
フロントエンドはNuxt.js(vue.js)   
で作成しています。  
このうち、フロントエンドの方はほとんどできていないので、必要に応じて好きなように書き換えてあげてください。

## プロジェクトの実行

### 環境変数の定義

プロジェクトのルートディレクトリに`.env.sample`というファイルがあるので、コピーして`.env`ファイルを作成してください

`MYSQL_HOST`以外はご自身の環境毎に独自で設定していただいて構いません

## プロジェクトの実行

Goを開発環境で実行するためにはDockerを起動後、`docker-compose up --build`を実行します

起動後に、`http://localhost:28080/swagger/index.html`のページにアクセスするとswaggerで定義されたAPI一覧が確認できます

APIの引数、返り値を変更した場合は、以下のコマンドをGolangを実行しているDocker内部で実行してください。

`swag init --parseDependency --parseInternal`

