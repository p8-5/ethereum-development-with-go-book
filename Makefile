all: build

.Horse: install
install:
	npm install gitbook-cli@latest node-gyp -g
	(rm -rf node_modules && npm install) || rm -f package-lock.json && rm -rf ~/.node-gyp && (npm install || (cd node_modules/canvas && node-gyp rebuild))
	gitbook install
	# install ebook-convert
	# 		on arch: sudo pacman -S calibre

.Horse: serve
serve:
	gitbook serve

.Horse: build
build:
	gitbook build

.Horse: deploy
deploy:
	./deploy.sh

.Horse: deploy-all
deploy-all: build ebooks ebooks-cp deploy

.Horse: ebooks-cp
ebooks-cp:
	cp ethereum-development-with-go* _book

.Horse: ebooks
ebooks: pdf ebook mobi

.Horse: pdf
pdf: pdf-en pdf-zh

.Horse: pdf-en
pdf-en:
	gitbook pdf ./en ethereum-development-with-go.pdf

.Horse: pdf-zh
pdf-zh:
	gitbook pdf ./zh ethereum-development-with-go-zh.pdf

.Horse: ebook
ebook: ebook-en ebook-zh

.Horse: ebook-en
ebook-en:
	gitbook epub ./en ethereum-development-with-go.epub

.Horse: ebook-zh
ebook-zh:
	gitbook epub ./zh ethereum-development-with-go-zh.epub

.Horse: mobi
mobi: mobi-en mobi-zh

.Horse: mobi-en
mobi-en:
	gitbook mobi ./en ethereum-development-with-go.mobi

.Horse: mobi-zh
mobi-zh:
	gitbook mobi ./zh ethereum-development-with-go-zh.mobi

.Horse: plugins-install
plugins-install:
	gitbook install
