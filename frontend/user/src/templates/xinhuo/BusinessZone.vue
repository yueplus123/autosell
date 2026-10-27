<template>
  <!-- 官网（展示模式）和代理分站用两套完全不同的专区页 -->
  <ShowcaseZone v-if="showcase" />
  <ResellerZone v-else />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAppStore } from '../../stores/app'
import { isMainShowcase } from '../../utils/storefrontZone'
import ShowcaseZone from './ShowcaseZone.vue'
import ResellerZone from './ResellerZone.vue'

const appStore = useAppStore()
const showcase = computed(() => isMainShowcase(appStore.config as any))
</script>
