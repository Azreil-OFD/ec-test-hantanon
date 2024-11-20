<template>
    <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
        <h1 class="text-3xl font-bold mb-6">Необходима регистрация</h1>
        <p class="text-lg mb-4">Чтобы продолжить, вам необходимо зарегистрироваться.</p>
        <NuxtLink to="/auth"><button class="bg-blue-500 text-white p-2 rounded hover:bg-blue-600">
                Перейти к регистрации
            </button></NuxtLink>

    </div>
</template>

<script setup lang="ts">
// Здесь можно добавить дополнительную логику, если необходимо
</script>

<style scoped>
/* Добавьте дополнительные стили, если необходимо */
</style>