import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '../stores/app'
import { postAPI } from '../api'
import { debounceAsync } from '../utils/debounce'
import { useLocalized } from './useProduct'
import { usePageSeo } from './usePageSeo'
import { isMainShowcase } from '../utils/storefrontZone'

/**
 * 文章/公告详情共享逻辑（classic + vault 双模板共用）。
 * 完整保留原 views/BlogDetail.vue 的行为，仅抽离为 composable。
 */
export function useBlogDetail() {
  const route = useRoute()
  const { t } = useI18n()
  const appStore = useAppStore()
  const { formatPrice } = useLocalized()

  const loading = ref(true)
  const post = ref<any>(null)
  const relatedProducts = computed<any[]>(() => post.value?.related_products || [])

  const getLocalizedText = (jsonData: any) => {
    if (!jsonData) return ''
    const locale = appStore.locale
    return jsonData[locale] || jsonData['zh-CN'] || jsonData['en-US'] || ''
  }

  const formatDate = (dateString: string) => {
    if (!dateString) return ''
    const date = new Date(dateString)
    return date.toLocaleDateString(appStore.locale, {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    })
  }

  usePageSeo({
    title: () => post.value ? getLocalizedText(post.value.title) : '',
    description: () => post.value ? getLocalizedText(post.value.summary) : '',
    image: () => post.value?.thumbnail || '',
    canonicalPath: () => `/blog/${(route.params.slug as string) || ''}`,
    type: () => 'article',
  })

  // 官网（展示模式）的首页就是文章列表，文章详情直接返回首页
  const showcase = computed(() => isMainShowcase(appStore.config as any))

  const backLink = computed(() => {
    if (showcase.value && post.value?.type !== 'notice') return '/'
    if (!post.value) return '/blog'
    return post.value.type === 'notice' ? '/notice' : '/blog'
  })

  const backText = computed(() => {
    if (showcase.value && post.value?.type !== 'notice') return t('blogDetail.backToHome')
    if (!post.value) return t('blogDetail.backToBlog')
    return post.value.type === 'notice' ? t('blogDetail.backToNotice') : t('blogDetail.backToBlog')
  })

  const loadPost = async () => {
    loading.value = true
    try {
      const slug = route.params.slug as string
      const response = await postAPI.detail(slug)
      post.value = response.data.data || null
    } catch (error) {
      console.error('Failed to load post:', error)
      post.value = null
    } finally {
      loading.value = false
    }
  }

  const debouncedLoadPost = debounceAsync(loadPost, 300)

  onMounted(() => {
    loadPost()
  })

  onUnmounted(() => {
    debouncedLoadPost.cancel()
  })

  return {
    loading,
    post,
    relatedProducts,
    getLocalizedText,
    formatDate,
    formatPrice,
    backLink,
    backText,
  }
}
