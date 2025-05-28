SPEC_PATH := ../jpid-spec/openapi.json

all: pretty_model remove_blobs put_readme

.PHONY: model generate

put_readme:
	echo "\
		# \`jpid-go\`\n\
		\n\
		Go言語で日本郵便株式会社（以下、会社といいます）の「郵便番号・デジタルアドレスAPI」(以下、当該役務といいます）を利用するためのクライアントライブラリです。\n\
		APIドキュメント バージョン `jq .info.version < $(SPEC_PATH)` に対応しています。\n\
		\n\
		このSDKは当該役務を活用した製品等の開発者の利便のために試験目的で提供されるものであり、会社の承認および推奨の元に配布されるものではありません。\n\
		本SDKは Apache 2.0 ライセンスで提供されており、本リポジトリに含まれる成果物には当該ライセンスが文面通り適用されます。\n\
		なお、本リポジトリからの成果物の公衆送信は予告なく停止される可能性があります。予めご了承ください。\n\
		\n\
		## 本リポジトリの成果物について\n\
		\n\
		本リポジトリの成果物は[OpenAPI generator](https://github.com/OpenAPITools/openapi-generator)（バージョン: `openapi-generator version`）により生成されました。\n\
		OpenAPI generatorの開発チームおよびGo generatorのメンテナに心より敬意を表します。\n\
		また、仕様書を提供頂きました日本郵便株式会社の関係各位に対し、深く御礼申し上げます。\n\
		\n\
		## 著者\n\
		\n\
		Nomura Suzume （S ｕｚｕＭＥ[att]ea.g1e.org）\n\
		\n\
		## Credit\n\
		\n\
		「§ゆ∞ＩＤ」および「デジタルアドレス」は日本郵便株式会社の登録商標（6881754,6872251）です。\n\
		" | sed -r -e "s/^[[:blank:]]*//g" > README.md


remove_blobs: README.md api
	rm -r api docs

pretty_model: generate
	go run tools/remove_comments/main.go
	gofmt -w *.go

generate:
	openapi-generator generate -g go \
		-i $(SPEC_PATH) \
		--package-name jpid \
		--skip-validate-spec \
		--git-repo-id jpid-go \
		--git-user-id g1eng \
		-o . 

clean: go.mod
	rm -r *.go test 
