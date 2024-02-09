# poc-go-clean-architecture

## すぐに理解できなかった箇所

### ドメインのスコープ

ドメインの範囲は Usecase(Apprication)層に対して提供する業務ロジックを実装するための層であり、
Entities と Repostiry と Repository を介して、Entity オブジェクトとして取得する Service の実装が含まれる。

## SQLite 構築手順

```console
# db作成
sqlite3 poc_clean_architecture.db

# テーブル作成
create table user(id text primary key, name text, age int, created_at text);

# 終了
.quit
```

## デモ

```console
curl -X POST localhost:8080/create-user -H 'Content-Type: application/json' -d '{"name":"sato", "age":20}'
```

## 参考

- [実装クリーンアーキテクチャ](https://qiita.com/nrslib/items/a5f902c4defc83bd46b8)
- [go-clean-architecture](https://github.com/GSabadini/go-clean-architecture/tree/master)
- [3.2. ドメイン層の実装](https://terasolunaorg.github.io/guideline/current/ja/ImplementationAtEachLayer/DomainLayer.html#id3)
- [データベース管理ソフト SQLite 3 の使い方](https://funatsu-lab.github.io/open-course-ware/resource/sqlite3/)
