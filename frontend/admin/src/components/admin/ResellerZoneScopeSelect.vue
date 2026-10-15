<template>
  <div class="grid gap-2">
    <Label>{{ t('admin.resellerZone.label') }}</Label>
    <Select :model-value="modelValue" @update:model-value="(v) => emit('update:modelValue', String(v || 'all'))">
      <SelectTrigger>
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="all">{{ t('admin.resellerZone.all') }}</SelectItem>
        <SelectItem value="ai">{{ t('admin.resellerZone.ai') }}</SelectItem>
        <SelectItem value="game">{{ t('admin.resellerZone.game') }}</SelectItem>
      </SelectContent>
    </Select>
    <p class="text-xs text-muted-foreground">{{ t('admin.resellerZone.hint') }}</p>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { Label } from '@/components/ui/label'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'

defineProps<{ modelValue: string }>()
const emit = defineEmits<{ (e: 'update:modelValue', value: string): void }>()
const { t } = useI18n()
</script>
