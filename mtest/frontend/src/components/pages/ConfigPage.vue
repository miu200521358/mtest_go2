<template>
    <AppHeader />
    <main>
        <InputOrganism />
    </main>
</template>

<script lang="ts">
import AppHeader from '@/components/organisms/AppHeader.vue'
import InputOrganism from '@/components/organisms/InputOrganism.vue'

export default {
    components: {
        AppHeader,
        InputOrganism
    }
}
</script>
