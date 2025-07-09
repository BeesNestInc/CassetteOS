# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

### Changed

### Removed

### Security

## [v0.4.15-cs1.3.7]
- WiFiセットアップ機能にて、インタフェースのプロセスが残ってしまって、次の起動に失敗する不具合の修正

## [v0.4.15-cs1.3.6]
- ファイル名、フォルダ名にCasaOSが残っているのを修正

## [v0.4.15-cs1.3.5]
### 削除
- 使用しないクラウドドライブに関するコードを削除

## [v0.4.15-cs1.3.4]
### 削除
- Rcloneを使用しない方針にしたため、関連するコードを削除


## [v0.4.15-cs1.3.3]
### Changed
- セットアップスクリプトやそのフォルダ名をCasaOSからCassetteOSへ変更
- インストール時のマイグレーションスクリプトは不要だと判断しているが、まだ削除まではしていない。(build/scripts/migration)


## [v0.4.15-cs1.3.2]
### Changed
- OpenAPI上のWiFiAPIの定義で使用している`client`がほかのリポジトリと衝突していたため`wifi-client`に修正した

## [v0.4.15-cs1.3.1]

### 修正
- goreleaserから使用しないクラウドドライブのための関連のビルド時フラグを削除
- クラウドドライブ関連のコードはまだ残っている状態

## [v0.4.15-cs1.3.0]

### 変更
- モジュールパスなどを修正。  
  (e.g., `github.com/mayumigit/CasaOS` → `github.com/BeesNestInc/CassetteOS`)
- まだCasaOS固有部分は残された状態ではあるが、動作することを優先してあえてそのままにしてあります。
- READMEをCassetteOSように修正
- GitHubActionをCassetteOSように見直し

### 削除
- CasaOSのスナップショット画像や固有のドキュメントを削除

### 修正
- sync.Mapを含む構造体の値コピーにより発生するgo vetエラーを修正

## [v0.4.15-cs1.2.0]
### Changed
- Switched version update check to custom CassetteOS update server.
- Replaced version comparison logic to support custom versioning scheme (e.g. v0.4.18-cs1.0.0).
- Updated `VERSION` constant in `constants.go` to reflect CassetteOS versioning.

### Removed
- Removed deprecated ZeroTier-related APIs and constants.

## [0.4.15-cs1.1.0]
### Added
- Added API endpoints to configure Wi-Fi settings in both Access Point (AP) mode and Wi-Fi client mode.


## [0.4.15-cs1.0.0]

### Changed
- Replaced module paths to use our own GitHub fork instead of the original IceWhaleTech repository.
  (e.g., `github.com/IceWhaleTech/CasaOS` → `github.com/mayumigit/CasaOS`)