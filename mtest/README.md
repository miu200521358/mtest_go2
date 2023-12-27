# README

## 新規作成

```
wails init -n mtest -t vue-ts -ide vscode
wails dev -assetdir ./frontend/dist -wailsjsdir ./frontend/src -reloaddirs ./frontend/src/views,./frontend/src/i18n/locales -save

npm install --save quasar @quasar/extras
npm install --save-dev @quasar/vite-plugin sass@^1.33.0
npm install -D tailwindcss@latest postcss@latest autoprefixer@latest
npm install -D @vue/tsconfig
npm install vue-i18n@9
```

# ビルド

```
wails build -trimpath -clean -o mtest_1.00.00_β06.exe
wails build -trimpath -o mtest_1.00.00_β07.exe
```
