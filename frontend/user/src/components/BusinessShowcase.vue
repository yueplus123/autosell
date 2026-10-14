<template>
  <div class="business-showcase">
    <!-- 双业务入口卡片 -->
    <section class="relative z-10 py-10 md:py-12">
      <div :class="containerClass">
        <div class="mb-6 md:mb-8">
          <h2 :class="headingClass">{{ t('home.business.title') }}</h2>
          <p class="mt-2 text-sm text-muted-foreground">{{ t('home.business.description') }}</p>
        </div>

        <div class="grid grid-cols-1 gap-4 md:grid-cols-2 md:gap-5">
          <div
            v-for="line in lines"
            :key="line.key"
            class="group relative flex cursor-pointer min-h-[230px] flex-col justify-between overflow-hidden rounded-2xl p-6 text-white shadow-sm transition duration-300 hover:-translate-y-1 hover:shadow-xl md:min-h-[280px] md:p-8"
            :class="theme(line.key).card"
            @click="router.push(line.link)"
          >
            <!-- 装饰：光晕与大图标 -->
            <span class="pointer-events-none absolute -right-16 -top-16 h-56 w-56 rounded-full blur-3xl" :class="theme(line.key).glowA"></span>
            <span class="pointer-events-none absolute -bottom-20 -left-10 h-48 w-48 rounded-full blur-3xl" :class="theme(line.key).glowB"></span>
            <component
              :is="theme(line.key).icon"
              class="pointer-events-none absolute -bottom-6 -right-4 h-40 w-40 text-white/10 transition-transform duration-500 group-hover:-rotate-6 group-hover:scale-110 md:h-48 md:w-48"
              :stroke-width="1.25"
              aria-hidden="true"
            />

            <div class="relative space-y-3">
              <span class="inline-flex items-center gap-1.5 rounded-full bg-white/15 px-3 py-1 text-xs font-semibold tracking-wide backdrop-blur-sm">
                <component :is="theme(line.key).icon" class="h-3.5 w-3.5" aria-hidden="true" />
                {{ t(`home.business.${line.key}.eyebrow`) }}
              </span>
              <h3 class="text-2xl font-bold tracking-tight md:text-3xl">{{ t(`home.business.${line.key}.title`) }}</h3>
              <p class="max-w-md text-sm leading-relaxed text-white/85">{{ t(`home.business.${line.key}.description`) }}</p>
            </div>

            <div class="relative mt-6 space-y-4">
              <div class="flex flex-wrap gap-2">
                <template v-if="line.children.length">
                  <RouterLink
                    v-for="child in line.children"
                    :key="child.id"
                    :to="`/categories/${encodeURIComponent(String(child.slug))}`"
                    class="rounded-full border border-white/25 bg-white/10 px-3 py-1 text-xs font-medium transition hover:bg-white/25"
                    @click.stop
                  >
                    {{ getLocalizedText(child.name) || child.slug }}
                  </RouterLink>
                </template>
                <template v-else>
                  <span
                    v-for="tag in highlights(line.key)"
                    :key="tag"
                    class="inline-flex items-center gap-1 rounded-full border border-white/25 bg-white/10 px-3 py-1 text-xs font-medium"
                  >
                    <Check class="h-3 w-3" aria-hidden="true" />
                    {{ tag }}
                  </span>
                </template>
              </div>
              <RouterLink
                :to="line.link"
                class="inline-flex items-center gap-2 rounded-lg bg-white px-4 py-2.5 text-sm font-semibold transition group-hover:gap-3"
                :class="theme(line.key).cta"
                @click.stop
              >
                {{ t(`home.business.${line.key}.cta`) }}
                <ArrowRight class="h-4 w-4" aria-hidden="true" />
              </RouterLink>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 各业务线热门商品 -->
    <template v-if="showProducts">
      <section
        v-for="line in linesWithProducts"
        :key="`products-${line.key}`"
        class="relative z-10 py-8"
      >
        <div :class="containerClass">
          <div class="mb-5 flex items-end justify-between gap-4">
            <div class="flex items-center gap-3">
              <span class="grid h-9 w-9 place-items-center rounded-xl text-white" :class="theme(line.key).card">
                <component :is="theme(line.key).icon" class="h-[18px] w-[18px]" aria-hidden="true" />
              </span>
              <h2 :class="headingClass">{{ t(`home.business.${line.key}.hotTitle`) }}</h2>
            </div>
            <RouterLink :to="line.link" class="text-sm font-semibold text-muted-foreground transition-colors hover:text-foreground">
              {{ t('home.featured.viewAll') }}
            </RouterLink>
          </div>
          <div :class="productGridClass">
            <template v-for="(product, idx) in productsByKey[line.key]" :key="product.id">
              <slot name="product" :product="product" :index="idx" />
            </template>
          </div>
        </div>
      </section>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ArrowRight, Check, Gamepad2, Sparkles } from 'lucide-vue-next'
import { categoryAPI, productAPI } from '../api'
import { useLocalized } from '../composables/useProduct'
import type { PublicCategory } from '../utils/category'
import { resolveBusinessLines, type BusinessLineKey } from '../utils/businessLines'

const props = withDefaults(defineProps<{
  showProducts?: boolean
  productLimit?: number
  containerClass?: string
  headingClass?: string
  productGridClass?: string
}>(), {
  showProducts: false,
  productLimit: 8,
  containerClass: 'container mx-auto px-4',
  headingClass: 'theme-section-heading text-2xl md:text-3xl',
  productGridClass: 'grid grid-cols-2 gap-3 md:grid-cols-3 md:gap-4 lg:grid-cols-4',
})

const router = useRouter()
const { t } = useI18n()
const { getLocalizedText } = useLocalized()

const THEMES = {
  ai: {
    card: 'bg-linear-to-br from-indigo-600 via-violet-600 to-fuchsia-600',
    glowA: 'bg-sky-400/40',
    glowB: 'bg-fuchsia-400/30',
    cta: 'text-violet-700',
    icon: Sparkles,
  },
  game: {
    card: 'bg-linear-to-br from-rose-500 via-pink-500 to-amber-400',
    glowA: 'bg-yellow-300/40',
    glowB: 'bg-purple-400/30',
    cta: 'text-rose-600',
    icon: Gamepad2,
  },
} as const

const theme = (key: BusinessLineKey) => THEMES[key]

const HIGHLIGHT_COUNT = 3
const highlights = (key: BusinessLineKey): string[] =>
  Array.from({ length: HIGHLIGHT_COUNT }, (_, i) => t(`home.business.${key}.highlight${i + 1}`)).filter(Boolean)

const categories = ref<PublicCategory[]>([])
const lines = computed(() => resolveBusinessLines(categories.value))

const productsByKey = reactive<Record<BusinessLineKey, any[]>>({ ai: [], game: [] })
const linesWithProducts = computed(() => lines.value.filter((line) => line.category && productsByKey[line.key].length > 0))

const loadProducts = async () => {
  await Promise.all(
    lines.value.map(async (line) => {
      if (!line.category) return
      try {
        const res = await productAPI.list({ page: 1, page_size: props.productLimit, category_id: line.category.id })
        productsByKey[line.key] = res.data.data || []
      } catch (error) {
        console.error('Failed to load business line products:', error)
      }
    }),
  )
}

onMounted(async () => {
  try {
    const res = await categoryAPI.list()
    categories.value = res.data.data || []
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
  if (props.showProducts) await loadProducts()
})
</script>
