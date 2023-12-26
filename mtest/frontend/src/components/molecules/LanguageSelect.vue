<template>
    <div class="text-right mt-2 mr-2 px-2 py-1">
        <select class="bg-gray-500" v-model="locale" @change="onchangeLanguageHandle(locale)">
            <option v-for="item in languages" :key="item" :class="{ active: item === locale }" :value="item">
                {{ t("languages." + item) }}
            </option>
        </select>
    </div>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { LoadConfig, SaveConfig } from "@/wailsjs/go/main/App";

const { t, availableLocales: languages, locale } = useI18n();

const onchangeLanguageHandle = (item: string) => {
    console.log(item);
    if (item !== locale.value) {
        locale.value = item;
        SaveConfig("locale", [item], 1);
    }
};

onMounted(() => {
    onMounted(async () => {
        const savedLocale = await LoadConfig("locale");
        if (savedLocale) {
            locale.value = savedLocale[0];
        }
    });
});
</script>





