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
wails build -v 2 -trimpath -clean -o mtest.exe
```
