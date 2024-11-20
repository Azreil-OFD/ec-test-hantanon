import Aura from '@primevue/themes/aura';
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  mode: 'static',
  modules: [
   '@primevue/nuxt-module'
   ],
   css: ['~/assets/css/main.css'],
  router: {
      base: '/ec-test-hantanon/'
   },
   primevue: {
      options: {
          theme: {
              preset: Aura
          }
      }
  },
  postcss: {
   plugins: {
     tailwindcss: {},
     autoprefixer: {},
   },
 },

})
