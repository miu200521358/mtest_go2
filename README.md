# mmd_base_go

## WSLの環境構築

```
sudo apt -y install language-pack-ja language-pack-zh-hans language-pack-ko language-pack-en
sudo apt-get install unifont
sudo update-locale LANG=en_US.UTF8

wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
source ~/.bashrc

curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/master/install.sh | bash
source ~/.bashrc
nvm install v21.5.0

npm install reload --save-dev

go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

~/.bashrc
```
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin
```

```
wails init -n hello_wails -t vue-ts

cd hello_wails
wails dev -assetdir ./frontend/dist -wailsjsdir ./frontend/src

GOOS=windows GOARCH=amd64 wails build -v 2 -trimpath -clean -o hello_wails.exe
```

## Command (WSL)

```
wails init -n mtest -t vue-ts
cd mtest

初回
wails dev -assetdir ./frontend/dist -wailsjsdir ./frontend/src -reloaddirs ./frontend/src/views,./frontend/src/i18n/locales -save

2回目以降
wails dev
```

## フロントエンドのホットリロード

// https://c-a-p-engineer.github.io/tech/2022/12/08/vue3-vite-hotreload/
./frontend/vite.config.ts

```
export default defineConfig({
  plugins: [vue(), vueJsx()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  build: {
    rollupOptions: {
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`,
      },
    },
  },
  server: {
    watch: {
      usePolling: true,
    },
  },
});
```

## TypeScript + Vue3 + Vite + Quasar

```
wails init -n mtest -t vue-ts

frontend を削除

(base) miu@garnet:/mnt/c/MMD/mmd_base_go/mtest$ npm create vite@latest
✔ Project name: … frontend
✔ Select a framework: › Vue
✔ Select a variant: › TypeScript

Scaffolding project in /mnt/c/MMD/mmd_base_go/mtest/frontend...

Done. Now run:

  cd frontend
  npm install
  npm run dev

wails dev -assetdir ./frontend/dist -wailsjsdir ./frontend/src -reloaddirs ./frontend/src/components -save

npm install --save quasar @quasar/extras
npm install --save-dev @quasar/vite-plugin sass@^1.33.0

npm install vue-router@4

GOOS=windows GOARCH=amd64 wails build -v 2 -trimpath -clean -o hello_wails.exe

sudo apt install fonts-material-design-icons-iconfont -y

npm install -D tailwindcss@latest postcss@latest autoprefixer@latest

npm install -D @vue/tsconfig
```

```
wails init -n mtest2 -t vue-ts

wails dev -assetdir ./frontend/dist -wailsjsdir ./frontend/src -reloaddirs ./frontend/src/components -save

npm create vite@latest

```

```
https://github.com/iqonicdesignofficial/hope-ui-design-system?tab=readme-ov-file#method-2-using-cdn

wails dev

```


```
wails init -n mtest -t vue-ts
wails dev -assetdir ./frontend/dist -wailsjsdir ./frontend/src -reloaddirs ./frontend/src/views,./frontend/src/i18n/locales -save

npm install --save quasar @quasar/extras
npm install --save-dev @quasar/vite-plugin sass@^1.33.0
npm install -D tailwindcss@latest postcss@latest autoprefixer@latest
npm install -D @vue/tsconfig

npm install vue-i18n@9

```