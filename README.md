[![LinuxUnitTest](https://github.com/nao1215/spare/actions/workflows/linux_test.yml/badge.svg)](https://github.com/nao1215/spare/actions/workflows/linux_test.yml)
[![MacUnitTest](https://github.com/nao1215/spare/actions/workflows/mac_test.yml/badge.svg)](https://github.com/nao1215/spare/actions/workflows/mac_test.yml)
[![reviewdog](https://github.com/nao1215/spare/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/spare/actions/workflows/reviewdog.yml)
[![Run Gosec](https://github.com/nao1215/spare/actions/workflows/security.yml/badge.svg)](https://github.com/nao1215/spare/actions/workflows/security.yml)
![Coverage](https://raw.githubusercontent.com/nao1215/octocovs-central-repo/main/badges/nao1215/spare/coverage.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/spare)](https://goreportcard.com/report/github.com/nao1215/spare)
![GitHub](https://img.shields.io/github/license/nao1215/spare)

# spare - Single Page Application Release Easily
Work in progress: 開発中の期間は、日本語で内容を説明する

## 開発する機能
- [ ] initサブコマンド
  - 現在のディレクトリに設定ファイル（yml）を作成する
  - 初期化時に対話形式、もしくはオプションでymlに設定する値を指定する
  - 設定ファイルには、各設定項目のコメントが付与される
  - 設定ファイルのモデルは、configパッケージで定義する
  - この設定ファイルの内容に基づいて、生成するAWSインフラの構成が変わる
  
- [ ] createサブコマンド
  - AWSインフラを構築する
  - 最もシンプルな構成は、S3バケットとCloudFront
  - DBを利用する場合（API経由ではなくSPAが直接DB2アクセスする場合）、AppRunnerとRDSを構築する。ただし、このインフラ構成の需要がない、もしくはコードの複雑化が見込まれる場合は劣後とする。
  
- [ ] deployサブコマンド
  - ビルドした成果物をS3バケットにアップロードする
  - ビルドの責務は、別のツールに任せる
  - CloudFrontのキャッシュを削除する
    
- [ ] deleteサブコマンド
  - AWSインフラ、SPAを削除する
    
- [ ] cloudformationサブコマンド
  - createサブコマンドが構築するAWSインフラのCloudFormationテンプレートを出力する

- [ ] validateサブコマンド
  - 設定ファイルの内容を検証する
  
### メモ
spareは、数あるSPAデプロイ手段の1つ（スペア）という意味で名付けた。
フロントエンドの開発者は、Amplifyの方が好みかもしれない。そのため、以下のインフラ構成を取り扱うことを視野に入れている。
- S3 + CloudFront
- Amplify 
- AppRunner + RDS