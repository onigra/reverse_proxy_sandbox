![SmokeTest](https://github.com/onigra/reverse_proxy_sandbox/workflows/SmokeTest/badge.svg)

## やりたいこと

- nginxのリバプロの設定をスモークテストしたい

## 方針

- proxy先のhttpサーバをdocker-composeでまとめて立ち上げる
- テスト用の `server_name` を指定する
  - hostsにも追記する
- VirtualHostに対してhttpリクエスト投げるテストを実行

## やったやつ

- https://github.com/onigra/reverse_proxy_sandbox
