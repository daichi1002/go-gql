# マイグレーション実行手順

1. docker exec -it server bash
2. migrate --path migrations --database 'mysql://root@tcp(db:3306)/gql?charset=utf8&parseTime=True&loc=Local&timeout=10s' -verbose up
