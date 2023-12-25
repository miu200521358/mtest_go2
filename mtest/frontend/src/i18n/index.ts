import { createI18n } from "vue-i18n";

import zhHans from "@/i18n/locales/zh-Hans.json";
import en from "@/i18n/locales/en.json";
import ja from "@/i18n/locales/ja.json";

const i18n = createI18n({
  locale: "ja",
  fallbackLocale: "ja",
  legacy: false,
  messages: {
    ja: ja,
    en: en,
    "zh-Hans": zhHans,
  },
});

export default i18n;
