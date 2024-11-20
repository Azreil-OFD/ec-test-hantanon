import Aura from '@primevue/themes/aura';

export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    '@primevue/nuxt-module',
    '@nuxtjs/tailwindcss',
  ],
  css: ["~/assets/css/tailwind.css"],
  primevue: {
    options: {
        theme: {
            preset: Aura
        }
    }
}

})